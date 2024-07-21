package websocket

import (
	"fmt"
	"net/http"
	"time"
)

// websocket连接鉴权

type Authentication interface {
	Auth(w http.ResponseWriter, r *http.Request) bool
	UserId(r *http.Request) string
}

// 实现接口
type authentication struct{}

func (*authentication) Auth(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func (*authentication) UserId(r *http.Request) string {
	// 1. 如果http的请求参数中带了userId
	query := r.URL.Query()
	if query != nil && query["userId"] != nil {
		return fmt.Sprintf("%v", query["userId"])
	}

	return fmt.Sprintf("%v", time.Now().UnixMilli())
}
