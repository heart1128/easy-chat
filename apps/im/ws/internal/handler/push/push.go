package push

import (
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
	"easy-chat/apps/im/ws/ws"
	"github.com/mitchellh/mapstructure"
)

/*

	kafka消费者收到数据之后，插入数据库，然后通过这个push推送给消息目标用户

*/

// Push
//
//	@Description: 路由执行的函数，push给收信人
//	@param svc
//	@return websocket.HandlerFunc
func Push(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {

		var data ws.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err))
			return
		}

		// 发送的目标,取出conn
		rconn := srv.GetConn(data.RecvId)
		if rconn == nil {
			// todo 离线状态
		}

		srv.Infof("push msg %v", data)

		srv.Send(websocket.NewMessage(data.SendId, &ws.Chat{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendTime:       data.SendTime,
			Msg: ws.Msg{
				MType:   data.MType,
				Content: data.Content,
			},
		}), rconn)
	}
}
