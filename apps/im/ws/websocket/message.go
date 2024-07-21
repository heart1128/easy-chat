package websocket

type FrameType uint8

const (
	FrameData FrameType = 0x0
	FramePing FrameType = 0x1

	// FrameHeaders      FrameType = 0x1
	//FramePriority     FrameType = 0x2
	//FrameRSTStream    FrameType = 0x3
	//FrameSettings     FrameType = 0x4
	//FramePushPromise  FrameType = 0x5
	//FrameGoAway       FrameType = 0x7
	//FrameWindowUpdate FrameType = 0x8
	//FrameContinuation FrameType = 0x9
)

type Message struct {
	FrameType `json:"frameType"` // 消息类型，需要根据消息类型进行分别处理
	Method    string             `json:"method"`
	FormId    string             `json:"formId"`
	Data      interface{}        `json:"data"`
}

func NewMessage(fromId string, data interface{}) *Message {
	return &Message{
		FrameType: FrameData,
		FormId:    fromId,
		Data:      data,
	}
}
