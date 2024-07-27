package svc

import (
	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/config"
	"easy-chat/pkg/constants"
	"easy-chat/pkg/ctxdata"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

// ServiceContext 配置服务上下文
type ServiceContext struct {
	Config config.Config
	*redis.Redis
	// user表模型
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,
		Redis:  redis.MustNewRedis(c.Redisx),
		// 添加user模型到上下文
		UsersModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}

// SetRootToken
//
//	@Description: 用一个指定不变的root作为key生成jwt，存在redis中持续使用，就是root token
//	@receiver svc
//	@return error
func (svc *ServiceContext) SetRootToken() error {
	// 生成jwt
	systemToken, err := ctxdata.GetJwtToken(svc.Config.Jwt.AccessSecret, time.Now().Unix(),
		9999999, constants.SYSTEM_ROOT_UID)
	if err != nil {
		return err
	}
	// 写入redis
	return svc.Redis.Set(constants.REDIS_SYSTEM_ROOT_TOKEN, systemToken)
}
