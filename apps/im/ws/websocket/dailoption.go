package websocket

import "net/http"

type DailOptions func(option *dailOption)

type dailOption struct {
	pattern string // 请求websocket的地址
	header  http.Header
}

func newDailOptions(opts ...DailOptions) dailOption {
	o := dailOption{
		pattern: "/ws",
		header:  nil,
	}

	// 要设置很多个选项对象的时候，把上面的设置都传出去
	for _, opt := range opts {
		opt(&o)
	}

	return o
}

// WithClientPatten 设置路由参数
func WithClientPatten(patten string) DailOptions {
	return func(opt *dailOption) {
		opt.pattern = patten
	}
}

func WithClientHeader(header http.Header) DailOptions {
	return func(opt *dailOption) {
		opt.header = header
	}
}
