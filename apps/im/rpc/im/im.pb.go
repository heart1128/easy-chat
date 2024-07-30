// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: apps/im/rpc/im.proto

package im

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChatLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConversationId string `protobuf:"bytes,2,opt,name=conversationId,proto3" json:"conversationId,omitempty"`
	SendId         string `protobuf:"bytes,3,opt,name=sendId,proto3" json:"sendId,omitempty"`
	RecvId         string `protobuf:"bytes,4,opt,name=recvId,proto3" json:"recvId,omitempty"`
	MsgType        int32  `protobuf:"varint,5,opt,name=msgType,proto3" json:"msgType,omitempty"`
	MsgContent     string `protobuf:"bytes,6,opt,name=msgContent,proto3" json:"msgContent,omitempty"`
	ChatType       int32  `protobuf:"varint,7,opt,name=chatType,proto3" json:"chatType,omitempty"`
	SendTime       int64  `protobuf:"varint,8,opt,name=SendTime,proto3" json:"SendTime,omitempty"`
	ReadRecords    []byte `protobuf:"bytes,9,opt,name=readRecords,proto3" json:"readRecords,omitempty"`
}

func (x *ChatLog) Reset() {
	*x = ChatLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatLog) ProtoMessage() {}

func (x *ChatLog) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatLog.ProtoReflect.Descriptor instead.
func (*ChatLog) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{0}
}

func (x *ChatLog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChatLog) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

func (x *ChatLog) GetSendId() string {
	if x != nil {
		return x.SendId
	}
	return ""
}

func (x *ChatLog) GetRecvId() string {
	if x != nil {
		return x.RecvId
	}
	return ""
}

func (x *ChatLog) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *ChatLog) GetMsgContent() string {
	if x != nil {
		return x.MsgContent
	}
	return ""
}

func (x *ChatLog) GetChatType() int32 {
	if x != nil {
		return x.ChatType
	}
	return 0
}

func (x *ChatLog) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *ChatLog) GetReadRecords() []byte {
	if x != nil {
		return x.ReadRecords
	}
	return nil
}

type Conversation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversationId string `protobuf:"bytes,1,opt,name=conversationId,proto3" json:"conversationId,omitempty"`
	ChatType       int32  `protobuf:"varint,2,opt,name=chatType,proto3" json:"chatType,omitempty"`
	TargetId       string `protobuf:"bytes,3,opt,name=targetId,proto3" json:"targetId,omitempty"`
	IsShow         bool   `protobuf:"varint,4,opt,name=isShow,proto3" json:"isShow,omitempty"`
	Seq            int64  `protobuf:"varint,5,opt,name=seq,proto3" json:"seq,omitempty"`
	// 总消息数
	Total int32 `protobuf:"varint,6,opt,name=total,proto3" json:"total,omitempty"`
	// 未读消息数
	ToRead int32 `protobuf:"varint,7,opt,name=toRead,proto3" json:"toRead,omitempty"`
	// 已读消息
	Read int32    `protobuf:"varint,9,opt,name=Read,proto3" json:"Read,omitempty"`
	Msg  *ChatLog `protobuf:"bytes,8,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Conversation) Reset() {
	*x = Conversation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conversation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conversation) ProtoMessage() {}

func (x *Conversation) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conversation.ProtoReflect.Descriptor instead.
func (*Conversation) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{1}
}

func (x *Conversation) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

func (x *Conversation) GetChatType() int32 {
	if x != nil {
		return x.ChatType
	}
	return 0
}

func (x *Conversation) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *Conversation) GetIsShow() bool {
	if x != nil {
		return x.IsShow
	}
	return false
}

func (x *Conversation) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Conversation) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Conversation) GetToRead() int32 {
	if x != nil {
		return x.ToRead
	}
	return 0
}

func (x *Conversation) GetRead() int32 {
	if x != nil {
		return x.Read
	}
	return 0
}

func (x *Conversation) GetMsg() *ChatLog {
	if x != nil {
		return x.Msg
	}
	return nil
}

type GetConversationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetConversationsReq) Reset() {
	*x = GetConversationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConversationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConversationsReq) ProtoMessage() {}

func (x *GetConversationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConversationsReq.ProtoReflect.Descriptor instead.
func (*GetConversationsReq) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{2}
}

func (x *GetConversationsReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetConversationsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversationList map[string]*Conversation `protobuf:"bytes,2,rep,name=conversationList,proto3" json:"conversationList,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetConversationsResp) Reset() {
	*x = GetConversationsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConversationsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConversationsResp) ProtoMessage() {}

