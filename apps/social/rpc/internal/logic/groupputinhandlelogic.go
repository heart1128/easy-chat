/*
 * @Author: heart1128 1020273485@qq.com
 * @Date: 2024-07-08 21:35:00
 * @LastEditors: heart1128 1020273485@qq.com
 * @LastEditTime: 2024-07-20 22:42:02
 * @FilePath: /easy-chat/apps/social/rpc/internal/logic/groupputinhandlelogic.go
 * @Description:  learn
 */
package logic

import (
	"context"
	"database/sql"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/apps/social/socialmodels"
	"easy-chat/pkg/constants"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	ErrGroupReqBeforePass   = xerr.NewMsg("请求已通过")
	ErrGroupReqBeforeRefuse = xerr.NewMsg("请求已拒绝")
)

type GroupPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutInHandleLogic {
	return &GroupPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupPutInHandleLogic) GroupPutInHandle(in *social.GroupPutInHandleReq) (*social.GroupPutInHandleResp, error) {
	// todo: add your logic here and delete this line

	// 1. 查找请求
	groupReq, err := l.svcCtx.GroupRequestsModel.FindOne(l.ctx, int64(in.GroupReqId))
	if err != nil{
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group req err %v, req %v", err, in.GroupReqId)
	}

	// 2. 取出数据，如果是已经拒绝或者通过的，直接返回
	switch constants.HandlerResult(groupReq.HandleResult.Int64){
	case constants.PassHandlerResult:
		return nil, errors.WithStack(ErrGroupReqBeforePass)
	case constants.RefuseHandlerResult:
		return nil, errors.WithStack(ErrGroupReqBeforeRefuse)
	}

	// 3. 先更新群请求结果为通过，请求的成员通过事务插入到群成员表
	err = l.svcCtx.GroupRequestsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		
		// 更新群请求结果
		if err := l.svcCtx.GroupRequestsModel.Update(l.ctx, session, groupReq); err != nil{
			return errors.Wrapf(xerr.NewDBErr(), "update group req err %v req %v", err, groupReq)
		}

		if constants.HandlerResult(groupReq.HandleResult.Int64) != constants.PassHandlerResult{
			return nil
		}

		// 插入到群成员中
		groupMember := &socialmodels.GroupMembers{
			GroupId:     groupReq.GroupId,
			UserId:      groupReq.ReqId,
			RoleLevel:   int64(constants.AtLargeGroupRoleLevel),
			OperatorUid: sql.NullString{String: in.HandleUid}, // sql.NullString防止字段为空
		}

		_, err = l.svcCtx.GroupMembersModel.Insert(l.ctx, session, groupMember)
		if err != nil {
			return errors.Wrapf(xerr.NewDBErr(), "insert friend err %v req %v", err, groupMember)
		}

		return nil
	})


	return &social.GroupPutInHandleResp{}, err
}
