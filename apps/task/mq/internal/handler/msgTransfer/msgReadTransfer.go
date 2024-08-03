package msgTransfer

import (
	"context"
	"easy-chat/apps/im/ws/ws"
	"easy-chat/apps/task/mq/internal/svc"
	"easy-chat/apps/task/mq/mq"
	"easy-chat/pkg/bitmap"
	"easy-chat/pkg/constants"
	"encoding/base64"
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"sync"
	"time"
)

var (
	GroupMsgReadRecordDelayTime  = time.Second
	GroupMsgReadRecordDelayCount = 10
)

const (
	GroupMsgReadHandlerAtTransfer = iota // 默认不不开启消息合并
	GroupMsgReadHandlerDelayTransfer
)

type MsgReadTransfer struct {
	*baseMsgTransfer

	cache.Cache

	mu        sync.Mutex
	groupMsgs map[string]*groupMsgRead // 群消息已读处理
	push      chan *ws.Push
}

func NewMsgReadTransfer(svc *svc.ServiceContext) kq.ConsumeHandler {
	m := &MsgReadTransfer{
		baseMsgTransfer: NewBaseMsgTransfer(svc),
		groupMsgs:       make(map[string]*groupMsgRead, 1),
		push:            make(chan *ws.Push, 1),
	}

	if svc.Config.MsgReadHandler.GroupMsgReadHandler != GroupMsgReadHandlerAtTransfer {
		if svc.Config.MsgReadHandler.GroupMsgReadRecordDelayCount > 0 {
			GroupMsgReadRecordDelayCount = svc.Config.MsgReadHandler.GroupMsgReadRecordDelayCount
		}

		if svc.Config.MsgReadHandler.GroupMsgReadRecordDelayTime > 0 {
			GroupMsgReadRecordDelayTime = time.Duration(svc.Config.MsgReadHandler.GroupMsgReadRecordDelayTime)
		}
	}

	go m.transfer()

	return m
}

// Consume 实现kafka消费者接口
func (m *MsgReadTransfer) Consume(ctx context.Context, key, value string) error {

	m.Info("MsgReadTransfer", value)

	var (
		data mq.MsgMarkRead // 消息类型结构体
	)

	// kafka保存的是json序列化数据，这里消费者拿到反序列化
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}

	// 业务处理： 更新用户对消息已读未读的记录
	readRecords, err := m.updateChatLogRead(ctx, &data)
	if err != nil {
		return err
	}
	// 已读记录

	push := &ws.Push{
		ConversationId: data.ConversationId,
		ChatType:       data.ChatType,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		ReadRecords:    readRecords,
		ContentType:    constants.ContentMakeRead,
	}

	switch data.ChatType {
	case constants.SingleChatType:
		// 直接推送
		m.push <- push
	case constants.GroupChatType:
		// 判断是否开启合并消息的处理， 没开启和私聊一样
		if m.svcCtx.Config.MsgReadHandler.GroupMsgReadHandler == GroupMsgReadHandlerAtTransfer {
			m.push <- push
		}

		m.mu.Lock()
		defer m.mu.Unlock()

		push.SendId = ""
		// 在已经保存的记录中存在，合并存在的请求
		if _, ok := m.groupMsgs[push.ConversationId]; ok {
			m.Infof("merge push %v", push.ConversationId)
			// 合并
			m.groupMsgs[push.ConversationId].mergerPush(push)
		} else {
			m.Infof("create push %v", push.ConversationId)
			// 没有就创建
			m.groupMsgs[push.ConversationId] = newGroupMsgRead(push, m.push)
		}
	}

	return nil
}

// updateChatLogRead
//
//	@Description: 更新已读未读。什么时候更新已读？
//	@receiver m
//	@param ctx
//	@param data
//	@return map[string]string  返回的是哪些用户已读
//	@return error
func (m *MsgReadTransfer) updateChatLogRead(ctx context.Context, data *mq.MsgMarkRead) (map[string]string, error) {

	res := make(map[string]string)

	chatLogs, err := m.svcCtx.ChatLogModel.ListByMsgIds(ctx, data.MsgIds)
	if err != nil {
		return nil, err
	}

	// 处理已读
	for _, chatLog := range chatLogs {
		switch chatLog.ChatType {
		case constants.SingleChatType:
			chatLog.ReadRecords = []byte{1}
		case constants.GroupChatType:
			newBitmap := bitmap.NewBitmap(0)
			readRecords := newBitmap.Load(chatLog.ReadRecords)
			readRecords.Set(data.SendId)
			chatLog.ReadRecords = readRecords.Export()
		}
		// 记录已读
		res[chatLog.ID.Hex()] = base64.StdEncoding.EncodeToString(chatLog.ReadRecords)
		// 更新到数据库
		err = m.svcCtx.ChatLogModel.UpdateMarkRead(ctx, chatLog.ID, chatLog.ReadRecords)
		if err != nil {
			return nil, err
		}
	}
	return res, err
}

// transfer
//
//	@Description: 使用协程和chan做出私聊的异步推送
//	@receiver m
func (m *MsgReadTransfer) transfer() {
	// 监听chan
	for push := range m.push {
		if push.RecvId != "" || len(push.RecvIds) > 0 {
			if err := m.Transfer(context.Background(), push); err != nil {
				m.Errorf("m transfer err %v", err, push)
			}
		}

		// 如果是私聊，不用管下面的合并
		if push.ChatType == constants.SingleChatType {
			continue
		}

		// 没开启群聊合并
		if m.svcCtx.Config.MsgReadHandler.GroupMsgReadHandler == GroupMsgReadHandlerAtTransfer {
			continue
		}

		// 清空数据
		m.mu.Lock()
		// 如果是存在这个消息，并且当前的消息是空闲的，就可以删除
		// 也就是之前有，但是后面不发了
		// 超时或者超累计数量就清空，
		if _, ok := m.groupMsgs[push.ConversationId]; ok && m.groupMsgs[push.ConversationId].IsIdle() {
			m.groupMsgs[push.ConversationId].clear()
			delete(m.groupMsgs, push.ConversationId)
		}

		m.mu.Unlock()
	}

}
