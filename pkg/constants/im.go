package constants

type MType int

const (
	TextMType MType = iota // 文本消息类型
)

type ChatType int

const (
	GroupChatType ChatType = iota + 1 // 群聊天
	SingleChatType
)

type ContentType int

const (
	ContentChatMsg ContentType = iota
	ContentMakeRead
)