func (x *GetConversationsResp) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConversationsResp.ProtoReflect.Descriptor instead.
func (*GetConversationsResp) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{3}
}

func (x *GetConversationsResp) GetConversationList() map[string]*Conversation {
	if x != nil {
		return x.ConversationList
	}
	return nil
}

type PutConversationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId           string                   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ConversationList map[string]*Conversation `protobuf:"bytes,3,rep,name=conversationList,proto3" json:"conversationList,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PutConversationsReq) Reset() {
	*x = PutConversationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutConversationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutConversationsReq) ProtoMessage() {}

func (x *PutConversationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutConversationsReq.ProtoReflect.Descriptor instead.
func (*PutConversationsReq) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{4}
}

func (x *PutConversationsReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PutConversationsReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PutConversationsReq) GetConversationList() map[string]*Conversation {
	if x != nil {
		return x.ConversationList
	}
	return nil
}

type PutConversationsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutConversationsResp) Reset() {
	*x = PutConversationsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutConversationsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutConversationsResp) ProtoMessage() {}

func (x *PutConversationsResp) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutConversationsResp.ProtoReflect.Descriptor instead.
func (*PutConversationsResp) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{5}
}

type GetChatLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversationId string `protobuf:"bytes,1,opt,name=conversationId,proto3" json:"conversationId,omitempty"`
	StartSendTime  int64  `protobuf:"varint,2,opt,name=startSendTime,proto3" json:"startSendTime,omitempty"`
	EndSendTime    int64  `protobuf:"varint,3,opt,name=endSendTime,proto3" json:"endSendTime,omitempty"`
	Count          int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	MsgId          string `protobuf:"bytes,5,opt,name=msgId,proto3" json:"msgId,omitempty"`
}

func (x *GetChatLogReq) Reset() {
	*x = GetChatLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatLogReq) ProtoMessage() {}

func (x *GetChatLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatLogReq.ProtoReflect.Descriptor instead.
func (*GetChatLogReq) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{6}
}

func (x *GetChatLogReq) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

func (x *GetChatLogReq) GetStartSendTime() int64 {
	if x != nil {
		return x.StartSendTime
	}
	return 0
}

func (x *GetChatLogReq) GetEndSendTime() int64 {
	if x != nil {
		return x.EndSendTime
	}
	return 0
}

func (x *GetChatLogReq) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetChatLogReq) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

type GetChatLogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*ChatLog `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *GetChatLogResp) Reset() {
	*x = GetChatLogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatLogResp) ProtoMessage() {}

func (x *GetChatLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatLogResp.ProtoReflect.Descriptor instead.
func (*GetChatLogResp) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{7}
}

func (x *GetChatLogResp) GetList() []*ChatLog {
	if x != nil {
		return x.List
	}
	return nil
}

type SetUpUserConversationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendId   string `protobuf:"bytes,1,opt,name=SendId,proto3" json:"SendId,omitempty"`
	RecvId   string `protobuf:"bytes,2,opt,name=recvId,proto3" json:"recvId,omitempty"`
	ChatType int32  `protobuf:"varint,3,opt,name=chatType,proto3" json:"chatType,omitempty"`
}

func (x *SetUpUserConversationReq) Reset() {
	*x = SetUpUserConversationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUpUserConversationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUpUserConversationReq) ProtoMessage() {}

func (x *SetUpUserConversationReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUpUserConversationReq.ProtoReflect.Descriptor instead.
func (*SetUpUserConversationReq) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{8}
}

func (x *SetUpUserConversationReq) GetSendId() string {
	if x != nil {
		return x.SendId
	}
	return ""
}

