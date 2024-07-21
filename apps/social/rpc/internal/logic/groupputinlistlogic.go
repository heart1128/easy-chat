/*
 * @Author: heart1128 1020273485@qq.com
 * @Date: 2024-07-08 21:35:00
 * @LastEditors: heart1128 1020273485@qq.com
 * @LastEditTime: 2024-07-21 10:17:45
 * @FilePath: /easy-chat/apps/social/rpc/internal/logic/groupputinlistlogic.go
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

type GroupPutinListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutinListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutinListLogic {
	return &GroupPutinListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
 * @description: 查找群申请列表
 * @param {*social.GroupPutinListReq} in
 * @return {*}
 */
func (l *GroupPutinListLogic) GroupPutinList(in *social.GroupPutinListReq) (*social.GroupPutinListResp, error) {
	// todo: add your logic here and delete this line

	// 查找未处理的请求
	groupReqs, err := l.svcCtx.GroupRequestsModel.ListNoHandler(l.ctx, in.GroupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "list group req err %v req %v", err, in.GroupId)
	}

	var respList []*social.GroupRequests
	copier.Copy(&respList, &groupReqs)

	return &social.GroupPutinListResp{
		List: respList,
	}, nil
}
