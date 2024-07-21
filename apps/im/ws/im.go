package main

import (
	"easy-chat/apps/im/ws/internal/config"
	"easy-chat/apps/im/ws/internal/handler"
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"time"
)

// 启动配置加载
var configFile = flag.String("f", "etc/dev/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 加载go zero中的启动
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c) // 加载配置上下文，需要传入配置文件

	srv := websocket.NewServer(c.ListenOn,
		websocket.WithServerAuthentication(handler.NewJwtAuth(ctx)), //加入jwt鉴权
		websocket.WithServerMaxConnectionIdle(10*time.Second),       // 设置最大空闲时间10s
	) // 创建一个websocket,传入的是配置文件的监听地址
	defer srv.Stop()

	handler.RegisterHandlers(srv, ctx) // 加载路由

	fmt.Println("start websocket server at ", c.ListenOn, " ..... ")
	srv.Start()

}