func (x *SetUpUserConversationReq) GetRecvId() string {
	if x != nil {
		return x.RecvId
	}
	return ""
}

func (x *SetUpUserConversationReq) GetChatType() int32 {
	if x != nil {
		return x.ChatType
	}
	return 0
}

type SetUpUserConversationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetUpUserConversationResp) Reset() {
	*x = SetUpUserConversationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUpUserConversationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUpUserConversationResp) ProtoMessage() {}

func (x *SetUpUserConversationResp) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUpUserConversationResp.ProtoReflect.Descriptor instead.
func (*SetUpUserConversationResp) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{9}
}

type CreateGroupConversationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId  string `protobuf:"bytes,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	CreateId string `protobuf:"bytes,2,opt,name=CreateId,proto3" json:"CreateId,omitempty"`
}

func (x *CreateGroupConversationReq) Reset() {
	*x = CreateGroupConversationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupConversationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupConversationReq) ProtoMessage() {}

func (x *CreateGroupConversationReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupConversationReq.ProtoReflect.Descriptor instead.
func (*CreateGroupConversationReq) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{10}
}

func (x *CreateGroupConversationReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CreateGroupConversationReq) GetCreateId() string {
	if x != nil {
		return x.CreateId
	}
	return ""
}

type CreateGroupConversationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateGroupConversationResp) Reset() {
	*x = CreateGroupConversationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_im_rpc_im_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupConversationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupConversationResp) ProtoMessage() {}

func (x *CreateGroupConversationResp) ProtoReflect() protoreflect.Message {
	mi := &file_apps_im_rpc_im_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupConversationResp.ProtoReflect.Descriptor instead.
func (*CreateGroupConversationResp) Descriptor() ([]byte, []int) {
	return file_apps_im_rpc_im_proto_rawDescGZIP(), []int{11}
}

var File_apps_im_rpc_im_proto protoreflect.FileDescriptor

var file_apps_im_rpc_im_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x69, 0x6d, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6d, 0x22, 0x85, 0x02, 0x0a, 0x07, 0x43,
	0x68, 0x61, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x76, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x63, 0x76, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x73,
	0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x22, 0xf9, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63,
	0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x53, 0x68, 0x6f, 0x77, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x65, 0x71, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x52, 0x65, 0x61, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x6f, 0x52, 0x65, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x52,
	0x65, 0x61, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x1d, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69,
	0x6d, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2d,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xc9, 0x01,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x69, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x1a, 0x55, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69,
	0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xef, 0x01, 0x0a, 0x13, 0x50, 0x75,
	0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x59, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6d, 0x2e, 0x50, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x55, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x69, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x16, 0x0a, 0x14, 0x50,
	0x75, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x22, 0xab, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x73, 0x67, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x22, 0x31, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x66, 0x0a, 0x18, 0x53, 0x65, 0x74, 0x55, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x76,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x63, 0x76, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x1b, 0x0a, 0x19,
	0x53, 0x65, 0x74, 0x55, 0x70, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x52, 0x0a, 0x1a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x1d, 0x0a,
	0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x32, 0xf9, 0x02, 0x0a,
	0x02, 0x49, 0x6d, 0x12, 0x33, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x6f,
	0x67, 0x12, 0x11, 0x2e, 0x69, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x69, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x55,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x2e, 0x69, 0x6d, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x69, 0x6d, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x70, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x17, 0x2e, 0x69, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x69, 0x6d,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x2e, 0x69, 0x6d, 0x2e, 0x50,
	0x75, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x69, 0x6d, 0x2e, 0x50, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5a, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x69, 0x6d, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x69, 0x6d, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x69, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_im_rpc_im_proto_rawDescOnce sync.Once
	file_apps_im_rpc_im_proto_rawDescData = file_apps_im_rpc_im_proto_rawDesc
)

func file_apps_im_rpc_im_proto_rawDescGZIP() []byte {
	file_apps_im_rpc_im_proto_rawDescOnce.Do(func() {
		file_apps_im_rpc_im_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_im_rpc_im_proto_rawDescData)
	})
	return file_apps_im_rpc_im_proto_rawDescData
}

var file_apps_im_rpc_im_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_apps_im_rpc_im_proto_goTypes = []interface{}{
	(*ChatLog)(nil),                     // 0: im.ChatLog
	(*Conversation)(nil),                // 1: im.Conversation
	(*GetConversationsReq)(nil),         // 2: im.GetConversationsReq
	(*GetConversationsResp)(nil),        // 3: im.GetConversationsResp
	(*PutConversationsReq)(nil),         // 4: im.PutConversationsReq
	(*PutConversationsResp)(nil),        // 5: im.PutConversationsResp
	(*GetChatLogReq)(nil),               // 6: im.GetChatLogReq
	(*GetChatLogResp)(nil),              // 7: im.GetChatLogResp
	(*SetUpUserConversationReq)(nil),    // 8: im.SetUpUserConversationReq
	(*SetUpUserConversationResp)(nil),   // 9: im.SetUpUserConversationResp
	(*CreateGroupConversationReq)(nil),  // 10: im.CreateGroupConversationReq
	(*CreateGroupConversationResp)(nil), // 11: im.CreateGroupConversationResp
	nil,                                 // 12: im.GetConversationsResp.ConversationListEntry
	nil,                                 // 13: im.PutConversationsReq.ConversationListEntry
}
var file_apps_im_rpc_im_proto_depIdxs = []int32{
	0,  // 0: im.Conversation.msg:type_name -> im.ChatLog
	12, // 1: im.GetConversationsResp.conversationList:type_name -> im.GetConversationsResp.ConversationListEntry
	13, // 2: im.PutConversationsReq.conversationList:type_name -> im.PutConversationsReq.ConversationListEntry
	0,  // 3: im.GetChatLogResp.List:type_name -> im.ChatLog
	1,  // 4: im.GetConversationsResp.ConversationListEntry.value:type_name -> im.Conversation
	1,  // 5: im.PutConversationsReq.ConversationListEntry.value:type_name -> im.Conversation
	6,  // 6: im.Im.GetChatLog:input_type -> im.GetChatLogReq
	8,  // 7: im.Im.SetUpUserConversation:input_type -> im.SetUpUserConversationReq
	2,  // 8: im.Im.GetConversations:input_type -> im.GetConversationsReq
	4,  // 9: im.Im.PutConversations:input_type -> im.PutConversationsReq
	10, // 10: im.Im.CreateGroupConversation:input_type -> im.CreateGroupConversationReq
	7,  // 11: im.Im.GetChatLog:output_type -> im.GetChatLogResp
	9,  // 12: im.Im.SetUpUserConversation:output_type -> im.SetUpUserConversationResp
	3,  // 13: im.Im.GetConversations:output_type -> im.GetConversationsResp
	5,  // 14: im.Im.PutConversations:output_type -> im.PutConversationsResp
	11, // 15: im.Im.CreateGroupConversation:output_type -> im.CreateGroupConversationResp
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_apps_im_rpc_im_proto_init() }
func file_apps_im_rpc_im_proto_init() {
	if File_apps_im_rpc_im_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_im_rpc_im_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_im_rpc_im_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conversation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_im_rpc_im_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConversationsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_im_rpc_im_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConversationsResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_im_rpc_im_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutConversationsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_im_rpc_im_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutConversationsResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_im_rpc_im_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatLogReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_im_rpc_im_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatLogResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_im_rpc_im_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUpUserConversationReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_im_rpc_im_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUpUserConversationResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_im_rpc_im_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupConversationReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_im_rpc_im_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupConversationResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_im_rpc_im_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_im_rpc_im_proto_goTypes,
		DependencyIndexes: file_apps_im_rpc_im_proto_depIdxs,
		MessageInfos:      file_apps_im_rpc_im_proto_msgTypes,
	}.Build()
	File_apps_im_rpc_im_proto = out.File
	file_apps_im_rpc_im_proto_rawDesc = nil
	file_apps_im_rpc_im_proto_goTypes = nil
	file_apps_im_rpc_im_proto_depIdxs = nil
}
