/*
 * @Author: heart1128 1020273485@qq.com
 * @Date: 2024-07-08 21:35:00
 * @LastEditors: heart1128 1020273485@qq.com
 * @LastEditTime: 2024-07-21 10:15:41
 * @FilePath: /easy-chat/apps/social/rpc/internal/logic/grouplistlogic.go
 * @Description:  learn
 */
package logic

import (
	"context"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
 * @description: 查找用户的群列表
 * @param {*social.GroupListReq} in
 * @return {*}
 */
func (l *GroupListLogic) GroupList(in *social.GroupListReq) (*social.GroupListResp, error) {
	// todo: add your logic here and delete this line

	// 1. 首先要判断请求的用户加入的群有几个
	// 2. 然后通过群id查找群

	userGroup, err := l.svcCtx.GroupMembersModel.ListByGroupId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "list group member err %v req %v", err, in.UserId)
	}
		// 空列表
	if len(userGroup) == 0{
		return &social.GroupListResp{}, nil
	}

	// 开始查找群
	ids := make([]string, 0, len(userGroup))
	for _,v := range userGroup{
		ids = append(ids, v.GroupId)
	}

	groups, err := l.svcCtx.GroupsModel.ListByGroupIds(l.ctx, ids)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "list group err %v req %v", err, ids)
	}

	var respList []*social.Groups
	copier.Copy(&respList, &groups)

	return &social.GroupListResp{
		List: respList,
	}, nil
}
