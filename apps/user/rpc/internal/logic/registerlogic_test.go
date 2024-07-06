package logic

import (
	"context"
	"easy-chat/apps/user/rpc/user"
	"testing"
)

func TestRegisterLogic_Register(t *testing.T) {
	type args struct {
		in *user.RegisterReq
	}
	tests := []struct {
		name      string
		args      args
		wantPrint bool // 期待的值
		wantErr   bool
	}{
		{
			"1", args{in: &user.RegisterReq{
				Phone:    "15141882775",
				Nickname: "heart",
				Password: "12345",
				Avatar:   "png.jpg",
				Sex:      1,
			}}, true, false, // 预期返回注册成功
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantPrint {
				t.Log(tt.name, got)
			}
		})
	}
}
