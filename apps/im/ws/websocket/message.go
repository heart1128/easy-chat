package websocket

type Message struct {
	Method string      `json:"method"`
	FormId string      `json:"formId"`
	Data   interface{} `json:"data"`
}

func NewMessage(fromId string, data interface{}) *Message {
	return &Message{
		FormId: fromId,
		Data:   data,
	}
}
