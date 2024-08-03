package svc

import (
	"easy-chat/apps/user/api/internal/config"
	"easy-chat/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// 注册一个rpc client到服务上下文
	userclient.User
	*redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Redis:  redis.MustNewRedis(c.Redisx),
	}
}
