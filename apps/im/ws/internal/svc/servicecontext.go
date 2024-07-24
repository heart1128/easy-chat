package svc

import (
	"easy-chat/apps/im/immodels"
	"easy-chat/apps/im/ws/internal/config"
)

// ServiceContext 服务上下文对象
type ServiceContext struct {
	Config config.Config

	immodels.ChatLogModel // 数据库模型，MongoDB的
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		ChatLogModel: immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
	}
}
