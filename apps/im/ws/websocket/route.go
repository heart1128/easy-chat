package websocket

import "github.com/gorilla/websocket"

type Route struct {
	Method  string
	Handler HandlerFunc
}

// HandlerFunc 设置路由处理
type HandlerFunc func(srv *Server, conn *websocket.Conn, mes *Message)
