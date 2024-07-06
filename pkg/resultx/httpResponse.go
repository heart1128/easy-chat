package resultx

import (
	"context"
	"easy-chat/pkg/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	zrpcErr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc/status"
	"net/http"
)

// / 处理http响应
type Response struct {
	Code int         `json:"code"` // 状态
	Msg  string      `json:"msg"`  // 描述
	Data interface{} `json:"data"` // 数据
}

func Success(data interface{}) *Response {
	return &Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

func Fail(code int, err string) *Response {
	return &Response{
		Code: code,
		Msg:  err,
		Data: nil,
	}
}

// 这里是要在不改动api的handler下接管go-zero httpx的行为
func OkHandler(_ context.Context, v interface{}) any {
	return Success(v)
}

func ErrHandler(name string) func(ctx context.Context, err error) (int, any) {

	return func(ctx context.Context, err error) (int, any) {
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := xerr.ErrMsg(errcode)
		// 拿到原始错误对象
		causeErr := errors.Cause(err)
		// 断言是自定义错误（因为在xerr的errors.go中，返回一个CodeMsg对象）
		if e, ok := causeErr.(*zrpcErr.CodeMsg); ok {
			errcode = e.Code
			errmsg = e.Msg
		} else {
			// 是grpc的错误，是Log中间件设置进去的
			if gstatus, ok := status.FromError(causeErr); ok {
				errcode = int(gstatus.Code())
				errmsg = gstatus.Message()
			}
		}
		// 日志记录
		logx.WithContext(ctx).Errorf(" 【%s】err %v", name, err)

		return http.StatusBadRequest, Fail(errcode, errmsg)
	}
}
