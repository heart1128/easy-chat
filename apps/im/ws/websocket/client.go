package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/url"
)

type Client interface {
	Close() error

	Send(v any) error
	Read(v any) error
}

type client struct {
	*websocket.Conn
	host string

	opt dailOption
}

func NewClient(host string, opts ...DailOptions) *client {
	opt := newDailOptions(opts...)

	c := client{
		Conn: nil,
		host: host,
		opt:  opt,
	}
	conn, err := c.dail()
	if err != nil {
		panic(err)
	}

	c.Conn = conn
	return &c
}

// dail
//
//	@Description: 创建一个websocket的dial连接，给客户端使用
//	@receiver c
//	@return *websocket.Conn
//	@return error
func (c *client) dail() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: c.host, Path: c.opt.pattern}
	// 这里传入了一个header，就是http的header，从里面获取jwt token
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), c.opt.header)
	return conn, err
}

// Send
//
//	@Description: 客户端发送消息，调用websocket连接的WriteMessage就行
//	@receiver c
//	@param v
//	@return error
func (c *client) Send(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = c.WriteMessage(websocket.TextMessage, data)
	if err == nil {
		return nil
	}

	// todo：有错误再增加一个重连发送（可能发送失败）
	conn, err := c.dail()
	if err != nil {
		panic(err)
	}
	c.Conn = conn

	return c.WriteMessage(websocket.TextMessage, data)
}

func (c *client) Read(v any) error {
	_, msg, err := c.Conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	return json.Unmarshal(msg, v) // 反序列化到v
}

func (c *client) Close() error {
	return c.Conn.Close()
}
