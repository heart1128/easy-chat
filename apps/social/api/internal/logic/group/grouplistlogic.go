/*
 * @Author: heart1128 1020273485@qq.com
 * @Date: 2024-07-13 10:04:13
 * @LastEditors: heart1128 1020273485@qq.com
 * @LastEditTime: 2024-07-21 10:31:11
 * @FilePath: /easy-chat/apps/social/api/internal/logic/group/grouplistlogic.go
 * @Description:  learn
 */
package group

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/pkg/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户申群列表
func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListLogic) GroupList(req *types.GroupListRep) (resp *types.GroupListResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)
	groupList, err := l.svcCtx.Social.GroupList(l.ctx, &social.GroupListReq{
		UserId: uid,
	})

	if err != nil{
		return nil, err
	}

	var respList []*types.Groups
	copier.Copy(&respList, &groupList.List)

	return &types.GroupListResp{
		List: respList,
	}, nil
}
