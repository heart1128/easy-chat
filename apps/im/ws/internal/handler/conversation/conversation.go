package conversation

import (
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
	"easy-chat/apps/im/ws/ws"
	"easy-chat/apps/task/mq/mq"
	"easy-chat/pkg/constants"
	"easy-chat/pkg/wuid"
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
		// 没有会话id，分类讨论，有会话id，直接发送就行
		if data.ConversationId == "" {
			switch data.ChatType {
			case constants.SingleChatType: // 私聊
				data.ConversationId = wuid.CombineId(conn.Uid, data.RecvId)
			case constants.GroupChatType: // 群聊
				data.ConversationId = data.RecvId
			}
		}

		// push，作为product到kafka中
		err := svc.MsgChatTransferClient.Push(&mq.MsgChatTransfer{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			SendTime:       time.Now().UnixNano(),
			MType:          data.MType,
			Content:        data.Msg.Content,
		})
		if err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}
	}
}

// MarkRead
//
//	@Description: 已读未读的处理
//	@param svc
//	@return websocket.HandlerFunc
func MarkRead(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {

		var data ws.MarkRead
		// 解json，自动找对应的字段填充
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

		// 从kafka中使用websocket发送
		err := svc.MsgReadTransferClient.Push(&mq.MsgMarkRead{
			ChatType:       data.ChatType,
			ConversationId: data.ConversationId,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			MsgIds:         data.MsgIds,
		})
		if err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

	}
}
