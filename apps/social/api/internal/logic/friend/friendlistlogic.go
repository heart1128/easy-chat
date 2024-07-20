package friend

import (
	"context"
	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"easy-chat/apps/social/rpc/socialclient"
	"easy-chat/apps/user/rpc/userclient"
	"easy-chat/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友列表
func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	// todo: add your logic here and delete this line

	// 1. 获取好友列表
	uid := ctxdata.GetUId(l.ctx)

	friends, err := l.svcCtx.FriendList(l.ctx, &socialclient.FriendListReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	if len(friends.List) == 0 {
		return &types.FriendListResp{}, nil
	}

	// 2. 根据好友id获取好友 信息
	uids := make([]string, 0, len(friends.List))
	for _, i := range friends.List {
		uids = append(uids, i.FriendUid)
	}

	// 3. 根据uids查询所有好友信息
	users, err := l.svcCtx.FindUser(l.ctx, &userclient.FindUserReq{
		Ids: uids,
	})

	if err != nil {
		return &types.FriendListResp{}, err
	}

	// 因为查询到的users和查询的uids不是对应的，用一个map记录进行查找
	userRecords := make(map[string]*userclient.UserEntity, len(users.User))
	for i, _ := range friends.List {
		userRecords[users.User[i].Id] = users.User[i]
	}

	// 用好友的uid在map中查找，找到就是对应的
	respList := make([]*types.Friends, 0, len(friends.List))
	for _, v := range friends.List {
		friend := &types.Friends{
			Id:        v.Id,
			FriendUid: v.FriendUid,
		}

		if u, ok := userRecords[v.FriendUid]; ok {
			friend.Nickname = u.Nickname
			friend.Avatar = u.Avatar
		}

		respList = append(respList, friend)
	}

	return &types.FriendListResp{
		List: respList,
	}, nil
}
