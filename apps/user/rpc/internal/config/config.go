package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// yaml中添加了配置，这里要添加配置属性加载配置
	Mysql struct {
		DataSource string
	}

	Cache cache.CacheConf

	Redisx redis.RedisConf

	Jwt struct {
		AccessSecret string
		AccessExpire int64
	}
}
