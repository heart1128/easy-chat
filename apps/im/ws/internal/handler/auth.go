package handler

import (
	"context"
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/pkg/ctxdata"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/token"
	"net/http"
)

// go zero定义了处理jwt的过程，需要传入处理的方法等内容

type JwtAuth struct {
	svc    *svc.ServiceContext
	parser *token.TokenParser
	logx.Logger
}

// Auth jwt鉴权
func (j JwtAuth) Auth(w http.ResponseWriter, r *http.Request) bool {
	// 1. 从http request的header中解析token
	tok, err := j.parser.ParseToken(r, j.svc.Config.JwtAuth.AccessSecret, "")
	if err != nil {
		j.Errorf("parse token err %v", err)
		return false
	}

	// token无效
	if !tok.Valid {
		return false
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}
	// 设置token到请求中，设置到上下文中
	*r = *r.WithContext(context.WithValue(r.Context(), ctxdata.Identify, claims[ctxdata.Identify]))

	return true
}

func (j JwtAuth) UserId(r *http.Request) string {
	return ctxdata.GetUId(r.Context())
}

func NewJwtAuth(svc *svc.ServiceContext) *JwtAuth {
	return &JwtAuth{
		svc:    svc,
		parser: token.NewTokenParser(),
		Logger: logx.WithContext(context.Background()),
	}
}
