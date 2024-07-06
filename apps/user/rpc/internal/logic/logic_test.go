package logic

import (
	"easy-chat/apps/user/rpc/internal/config"
	"easy-chat/apps/user/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"path/filepath"
)

var svcCtx *svc.ServiceContext

// 初始化使用
func init() {
	var c config.Config
	// 加载配置
	conf.MustLoad(filepath.Join("../../etc/dev/user.yaml"), &c)
	svcCtx = svc.NewServiceContext(c)
}
