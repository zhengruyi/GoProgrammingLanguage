package service

import (
	"GoProgrammingLanguage/Go高并发与微服务实战/discovery/config"
	"GoProgrammingLanguage/Go高并发与微服务实战/discovery/discover"
	"context"
	"errors"
)

type Service interface {
	HealthCheck() bool
	SayHello() string
	DiscoveryService(ctx context.Context, serviceName string)([]interface{},error)
}

type DiscoveryServiceImpl struct {
	discoveryClient discover.DiscoveryClient
}


var ErrNotServiceInstances = errors.New("instances are not existed")

func (service *DiscoveryServiceImpl) HealthCheck() bool {
	return true
}

func (service *DiscoveryServiceImpl) SayHello() string {
	return "hello"
}

func (service *DiscoveryServiceImpl) DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error) {
	instances := service.discoveryClient.DiscoveryService(serviceName,config.Logger)
	if instances == nil || len(instances) == 0 {
		return nil, ErrNotServiceInstances
	}
	return instances, nil
}

func NewDiscoveryServiceImpl(discoveryClient discover.DiscoveryClient) Service  {
	return &DiscoveryServiceImpl{
		discoveryClient:discoveryClient,
	}
}
