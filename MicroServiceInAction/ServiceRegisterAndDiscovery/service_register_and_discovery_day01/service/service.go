package service

import (
	"context"
	"errors"
	"service_register_and_discovery_day01/config"
	"service_register_and_discovery_day01/discover"
)

type Service interface {
	HealthCheck() bool
	SayHello() string
	DiscoverService(ctx context.Context, serviceName string) ([]interface{}, error)
}

var ErrNotServiceInstance = errors.New("instance are not existed")

type DiscoverBase struct {
	discoverClient discover.DiscoveryClient
}

func (service *DiscoverBase) DiscoveryService(ctx context.Context, serverName string) ([]interface{}, error) {
	instance := service.discoverClient.DiscoverServices(serverName, config.Logger)
	if len(instance) == 0 {
		return nil, ErrNotServiceInstance
	}
	return instance, nil
}

func (service *DiscoverBase) SayHello() string {
	return "Hello World"
}

func (service *DiscoverBase) HealthCheck() bool {
	return true
}
