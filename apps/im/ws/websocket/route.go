package websocket

type Route struct {
	Method  string
	Handler HandlerFunc
}

// HandlerFunc 设置路由处理
type HandlerFunc func(srv *Server, conn *Conn, msg *Message)
