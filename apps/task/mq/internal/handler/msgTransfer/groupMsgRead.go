package msgTransfer

import (
	"easy-chat/apps/im/ws/ws"
	"easy-chat/pkg/constants"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

type groupMsgRead struct {
	mu sync.Mutex

	push           *ws.Push      // 发送的
	pushCh         chan *ws.Push //在transfer和read之间推送的
	count          int
	pushTime       time.Time // 上次推送的时间
	done           chan struct{}
	conversationId string
}

func newGroupMsgRead(push *ws.Push, pushCh chan *ws.Push) *groupMsgRead {
	m := &groupMsgRead{
		push:           push,
		pushCh:         pushCh,
		count:          0,
		pushTime:       time.Now(),
		done:           make(chan struct{}),
		conversationId: push.ConversationId,
	}

	go m.transfer()

	return m
}

// mergerPush
//
//	@Description: 就是在push之后，合并到一个map中，一起推送
//	@receiver m
//	@param push
func (m *groupMsgRead) mergerPush(push *ws.Push) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.count++ // 累计
	for msgId, read := range push.ReadRecords {
		// 有相同的已读记录消息直接替换，反正都是已读
		m.push.ReadRecords[msgId] = read
	}
}

// transfer
//
//	@Description: 合并之后的检查（什么时候发送，1. 超时， 2. 超量），不断轮询的
//	@receiver m
func (m *groupMsgRead) transfer() {
	// 1. 超时发送
	timer := time.NewTimer(GroupMsgReadRecordDelayTime / 2)
	defer timer.Stop()

	for {
		select {
		case <-m.done:
			return
		case <-timer.C: // 定时器触发（时间到了，每隔一段时间检测一次）
			m.mu.Lock()

			pushTime := m.pushTime
			val := GroupMsgReadRecordDelayTime - time.Since(pushTime)
			push := m.push

			// 没有超时，也没有超量。或者没有消息，不用做处理
			if val > 0 && m.count < GroupMsgReadRecordDelayCount || push == nil {
				// 不做处理要重置定时器
				if val > 0 {
					timer.Reset(val)
				}
				m.mu.Unlock()
				continue
			}

			m.pushTime = time.Now()
			m.pushCh = nil
			timer.Reset(GroupMsgReadRecordDelayTime / 2)
			m.count = 0
			m.mu.Unlock()

			// 推送
			logx.Infof("超过合并的条件，开始合并推送%v", push)
			m.pushCh <- push
		default:
			// 上面只是判断了超时，在这里单独判断超量的情况
			// 2. 超量发送
			m.mu.Lock()

			if m.count >= GroupMsgReadRecordDelayCount {
				push := m.push
				m.pushTime = time.Now()
				m.pushCh = nil
				m.count = 0
				m.mu.Unlock()

				// 推送
				logx.Infof("超量，开始合并推送%v", push)
				m.pushCh <- push
			}

			// 没有达到任何条件
			if m.IsIdle() {
				m.mu.Unlock()
				// 通道给msgReadTransfer判断没有满足条件进行清理
				m.pushCh <- &ws.Push{
					ChatType:       constants.GroupChatType,
					ConversationId: m.conversationId,
				}
				continue
			}

			m.mu.Unlock()
			tempDelay := GroupMsgReadRecordDelayTime / 4
			if tempDelay > time.Second {
				tempDelay = time.Second
			}
			time.Sleep(tempDelay)
		}
	}

}

func (m *groupMsgRead) IsIdle() bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.isIdle()
}

func (m *groupMsgRead) isIdle() bool {
	pushTime := m.pushTime
	// 超时的时间 - 上一次推送的时间到现在的时间
	// 如果大于0表示还没有超时
	val := GroupMsgReadRecordDelayTime*2 - time.Since(pushTime)

	// 小于0超时，并且没有消息，就是一个空闲状态
	if val <= 0 && m.push == nil && m.count == 0 {
		return true
	}

	return false
}

func (m *groupMsgRead) clear() {
	select {
	case <-m.done: // 等待关闭
	default:
		close(m.done)
	}

	m.push = nil
}
