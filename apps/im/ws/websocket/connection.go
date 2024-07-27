package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

// 加入心跳检测,因为websocket没有自带的心跳检测，所以还要在上层封装一个自定义的connection

type Conn struct {
	*websocket.Conn
	s                 *Server
	Uid               string
	idleMu            sync.Mutex
	idle              time.Time
	maxConnectionIdle time.Duration

	messageMu      sync.Mutex // ack机制锁
	readMessage    []*Message // 读消息的队列
	readMessageSeq map[string]*Message
	message        chan *Message // 用于ack确认和消息处理之间的异步通信（协程 + chan）

	done chan struct{} // 超时关闭通道
}

func NewConn(s *Server, w http.ResponseWriter, r *http.Request) *Conn {
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Errorf("upgrade err %v", err)
		return nil
	}

	// 把升级的websocket连接封装到自定义连接中
	conn := &Conn{
		Conn:              c,
		s:                 s,
		idle:              time.Now(),
		maxConnectionIdle: s.opt.maxConnectionIdle,
		readMessage:       make([]*Message, 0, 2),
		readMessageSeq:    make(map[string]*Message, 2), // 有容量可以保证seq的顺序
		message:           make(chan *Message, 1),
		done:              make(chan struct{}),
	}

	// keeplive
	go conn.keepalive()

	return conn
}

// appendMsgMq
//
//	@Description: 这个mq是连接conn和readAck的存在，需要做的就是两个任务，分别在第一次ack和第二次ack中
//	@receiver c
//	@param msg
func (c *Conn) appendMsgMq(msg *Message) {
	c.messageMu.Lock()
	defer c.messageMu.Unlock()

	// 读取队列
	// 情况一：队列为空，说明是第一次ack，直接记录
	// 情况二：不为空，说明是第二次ack，进行确认
	if m, ok := c.readMessageSeq[msg.Id]; ok {
		// 已经有消息记录
		if len(c.readMessage) == 0 {
			// 队列中没有消息
			return
		}

		// msg.AckSeq > m.AckSeq（ack是递增的）
		if m.AckSeq >= msg.AckSeq {
			// 没有ack的确认或者重复发送
			return
		}
		// 更新最新的ack
		c.readMessageSeq[msg.Id] = msg
		return
	}
	// 第一次ack，避免客户端发送多余的ack消息
	if msg.FrameType == FrameAck {
		return
	}

	// 记录消息
	c.readMessage = append(c.readMessage, msg)
	c.readMessageSeq[msg.Id] = msg
}

// ReadMessage
//
//	@Description: 包装websocket的readMessage,重新设置活跃时间点
//	@receiver c
//	@return messageType
//	@return p
//	@return err
func (c *Conn) ReadMessage() (messageType int, p []byte, err error) {
	// 包装websocket的readMessage
	messageType, p, err = c.Conn.ReadMessage()

	c.idleMu.Lock()
	defer c.idleMu.Unlock()
	// 有消息接受了，说明是保持连接的，重置空闲时间点
	c.idle = time.Time{}
	return
}

// WriteMessage
//
//	@Description: 包装websocket的WriteMessage,重新设置活跃时间点
//	@receiver c
//	@param messageType
//	@param data
//	@return error
func (c *Conn) WriteMessage(messageType int, data []byte) error {
	c.idleMu.Lock()
	defer c.idleMu.Unlock()

	// 写是线程不安全的，要加锁
	err := c.Conn.WriteMessage(messageType, data)
	c.idle = time.Now()
	return err
}

func (c *Conn) Close() error {
	// 先把通道关闭，防止多次调用close，关闭通道之后再写通道就panic了
	select {
	case <-c.done:
	default:
		close(c.done)
	}

	// 调用第三方库的关闭
	return c.Conn.Close()
}

// keepalive
//
//	@Description: 保持心跳连接，主要判断是空闲时间，使用的是定时器
//	@receiver c
func (c *Conn) keepalive() {
	idleTimer := time.NewTimer(c.maxConnectionIdle)

	defer func() {
		idleTimer.Stop()
	}()

	for {
		select {
		case <-idleTimer.C: // 定时器到期触发通道
			c.idleMu.Lock()
			idle := c.idle
			if idle.IsZero() { // 没有空闲时间
				c.idleMu.Unlock()
				idleTimer.Reset(c.maxConnectionIdle)
				continue
			}
			val := c.maxConnectionIdle - time.Since(idle) // 判断超时差值
			c.idleMu.Unlock()
			if val <= 0 { // 超时了，直接关闭
				c.s.Close(c)
				return
			}
			idleTimer.Reset(val) // 没超时，重设定时器的值
		case <-c.done: // 关闭
			return
		}
	}
}
