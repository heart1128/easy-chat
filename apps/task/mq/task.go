package main

import (
	"easy-chat/apps/task/mq/internal/config"
	"easy-chat/apps/task/mq/internal/handler"
	"easy-chat/apps/task/mq/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

// 启动配置加载
var configFile = flag.String("f", "etc/dev/task.yaml", "the config file")

func main() {
	flag.Parse()

	// 1. 加载配置文件，从yaml中
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 2. 配置文件设置，加载服务上下文，保存了config到上下文
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c) // 加载配置上下文，需要传入配置文件

	// 3. 获取listen对象，可以获取多个消费者
	listen := handler.NewListen(ctx)
	// 创建消费者组
	serviceGroup := service.NewServiceGroup()
	for _, s := range listen.Services() {
		serviceGroup.Add(s)
	}

	fmt.Println("Starting mq queue at ...")
	serviceGroup.Start()
}
