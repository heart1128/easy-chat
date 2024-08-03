// Code generated by goctl. DO NOT EDIT.
package types

type FriendListReq struct {
}

type FriendListResp struct {
	List []*Friends `json:"list"`
}

type FriendPutInHandleReq struct {
	FriendReqId  int32 `json:"friend_req_id,omitempty"`
	HandleResult int32 `json:"handle_result,omitempty"` // 处理结果
}

type FriendPutInHandleResp struct {
}

type FriendPutInListReq struct {
}

type FriendPutInListResp struct {
	List []*FriendRequests `json:"list"`
}

type FriendPutInReq struct {
	ReqMsg  string `json:"req_msg,omitempty"`
	ReqTime int64  `json:"req_time,omitempty"`
	UserId  string `json:"user_uid"`
}

type FriendPutInResp struct {
}

type FriendRequests struct {
	Id           int64  `json:"id,omitempty"`
	UserId       string `json:"user_id,omitempty"`
	ReqUid       string `json:"req_uid,omitempty"`
	ReqMsg       string `json:"req_msg,omitempty"`
	ReqTime      int64  `json:"req_time,omitempty"`
	HandleResult int    `json:"handle_result,omitempty"`
	HandleMsg    string `json:"handle_msg,omitempty"`
}

type Friends struct {
	Id        int32  `json:"id,omitempty"`
	FriendUid string `json:"friend_uid,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Remark    string `json:"remark,omitempty"`
}

type FriendsOnlineReq struct {
}

type FriendsOnlineResp struct {
	OnlineList map[string]bool `json:"onlineList"`
}

type GroupCreateReq struct {
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
}

type GroupCreateResp struct {
}

type GroupListRep struct {
}

type GroupListResp struct {
	List []*Groups `json:"list,omitempty"`
}

type GroupMembers struct {
	Id            int64  `json:"id,omitempty"`
	GroupId       string `json:"group_id,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	Nickname      string `json:"nickname,omitempty"`
	UserAvatarUrl string `json:"user_avatar_url,omitempty"`
	RoleLevel     int    `json:"role_level,omitempty"`
	InviterUid    string `json:"inviter_uid,omitempty"`
	OperatorUid   string `json:"operator_uid,omitempty"`
}

type GroupPutInHandleRep struct {
	GroupReqId   int32  `json:"group_req_id,omitempty"`
	GroupId      string `json:"group_id,omitempty"`
	HandleResult int32  `json:"handle_result,omitempty"` // 处理结果
}

type GroupPutInHandleResp struct {
}

type GroupPutInListRep struct {
	GroupId string `json:"group_id,omitempty"`
}

type GroupPutInListResp struct {
	List []*GroupRequests `json:"list,omitempty"`
}

type GroupPutInRep struct {
	GroupId    string `json:"group_id,omitempty"`
	ReqMsg     string `json:"req_msg,omitempty"`
	ReqTime    int64  `json:"req_time,omitempty"`
	JoinSource int64  `json:"join_source,omitempty"`
}

type GroupPutInResp struct {
}

type GroupRequests struct {
	Id            int64  `json:"id,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	GroupId       string `json:"group_id,omitempty"`
	ReqMsg        string `json:"req_msg,omitempty"`
	ReqTime       int64  `json:"req_time,omitempty"`
	JoinSource    int64  `json:"join_source,omitempty"`
	InviterUserId string `json:"inviter_user_id,omitempty"`
	HandleUserId  string `json:"handle_user_id,omitempty"`
	HandleTime    int64  `json:"handle_time,omitempty"`
	HandleResult  int64  `json:"handle_result,omitempty"`
}

type GroupUserListReq struct {
	GroupId string `json:"group_id,omitempty"`
}

type GroupUserListResp struct {
	List []*GroupMembers `json:"List,omitempty"`
}

type GroupUserOnlineReq struct {
	GroupId string `json:"groupId"`
}

type GroupUserOnlineResp struct {
	OnlineList map[string]bool `json:"onlineList"`
}

type Groups struct {
	Id              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Icon            string `json:"icon,omitempty"`
	Status          int64  `json:"status,omitempty"`
	GroupType       int64  `json:"group_type,omitempty"`
	IsVerify        bool   `json:"is_verify,omitempty"`
	Notification    string `json:"notification,omitempty"`
	NotificationUid string `json:"notification_uid,omitempty"`
}
