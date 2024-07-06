package ctxdata

import "context"

func GetUId(ctx context.Context) string {
	// http的上下文传入，利用上下文传递值
	// 这里的上下文就是token.go里面设置的
	if u, ok := ctx.Value(Identify).(string); ok {
		return u
	}
	return ""
}
