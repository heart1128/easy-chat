package msgTransfer

import (
	"context"
	"easy-chat/apps/im/immodels"
	"easy-chat/apps/im/ws/websocket"
	"easy-chat/apps/task/mq/internal/svc"
	"easy-chat/apps/task/mq/mq"
	"easy-chat/pkg/constants"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

// MsgChatTransfer
//
//	@Description: kafka消费者，只要实现指定消费者接口就行
type MsgChatTransfer struct {
	logx.Logger
	svc *svc.ServiceContext
}

func NewMsgChatTransfer(svc *svc.ServiceContext) *MsgChatTransfer {
	return &MsgChatTransfer{
		Logger: logx.WithContext(context.Background()),
		svc:    svc,
	}
}

// Consume
//
//	@Description: 实现zeromicro/go-queue/kq/queue.go 中的kafka接口，当订阅的topic有消息，就会通知这个消费者
//				消费者拿到之后，存入数据库，通知websocket，在推送给目标用户
//	@receiver m
//	@param ctx
//	@param key			kafka key
//	@param value
//	@return error
func (m *MsgChatTransfer) Consume(ctx context.Context, key, value string) error {
	fmt.Println("key ：", key, " value : ", value)

	var (
		data mq.MsgChatTransfer // 消息类型结构体
		c    = context.Background()
	)

	// kafka保存的是json序列化数据，这里消费者拿到反序列化
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		panic(err)
	}

	// 消费者拿到数据之后把数据保存到数据库
	if err := m.addChatLog(c, &data); err != nil {
		return err
	}

	// 使用websocket推送消息
	return m.svc.WsClient.Send(websocket.Message{
		FrameType: websocket.FrameData,
		Method:    "push",
		FormId:    constants.SYSTEM_ROOT_UID, // 当前系统角色
		Data:      data,
	})
}

// addChatLog
//
//	@Description: 组装一条消息到数据库中
//	@receiver m
//	@param data
//	@return error
func (m *MsgChatTransfer) addChatLog(ctx context.Context, data *mq.MsgChatTransfer) error {
	chatlog := immodels.ChatLog{
		ConversationId: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		ChatType:       data.ChatType,
		MsgFrom:        0,
		MsgType:        data.MType,
		MsgContent:     data.Content,
		SendTime:       data.SendTime,
	}
	// 插入数据库
	return m.svc.ChatLogModel.Insert(ctx, &chatlog)
}
