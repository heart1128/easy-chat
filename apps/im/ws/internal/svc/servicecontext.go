package svc

import "easy-chat/apps/im/ws/internal/config"

// ServiceContext 服务上下文对象
type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
