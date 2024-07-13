package logic

import (
	"context"
	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/apps/social/socialmodels"
	"easy-chat/pkg/constants"
	"easy-chat/pkg/xerr"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrFriendReqBeforePass   = xerr.NewMsg("好友申请已经通过")
	ErrFriendReqBeforeRefuse = xerr.NewMsg("好友已经被拒绝")
)

type FriendPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FriendPutInHandle
//
//	@Description: 处理好友请求列表
//	@receiver l
//	@param in
//	@return *social.FriendPutInHandleResp
//	@return error
func (l *FriendPutInHandleLogic) FriendPutInHandle(in *social.FriendPutInHandleReq) (*social.FriendPutInHandleResp, error) {

	// 1. 获取好友申请记录
	friendReq, err := l.svcCtx.FriendRequestsModel.FindOne(l.ctx, int64(in.FriendReqId))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friends request by rid err %v req %v", err, in.FriendReqId)
	}

	// 2. 验证是否已处理
	// friendReq.HandleResult.Int64返回处理结果
	switch constants.HandlerResult(friendReq.HandleResult.Int64) {
	case constants.PassHandlerResult: // 已经通过
		return nil, errors.WithStack(ErrFriendReqBeforePass)
	case constants.RefuseHandlerResult: // 被拒绝
		return nil, errors.WithStack(ErrFriendReqBeforeRefuse)
	}

	// 处理结果为当前传入的值
	friendReq.HandleResult.Int64 = int64(in.HandleResult)

	// 3. 修改申请结果——> 1. 通过【建立好友关系记录】（用事务:修改的地方基本都要用事务，否则不一致）
	err = l.svcCtx.FriendRequestsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 更新好友请求，同意，事务包含
		if err := l.svcCtx.FriendRequestsModel.Update(l.ctx, session, friendReq); err != nil {
			return errors.Wrapf(xerr.NewDBErr(), "update friend request err %v, req %v", err, friendReq)
		}

		if constants.HandlerResult(in.HandleResult) == constants.PassHandlerResult {
			return nil
		}
		// 添加两条关系，分别建立关系数据，有冗余但是简单
		friends := []*socialmodels.Friends{
			{
				UserId:    friendReq.UserId,
				FriendUid: friendReq.ReqUid,
			},
			{
				UserId:    friendReq.ReqUid,
				FriendUid: friendReq.UserId,
			},
		}
		fmt.Println("开始插入数据")
		_, err = l.svcCtx.FriendsModel.Inserts(l.ctx, session, friends...)
		if err != nil {
			// 包装成zero自己的errors类型
			return errors.Wrapf(xerr.NewDBErr(), "friends insters err %v, req %v", err, friendReq)
		}

		return nil
	})
	return &social.FriendPutInHandleResp{}, nil
}
