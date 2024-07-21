package websocket

// 为websocket动态传参使用（struct结构体包含多个参数使用）

type ServerOptions func(opt *ServerOption)

type ServerOption struct {
	Authentication
	patten string
}

func newServerOptions(opts ...ServerOptions) ServerOption {
	o := ServerOption{
		Authentication: new(authentication),
		patten:         "/ws", // 路由
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

// WithServerPatten 设置路由参数
func WithServerPatten(patten string) ServerOptions {
	return func(opt *ServerOption) {
		opt.patten = patten
	}
}
