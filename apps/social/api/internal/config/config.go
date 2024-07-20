package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// 添加rpc客户端配置
	SocialRpc zrpc.RpcClientConf
	// 因为要获取到user信息
	UserRpc zrpc.RpcClientConf

	JwtAuth struct {
		AccessSecret string
		//AccessExpire int64
	}
}
