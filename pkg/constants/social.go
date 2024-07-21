/*
 * @Author: heart1128 1020273485@qq.com
 * @Date: 2024-07-08 22:23:25
 * @LastEditors: heart1128 1020273485@qq.com
 * @LastEditTime: 2024-07-20 22:24:42
 * @FilePath: /easy-chat/pkg/constants/social.go
 * @Description:  learn
 */
package constants

// 模块的状态

// 处理结果 1. 未处理， 2、 处理 3、拒绝
type HandlerResult int

const (
	NoHandlerResult HandlerResult = iota + 1
	PassHandlerResult
	RefuseHandlerResult
	CancelHandlerResult
)

// 群等级 ， 1. 创建者（群主）， 2. 管理员  3.普通
type GroupRoleLevel int

const(
	CreatorGroupRoleLevel GroupRoleLevel = iota + 1
	ManagerGroupRoleLevel
	AtLargeGroupRoleLevel
)


// 进群的申请方式 1. 邀请 2. 申请
type GroupJoinSource int

const (
	InviteGroupJoinSource GroupJoinSource = iota + 1
	PutInGroupJoinSource
)