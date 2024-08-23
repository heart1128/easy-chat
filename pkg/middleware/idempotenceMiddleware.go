package middleware

import (
	"easy-chat/pkg/interceptor"
	"net/http"
)

type IdempotenceMiddleware struct {
}

// Handler
//
//	@Description: http处理中间件， 设置幂等的请求id，在上下文中
//	@receiver m
//	@param next
//	@return http.HandlerFunc
func (m *IdempotenceMiddleware) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(interceptor.ContextWithVal(r.Context()))

		next(w, r)
	}
}

func NewIdempotenceMiddleware() *IdempotenceMiddleware {
	return &IdempotenceMiddleware{}
}
