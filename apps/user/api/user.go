package main

import (
	"easy-chat/apps/user/api/internal/config"
	"easy-chat/apps/user/api/internal/handler"
	"easy-chat/apps/user/api/internal/svc"
	"easy-chat/pkg/configserver"
	"easy-chat/pkg/resultx"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"sync"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/dev/user.yaml", "the config file")
var wg sync.WaitGroup

func main() {
	flag.Parse()

	var c config.Config
	// conf.MustLoad(*configFile, &c)

	// 获取配置中心的配置文件信息
	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "10.3.237.104:3379",
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2", // sail项目管理的key
		Namespace:      "user",
		Configs:        "user-api.yaml",
		ConfigFilePath: "./etcd/conf",
		LogLevel:       "DEBUG",
	})).MustLoad(&c, func(bytes []byte) error {
		// var c config.Config
		// go-zero平滑重启
		proc.WrapUp()
		// 阻塞运行
		wg.Add(1)
		go func(c config.Config) {
			defer wg.Done()

			Run(c)
		}(c)
		return nil
	})

	if err != nil {
		panic(err)
	}

}

func Run(c config.Config) {
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 设置为自定义的错误和正确处理
	httpx.SetErrorHandlerCtx(resultx.ErrHandler(c.Name))
	httpx.SetOkHandler(resultx.OkHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
