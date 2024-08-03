package user

import (
	"context"
	"easy-chat/apps/user/api/internal/svc"
	"easy-chat/apps/user/api/internal/types"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/constants"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登入
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 调用rpc的server注册，rpc的逻辑已经完成了
	LoginResp, err := l.svcCtx.User.Login(l.ctx, &user.LoginReq{
		Phone:    req.Phone,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	var res types.LoginResp
	copier.Copy(&res, LoginResp)

	// 处理登录的业务，缓存登录用户
	// constants.REDIS_ONLINE_USER是redis的key, LoginResp.Id是哈希类型的key, value随便
	l.svcCtx.Redis.HsetCtx(l.ctx, constants.REDIS_ONLINE_USER, LoginResp.Id, "1")

	return &res, nil
}
