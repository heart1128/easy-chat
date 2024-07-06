package svc

import (
	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 配置服务上下文
type ServiceContext struct {
	Config config.Config

	// user表模型
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		// 添加user模型到上下文
		UsersModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}
