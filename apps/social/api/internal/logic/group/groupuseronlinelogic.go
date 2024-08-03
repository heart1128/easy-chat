package group

import (
	"context"
	"easy-chat/apps/social/rpc/socialclient"
	"easy-chat/pkg/constants"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUserOnlineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 群在线用户
func NewGroupUserOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUserOnlineLogic {
	return &GroupUserOnlineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GroupUserOnline
//
//	@Description: 查询群在线成员
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *GroupUserOnlineLogic) GroupUserOnline(req *types.GroupUserOnlineReq) (resp *types.GroupUserOnlineResp, err error) {
	// todo: add your logic here and delete this line

	// 查找群成员
	groupUsers, err := l.svcCtx.Social.GroupUsers(l.ctx, &socialclient.GroupUsersReq{
		GroupId: req.GroupId,
	})

	if err != nil || len(groupUsers.List) == 0 {
		return &types.GroupUserOnlineResp{}, err
	}

	// 2. 在线的redis中查找这些成员
	uids := make([]string, 0, len(groupUsers.List))
	for _, user := range groupUsers.List {
		uids = append(uids, user.UserId)
	}

	onlines, err := l.svcCtx.Redis.Hgetall(constants.REDIS_ONLINE_USER)
	if err != nil {
		return nil, err
	}

	resOnlineList := make(map[string]bool, len(onlines))
	// 查找每个uid在不在缓存中，在就在线，不在就离线
	for _, s := range uids {
		if _, ok := onlines[s]; ok {
			resOnlineList[s] = true
		} else {
			resOnlineList[s] = false
		}
	}

	return &types.GroupUserOnlineResp{OnlineList: resOnlineList}, nil

}
