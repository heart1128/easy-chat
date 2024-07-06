package logic

import (
	"context"
	"easy-chat/apps/user/models"
	"github.com/jinzhu/copier"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {

	var (
		userEntirys []*models.Users
		err         error
	)

	// 1. 首先根据手机号查，手机号没有就根据昵称或者id查
	if in.Phone != "" {
		userEntiry, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
		if err == nil {
			userEntirys = append(userEntirys, userEntiry)
		}
	} else if in.Name != "" {
		userEntirys, err = l.svcCtx.UsersModel.ListByName(l.ctx, in.Name)
	} else if len(in.Ids) != 0 {
		userEntirys, err = l.svcCtx.UsersModel.ListByIds(l.ctx, in.Ids)
	}

	if err != nil {
		return nil, err
	}

	var resp []*user.UserEntity
	copier.Copy(&resp, userEntirys)

	// user是proto的类
	return &user.FindUserResp{
		User: resp,
	}, nil
}
