package push

import (
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
	"easy-chat/apps/im/ws/ws"
	"easy-chat/pkg/constants"
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

		// 根据群聊或者私聊类型判断并发发送还是单个发送
		switch data.ChatType {
		case constants.SingleChatType:
			single(srv, &data, data.RecvId)
		case constants.GroupChatType:
			group(srv, &data)
		}
	}
}

// single
//
//	@Description: 封装私聊，单发
//	@param srv
//	@param data
//	@param recvId
//	@return error
func single(srv *websocket.Server, data *ws.Push, recvId string) error {
	// 发送的目标,取出conn
	rconn := srv.GetConn(recvId)
	if rconn == nil {
		// todo 离线状态
	}

	srv.Infof("push msg %v", data)

	return srv.Send(websocket.NewMessage(data.SendId, &ws.Chat{
		ConversationId: data.ConversationId,
		ChatType:       data.ChatType,
		SendTime:       data.SendTime,
		Msg: ws.Msg{
			MType:   data.MType,
			Content: data.Content,
		},
	}), rconn)
}

// group
//
//	@Description: 群聊发送，本质还是私聊，只是使用多协程并发发送
//	@param srv
//	@param data
//	@param recvId
//	@return error
func group(srv *websocket.Server, data *ws.Push) error {
	for _, id := range data.RecvIds {
		func(id string) {
			// 加入协程调度执行
			srv.Schedule(func() {
				single(srv, data, id)
			})
		}(id)
	}
	return nil
}
