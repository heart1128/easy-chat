package user

import (
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
)

// OnLine
//
//	@Description: 获取所有在线用户
//	@param svc
//	@return websocket.HandlerFunc
func OnLine(svc *svc.ServiceContext) websocket.HandlerFunc {
	// 返回处理函数，在server.go中就可以根据连接执行函数
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		uids := srv.GetUsers() // 为空就是获取所有连接上的用户，也就是在线用户

		u := srv.GetUsers(conn)
		err := srv.Send(websocket.NewMessage(u[0], uids), conn)
		srv.Info("err ", err)
	}
}
