package conversation

import (
	"context"
	"easy-chat/apps/im/ws/internal/logic"
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
	"easy-chat/apps/im/ws/ws"
	"easy-chat/pkg/constants"
	"github.com/mitchellh/mapstructure"
	"time"
)

// Chat
//
//	@Description: 聊天方法
//	@param svc
//	@return websocket.HandlerFunc
func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		var data ws.Chat
		// 解json，自动找对应的字段填充
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

		// 处理聊天类型，私聊，群聊等
		switch data.ChatType {
		case constants.SingleChatType: // 私聊
			err := logic.NewConversation(context.Background(), srv, svc).SingleChat(&data, conn.Uid)
			if err != nil {
				srv.Send(websocket.NewErrMessage(err), conn)
				return
			}

			srv.SendByUserId(websocket.NewMessage(conn.Uid, ws.Chat{
				ConversationId: data.ConversationId,
				ChatType:       data.ChatType,
				SendId:         conn.Uid,
				RecvId:         data.RecvId,
				SendTime:       time.Now().UnixMilli(),
				Msg:            data.Msg,
			}), data.RecvId)
		}

	}
}
