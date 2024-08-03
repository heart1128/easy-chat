package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// 加rpc, 就是用rpc的client调用
	UserRpc zrpc.RpcClientConf

	// 使用redis保存登录用户
	Redisx redis.RedisConf

	JwtAuth struct {
		AccessSecret string
		// AccessExpire int64
	}
}
