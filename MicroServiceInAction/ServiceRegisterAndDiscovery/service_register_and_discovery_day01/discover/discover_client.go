package discover

import "log"

type DiscoveryClient interface {
	/*
		服务注册接口
		@param serviceName 服务名称
		@param instanceId	服务实例ID
		@param instanceHost	服务实例地址
		@prams instancePort 服务实例端口
		@param healthCheckUrl	心跳地址
		@param meta	服务实例元数据
	*/
	Register(serviceName, instanceId, healthCheckUrl string, instanceHost string, instancePort int, meta map[string]string, logger *log.Logger) bool

	/*
		服务注销接口
	*/
	DeRegister(instanceId string, logger *log.Logger) bool

	/*
		发现服务实例接口
	*/
	DiscoverServices(serviceName string, logger *log.Logger) []interface{}
}
