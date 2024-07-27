package ws

import (
	"easy-chat/pkg/constants"
)

// mapstructure用于将通用的map[string]interface{}解码到对应的 Go 结构体中，或者执行相反的操作
// 因为message定义的data是interface{}，转换成json是 map[string]interface{}
type (
	// Msg 具体消息的结构体
	Msg struct {
		constants.MType `mapstructure:"mType"`
		Content         string `mapstructure:"content"`
	}

	// Chat 聊天会话的结构体
	Chat struct {
		ConversationId     string `mapstructure:"conversationId"` // 会话id
		constants.ChatType `mapstructure:"chatType"`
		SendId             string `mapstructure:"sendId"`
		RecvId             string `mapstructure:"recvId"`
		SendTime           int64  `mapstructure:"sendTime"`
		Msg                `mapstructure:"msg"`
	}

	Push struct {
		ConversationId     string                    `mapstructure:"conversationId"` // 会话id
		constants.ChatType `mapstructure:"chatType"` // 聊天类型（私聊，群聊）
		SendId             string                    `mapstructure:"sendId"` //发送者id
		RecvId             string                    `mapstructure:"recvId"`
		SendTime           int64                     `mapstructure:"sendTime"`

		constants.MType `mapstructure:"mType"` // 消息类型
		Content         string                 `mapstructure:"content"`
	}
)
