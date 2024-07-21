package handler

import (
	"easy-chat/apps/im/ws/internal/handler/user"
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
)

// 方法注册到路由

// RegisterHandlers
//
//	@Description: 注册路由到自定义的websocket中，加入到map
//	@param srv
//	@param svc
func RegisterHandlers(srv *websocket.Server, svc *svc.ServiceContext) {
	srv.AddRoutes([]websocket.Route{
		{
			Method:  "user.online",
			Handler: user.OnLine(svc),
		},
	})
}
