package mqclient

import (
	"context"
	"easy-chat/apps/task/mq/mq"
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
)

type MsgChatTransferClient interface {
	Push(msg *mq.MsgChatTransfer) error
}

type MsgReadTransferClient interface {
	Push(msg *mq.MsgMarkRead) error
}

type msgChatTransferClient struct {
	pusher *kq.Pusher // 生产者
}

// MsgReadTransfer
//
//	@Description: 对已读处理的客户端
type msgReadTransferClient struct {
	pusher *kq.Pusher
}

func NewMsgChatTransferClient(addr []string, topic string, opts ...kq.PushOption) MsgChatTransferClient {
	return &msgChatTransferClient{
		pusher: kq.NewPusher(addr, topic),
	}
}

func NewMsgReadTransferClient(addr []string, topic string, opts ...kq.PushOption) MsgReadTransferClient {
	return &msgReadTransferClient{
		pusher: kq.NewPusher(addr, topic),
	}
}

// Push
//
//	@Description: 消息生产者加入到kafka中
//	@receiver c
//	@param msg
//	@return error
func (c *msgChatTransferClient) Push(msg *mq.MsgChatTransfer) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return c.pusher.Push(context.Background(), string(body))
}

// Push
//
//	@Description: 实现已读未读，bitmap发送的websocket的客户端
//	@receiver c
//	@param msg
//	@return error
func (c *msgReadTransferClient) Push(msg *mq.MsgMarkRead) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return c.pusher.Push(context.Background(), string(body))
}
