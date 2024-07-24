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
		done:              make(chan struct{}),
	}

	// keeplive
	go conn.keepalive()

	return conn
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
