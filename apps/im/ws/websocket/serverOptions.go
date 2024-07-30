package websocket

import "time"

// 为websocket动态传参使用（struct结构体包含多个参数使用）

type ServerOptions func(opt *ServerOption)

type ServerOption struct {
	Authentication
	ack        AckType       // ack应答方式
	ackTimeout time.Duration // ack有等待时间，超过处理就错误了
	patten     string

	maxConnectionIdle time.Duration

	concurrency int // 群聊websocket并发数
}

func newServerOptions(opts ...ServerOptions) ServerOption {
	o := ServerOption{
		Authentication:    new(authentication),
		ackTimeout:        defaultAckTimeout,
		maxConnectionIdle: defaultMaxConnectionIdle,
		patten:            "/ws", // 路由
		concurrency:       defaultConcurrentcy,
	}

	// 执行func(opt *ServerOption)，参数是o
	for _, opt := range opts {
		opt(&o)
	}

	return o
}

// WithServerAuthentication 设置鉴权参数
func WithServerAuthentication(auth Authentication) ServerOptions {
	return func(opt *ServerOption) {
		opt.Authentication = auth
	}
}

// WithServerAck 设置路由参数
func WithServerAck(ack AckType) ServerOptions {
	return func(opt *ServerOption) {
		opt.ack = ack
	}
}

// WithServerPatten 设置路由参数
func WithServerPatten(patten string) ServerOptions {
	return func(opt *ServerOption) {
		opt.patten = patten
	}
}

// WithServerMaxConnectionIdle 设置最大空闲时间
func WithServerMaxConnectionIdle(maxConnectionIdle time.Duration) ServerOptions {
	return func(opt *ServerOption) {
		if maxConnectionIdle > 0 {
			opt.maxConnectionIdle = maxConnectionIdle
		}
	}
}
