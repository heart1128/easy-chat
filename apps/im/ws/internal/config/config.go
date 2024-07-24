package config

import "github.com/zeromicro/go-zero/core/service"

type Config struct {
	// 依赖go zero的log，性能监听等配置
	service.ServiceConf

	ListenOn string

	JwtAuth struct {
		AccessSecret string
	}

	Mongo struct {
		Url string
		Db  string
	}
}
