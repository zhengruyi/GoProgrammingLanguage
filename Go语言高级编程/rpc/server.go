package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	//注册对应的服务名称和对应的方法
	rpc.RegisterName("HelloService",new(HelloService))
	listener,err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal("ListenTCP error :", err)
	}
	//阻塞等待客户端的链接
	conn, err :=  listener.Accept()
	if err != nil {
		log.Fatal("Accept error:",err)
	}
	//服务这个连接
	rpc.ServeConn(conn)
}