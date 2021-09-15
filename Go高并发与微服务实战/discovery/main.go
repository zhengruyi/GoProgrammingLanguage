package main

import (
	"GoProgrammingLanguage/Go高并发与微服务实战/discovery/config"
	"GoProgrammingLanguage/Go高并发与微服务实战/discovery/discover"
	"GoProgrammingLanguage/Go高并发与微服务实战/discovery/endpoint"
	"GoProgrammingLanguage/Go高并发与微服务实战/discovery/service"
	"GoProgrammingLanguage/Go高并发与微服务实战/discovery/transport"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	var(
		servicePort = flag.Int("service.port",10086,"service port")
		serviceHost = flag.String("service.host","127.0.0.1","service host")
		serviceName = flag.String("service.name","SayHello","service name")
	)
	flag.Parse()
	ctx := context.Background()
	errChan := make(chan error)
	var discoveryClient discover.DiscoveryClient
	var svc = service.NewDiscoveryServiceImpl(discoveryClient)
	sayHelloEndpoint := endpoint.MakeSayHelloEndpoint(svc)
	discoveryEndpoint := endpoint.MakeDiscoveryEndpoint(svc)
	healthEndpoint := endpoint.MakeHealthCheckEndpoint(svc)
	endpts := endpoint.DiscoveryEndpoints{
		SayHelloEndpoint: sayHelloEndpoint,
		DiscoveryEndpoint: discoveryEndpoint,
		HealthCheckEndpoint: healthEndpoint,
	}
	//创建http.Handler
	r := transport.MakeHttpHandler(ctx, endpts, config.KitLogger)
	// 定义服务实例ID
	instanceId := *serviceName
	// 启动 http server
	go func() {
		config.Logger.Println("Http Server start at port:" + strconv.Itoa(*servicePort))
		//启动前执行注册
		if !discoveryClient.Register(*serviceName, instanceId, "/health", *serviceHost,  *servicePort, nil, config.Logger){
			config.Logger.Printf("string-service for service %s failed.", serviceName)
			// 注册失败，服务启动失败
			os.Exit(-1)
		}
		handler := r
		errChan <- http.ListenAndServe(":"  + strconv.Itoa(*servicePort), handler)
	}()

	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	//服务退出取消注册
	discoveryClient.DeRegister(instanceId, config.Logger)
	config.Logger.Println(error)
}

