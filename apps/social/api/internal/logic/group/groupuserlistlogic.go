/*
 * @Author: heart1128 1020273485@qq.com
 * @Date: 2024-07-13 10:04:13
 * @LastEditors: heart1128 1020273485@qq.com
 * @LastEditTime: 2024-07-21 10:36:54
 * @FilePath: /easy-chat/apps/social/api/internal/logic/group/groupuserlistlogic.go
 * @Description:  learn
 */
package group

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"easy-chat/apps/social/rpc/socialclient"
	"easy-chat/apps/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 成员列表列表
func NewGroupUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUserListLogic {
	return &GroupUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUserListLogic) GroupUserList(req *types.GroupUserListReq) (resp *types.GroupUserListResp, err error) {
	// todo: add your logic here and delete this line
	groupUsers, err := l.svcCtx.Social.GroupUsers(l.ctx, &socialclient.GroupUsersReq{
		GroupId: req.GroupId,
	})

	if err != nil{
		return nil, err
	}

	// 还需要获取用户的信息
	uids := make([]string, 0, len(groupUsers.List))
	for _, v := range groupUsers.List {
		uids = append(uids, v.UserId)
	}

	// 获取用户信息
	userList, err := l.svcCtx.User.FindUser(l.ctx, &userclient.FindUserReq{Ids: uids})
	if err != nil {
		return nil, err
	}

	userRecords := make(map[string]*userclient.UserEntity, len(userList.User))
	for i, _ := range userList.User {
		userRecords[userList.User[i].Id] = userList.User[i]
	}

	respList := make([]*types.GroupMembers, 0, len(groupUsers.List))
	for _, v := range groupUsers.List {

		member := &types.GroupMembers{
			Id:        int64(v.Id),
			GroupId:   v.GroupId,
			UserId:    v.UserId,
			RoleLevel: int(v.RoleLevel),
		}
		if u, ok := userRecords[v.UserId]; ok {
			member.Nickname = u.Nickname
			member.UserAvatarUrl = u.Avatar
		}
		respList = append(respList, member)
	}

	return &types.GroupUserListResp{List: respList}, err
}
