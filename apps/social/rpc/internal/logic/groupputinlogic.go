/*
 * @Author: heart1128 1020273485@qq.com
 * @Date: 2024-07-08 21:35:00
 * @LastEditors: heart1128 1020273485@qq.com
 * @LastEditTime: 2024-07-21 10:05:24
 * @FilePath: /easy-chat/apps/social/rpc/internal/logic/groupputinlogic.go
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
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutinLogic {
	return &GroupPutinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 申请进群，加入到申请群
func (l *GroupPutinLogic) GroupPutin(in *social.GroupPutinReq) (*social.GroupPutinResp, error) {
	// todo: add your logic here and delete this line

	//  1. 普通用户申请 ： 如果群无验证直接进入
	//  2. 群成员邀请： 如果群无验证直接进入
	//  3. 群管理员/群创建者邀请：直接进入群

	var (
		inviteGroupMember *socialmodels.GroupMembers
		userGroupMember   *socialmodels.GroupMembers
		groupInfo         *socialmodels.Groups

		err error
	)

	// 首先判断申请者是不是在群里了（在群成员表中查找）
	userGroupMember, err = l.svcCtx.GroupMembersModel.FindByGroudIdAndUserId(l.ctx, in. ReqId, in.GroupId)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group member by groud id and  req id err %v, req %v, %v", err,
			in.GroupId, in.ReqId)
	}
	if userGroupMember != nil{  		// 找到了，说明在群里，不用处理
		return &social.GroupPutinResp{}, nil
	}

	// 查找加群请求
	groupReq, err := l.svcCtx.GroupRequestsModel.FindByGroupIdAndReqId(l.ctx, in.GroupId, in.ReqId)
	if err != nil && err != socialmodels.ErrNotFound{
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group req by groud id and user id err %v, req %v, %v", err,
			in.GroupId, in.ReqId)
	}
	if groupReq != nil {		// 请求已经存在，不用重复请求
		return &social.GroupPutinResp{}, nil
	}

	// 组装请求数据，插入到请求加群表中
	groupReq = &socialmodels.GroupRequests{
		ReqId:         in.ReqId,
		GroupId:       in.GroupId,
		ReqMsg:        sql.NullString{
			String: in.ReqMsg,
			Valid: true,
		},
		ReqTime:       sql.NullTime{
			Time: time.Unix(in.ReqTime, 0),
			Valid: true,
		},
		JoinSource:    sql.NullInt64{
			Int64: int64(in.JoinSource),
			Valid: true,
		},
		InviterUserId: sql.NullString{
			String: in.InviterUid, 
			Valid:  true,
		},

		HandleResult:  sql.NullInt64{
			Int64: int64(constants.NoHandlerResult),
			Valid: true,
		},
	}	

	// 加入群成员数据库
	createGroupMember := func(){
		if err != nil{
			return
		}

		err = l.createGroupMember(in)
	}

	// 查找群
	groupInfo, err = l.svcCtx.GroupsModel.FindOne(l.ctx, in.GroupId)
	if err != nil{
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group by groud id err %v, req %v", err, in.GroupId)
	}

	// 是否需要验证
	if !groupInfo.IsVerify{
		// 不需要，直接加入群成员
		defer createGroupMember()

		// 更新请求为通过
		groupReq.HandleResult = sql.NullInt64{
			Int64: int64(constants.PassHandlerResult),
			Valid: true,
		}

		// 群申请插入（因为这里是不需要验证的，群存在直接就加入了）
		return l.createGroupReq(groupReq, true)
	}

	// 需要验证
		// 自己申请的
	if constants.GroupJoinSource(in.JoinSource) == constants.PutInGroupJoinSource{
		// 插入申请表，等待审批
		return l.createGroupReq(groupReq, false)
	}

		// 被邀请的
	inviteGroupMember, err = l.svcCtx.GroupMembersModel.FindByGroudIdAndUserId(l.ctx, in.InviterUid, in.GroupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group member by groud id and user id err %v, req %v",
			in.InviterUid, in.GroupId)
	}
		// 邀请者是管理员或者群主，直接通过
	if constants.GroupRoleLevel(inviteGroupMember.RoleLevel) == constants.CreatorGroupRoleLevel ||
		constants.GroupRoleLevel(inviteGroupMember.RoleLevel) == constants.ManagerGroupRoleLevel{
			defer createGroupMember()

			groupReq.HandleResult = sql.NullInt64{
				Int64: int64(constants.PassHandlerResult),
				Valid: true,
			}

			groupReq.HandleUserId = sql.NullString{
				String: in.InviterUid,
				Valid: true,
			}

			return l.createGroupReq(groupReq, true)
	}

	return l.createGroupReq(groupReq, false)
}

/**
 * @description: 在前面验证通过之后，创建一个加群申请，插入数据库
 * @param {*socialmodels.GroupRequests} groupReq
 * @param {bool} isPass
 * @return {*}
 */
func (l *GroupPutinLogic)createGroupReq(groupReq *socialmodels.GroupRequests, isPass bool)(*social.GroupPutinResp, error) {
	
	_, err := l.svcCtx.GroupRequestsModel.Insert(l.ctx, groupReq)
	if err != nil{
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert group req err %v req %v", err, groupReq)
	}
	if isPass{
		return &social.GroupPutinResp{GroupId: groupReq.GroupId}, nil
	}

	return &social.GroupPutinResp{}, nil
}

/**
 * @description: 插入到群成员数据表中
 * @param {*social.GroupPutinReq} in
 * @return {*}
 */
func (l *GroupPutinLogic) createGroupMember(in *social.GroupPutinReq) error {
	groupMember := &socialmodels.GroupMembers{
		GroupId:     in.GroupId,
		UserId:      in.ReqId,
		RoleLevel:   int64(constants.AtLargeGroupRoleLevel),
		OperatorUid: sql.NullString{
			String: in.InviterUid,
			Valid: true,
		},
	}

	_, err := l.svcCtx.GroupMembersModel.Insert(l.ctx, nil, groupMember)
	if err != nil{
		return errors.Wrapf(xerr.NewDBErr(), "insert group member err %v req %v", err, groupMember)
	}

	return nil
}