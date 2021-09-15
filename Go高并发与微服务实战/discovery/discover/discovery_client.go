package discover

import "log"

type DiscoveryClient interface{
	Register(serviceName, instanceId, healthCheckUrl string, instanceHost string, instancePort int, meta map[string] string, logger *log.Logger) bool
	DeRegister(instanceId string, logger *log.Logger) bool
	DiscoveryService(serviceName string, logger *log.Logger) []interface{}
}
