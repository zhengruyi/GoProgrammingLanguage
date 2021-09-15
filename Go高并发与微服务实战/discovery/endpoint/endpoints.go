package endpoint

import (
	"GoProgrammingLanguage/Go高并发与微服务实战/discovery/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type DiscoveryEndpoints struct {
	SayHelloEndpoint endpoint.Endpoint
	DiscoveryEndpoint endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}

type SayHelloRequest struct {
	
}

type SayHelloResponse struct {
	Message string `json:"message"`
}

func MakeSayHelloEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		message := svc.SayHello()
		return SayHelloResponse{Message: message},nil
	}
}

type DiscoveryRequest struct {
	ServiceName string
}

type DiscoveryResponse struct {
	Instances []interface{} `json:"instances"`
	Error string `json:"error"`
}

func MakeDiscoveryEndpoint(svc service.Service) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DiscoveryRequest)
		instances ,err := svc.DiscoveryService(ctx,req.ServiceName)
		var errString = ""
		if err != nil {
			errString = err.Error()
		}
		return &DiscoveryResponse{
			Instances: instances,
			Error: errString,
		},nil
	}
}

type HealthRequest struct {}
type HealthResponse struct {
	Status bool `json:"status"`
}

func MakeHealthCheckEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.HealthCheck()
		return HealthResponse{
			Status: status,
		}, nil
	}
}