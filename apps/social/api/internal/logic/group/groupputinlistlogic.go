/*
 * @Author: heart1128 1020273485@qq.com
 * @Date: 2024-07-13 10:04:13
 * @LastEditors: heart1128 1020273485@qq.com
 * @LastEditTime: 2024-07-21 10:36:28
 * @FilePath: /easy-chat/apps/social/api/internal/logic/group/groupputinlistlogic.go
 * @Description:  learn
 */
package group

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"easy-chat/apps/social/rpc/socialclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutInListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请进群列表
func NewGroupPutInListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutInListLogic {
	return &GroupPutInListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPutInListLogic) GroupPutInList(req *types.GroupPutInListRep) (resp *types.GroupPutInListResp, err error) {
	// todo: add your logic here and delete this line

	list, err := l.svcCtx.Social.GroupPutinList(l.ctx, &socialclient.GroupPutinListReq{
		GroupId: req.GroupId,
	})

	if err != nil{
		return nil, err
	}

	var respList []*types.GroupRequests
	copier.Copy(&respList, list.List)

	return &types.GroupPutInListResp{List: respList}, nil
}
