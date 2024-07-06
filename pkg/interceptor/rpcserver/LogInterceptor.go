package rpcserver

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	zerr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// rpc的日志事件拦截器，记录到中间件
// 在rpc的user.go中添加拦截器（Hook）
func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp any, err error) {

	// 给框架流程处理，得到response
	resp, err = handler(ctx, req)

	// 然后发送，因为这个拦截器是处理错误的，所以出错才会到下面去
	if err != nil {
		return resp, err
	}

	logx.WithContext(ctx).Errorf("【RPC SRV ERR】 %v", err)
	// 获取产生错误的原因
	causeErr := errors.Cause(err)
	// grpc有自己的错误处理，要转换，因为传进来是自定义的错误原因，自定义返回的事zerr.CideMsg结构体
	// 这里断言一下
	if e, ok := causeErr.(*zerr.CodeMsg); ok {
		err = status.Error(codes.Code(e.Code), e.Msg)
	}

	return resp, err
}
