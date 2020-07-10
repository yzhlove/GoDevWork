package service

import (
	"context"
	"errors"
	"service_register_and_discovery_day01/discover"
)

type Service interface {
	HealthCheck() bool
	SayHello() string
	DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error)
}

var ErrNotServiceInstance = errors.New("instance are not existed")

type DiscoverBase struct {
	discoverClient discover.DiscoveryClient
}


