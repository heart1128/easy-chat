package svc

import (
	"easy-chat/apps/im/immodels"
	"easy-chat/apps/im/ws/websocket"
	"easy-chat/apps/social/rpc/socialclient"
	"easy-chat/apps/task/mq/internal/config"
	"easy-chat/pkg/constants"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
)

type ServiceContext struct {
	config.Config

	WsClient websocket.Client
	*redis.Redis

	socialclient.Social   // 社交rpc，查询群用户
	immodels.ChatLogModel // mongo数据model
	immodels.ConversationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:            c,
		Redis:             redis.MustNewRedis(c.Redisx),
		ChatLogModel:      immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
		ConversationModel: immodels.MustConversationModel(c.Mongo.Url, c.Mongo.Db),
		Social:            socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
	}

	// websocket client需要单独创建
	// 在上下文中获取超级token
	token, err := svc.GetSystemToken()
	if err != nil {
		panic(err)
	}

	// 设置token到header中，websocket的client给server就会有token
	header := http.Header{}
	header.Set("Authorization", token)
	svc.WsClient = websocket.NewClient(c.Ws.Host, websocket.WithClientHeader(header))

	return svc
}

// GetSystemToken
//
//	@Description: 在user rpc服务启动中，生成了一个固定的token存在了redis中作为root，这里取出kafka就可以一直通过授权了
//	@receiver svc
//	@return string
//	@return error
func (svc *ServiceContext) GetSystemToken() (string, error) {
	return svc.Redis.Get(constants.REDIS_SYSTEM_ROOT_TOKEN)
}
