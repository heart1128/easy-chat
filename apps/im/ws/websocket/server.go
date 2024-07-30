package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"net/http"
	"sync"
	"time"
)

// 因为go-zero没有websocket，所以需要自己创建server，封装进去

type AckType int

// 三种ack方式，没有，一次，握手
const (
	NoAck AckType = iota
	OnlyAck
	RigorAck
)

func (t AckType) ToString() string {
	switch t {
	case OnlyAck:
		return "OnlyAck"
	case RigorAck:
		return "RigorAck"
	}
	return "NoAck"
}

type Server struct {
	routes map[string]HandlerFunc // 存储服务路由
	addr   string

	*threading.TaskRunner // 并发任务执行，内部使用waitGroup和channel实现的

	authentication Authentication // 连接鉴权(token)

	sync.RWMutex                  // 因为下面的map不是并发安全的，加读写锁保证安全
	connToUser   map[*Conn]string // 每个连接对象对应的用户
	userToConn   map[string]*Conn // 用户对应的连接对象

	patten string        // 路由名
	opt    *ServerOption // 设置选项，配置参数等

	upgrader websocket.Upgrader // websocket
	logx.Logger
}

// NewServer 返回websocket Server实例
func NewServer(addr string, opts ...ServerOptions) *Server {
	opt := newServerOptions(opts...)

	return &Server{
		routes:         make(map[string]HandlerFunc),
		addr:           addr,
		TaskRunner:     threading.NewTaskRunner(opt.concurrency),
		authentication: opt.Authentication,
		connToUser:     make(map[*Conn]string),
		userToConn:     make(map[string]*Conn),
		patten:         opt.patten,
		opt:            &opt,
		upgrader: websocket.Upgrader{
			// 解决跨域问题
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Logger: logx.WithContext(context.Background()),
	}
}

// 下面的函数由返回的NewServer调用

// ServerWs
//
//	@Description: 升级http到websocket
//	@receiver s
//	@param w
//	@param r
func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {

	defer func() {
		// 在协程中，如果捕捉到了panic，就返回，没有就返回空，recover只能在defer中使用，防止panic崩溃
		if r := recover(); r != nil {
			s.Error("server handler ws recover err %v", r)
		}
	}()

	// 根据传入的http升级成websocket
	//conn, err := s.upgrader.Upgrade(w, r, nil)
	//if err != nil {
	//	s.Errorf("upgrade err %v", err)
	//	return
	//}

	// 自定义封装conn
	conn := NewConn(s, w, r)
	if conn == nil {
		return
	}

	if !s.authentication.Auth(w, r) {
		s.Send(&Message{
			FrameType: FrameData,
			Data:      fmt.Sprint("不具备访问权限"),
		}, conn)
		// conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprint("不具备访问权限")))
		conn.Close()
		return
	}

	// 记录连接
	s.addConn(conn, r)

	// 根据连接对象（websocket对象）获取请求信息
	// method
	go s.handlerConn(conn)
}

// addConn
//
//	@Description: 添加连接对象映射的关系，连接之前的鉴权
//	@receiver s
//	@param conn
//	@param req
func (s *Server) addConn(conn *Conn, req *http.Request) {
	uid := s.authentication.UserId(req)

	// 写锁
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	// 防止用户重复登录（暂时不支持重复登录）
	// 关闭之前的连接，给一个新连接
	if c := s.userToConn[uid]; c != nil {
		c.Close()
	}

	s.connToUser[conn] = uid
	s.userToConn[uid] = conn
}

// handlerConn
//
//	@Description: 内部私有函数，处理路由连接任务处理
//	@receiver s
//	@param conn
func (s *Server) handlerConn(conn *Conn) {
	// 获取请求的用户id，方便聊天的时候获取用户使用
	uids := s.GetUsers(conn)
	conn.Uid = uids[0]

	// 处理消息任务
	go s.handlerWrite(conn)

	// 开启了ack
	if s.isAck(nil) {
		go s.readAck(conn)
	}

	for {
		// 1.获取请求消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.Errorf("websocket conn read message err %v", err)
			s.Close(conn)
			return
		}

		// 2.处理消息，json反序列化
		var message Message // 自定义的message，是一个json结构体
		if err = json.Unmarshal(msg, &message); err != nil {
			s.Errorf("json unmarshal err %v, mes %v", err, msg)
			s.Close(conn)
			return
		}

		// 3. TODO 读取消息之后，回复客户端一个ack

		// 4. 根据消息类型进行处理
		// 4.1. 使用了ack，添加消息到ack message qmq中进行ack判读
		if s.isAck(&message) {
			s.Infof("conn message read ack msg %v", message)
			conn.appendMsgMq(&message)
		} else {
			// 4.2 没有使用ack，直接进行发送
			conn.message <- &message
		}
	}
}

func (s *Server) isAck(message *Message) bool {
	// 没有消息，直接判断ack设置的类型
	if message == nil {
		return s.opt.ack != NoAck
	}
	// 同时判断设置的类型和消息中有没有带ack，有一个带了就算有
	return s.opt.ack != NoAck && message.FrameType != FrameNoAck
}

// readAck
//
//	@Description:  进行ack的处理，这时候message实在msgMq中
//	@receiver s
func (s *Server) readAck(conn *Conn) {

	// 1. 判断ack的重传次数
	//send := func(msg *Message, conn *Conn) error {
	//	err := s.Send(msg, conn)
	//	if err == nil {
	//		return nil
	//	}
	//
	//	s.Errorf("message ack OnlyAck send err %v mesasage %v", err, msg)
	//	conn.readMessage[0].errCount++
	//	conn.messageMu.Unlock()
	//
	//	tempDelay := time.Duration(200*conn.readMessage[0].errCount) * time.Microsecond
	//	if max := 1 * time.Second; tempDelay > max {
	//		tempDelay = max
	//	}
	//
	//	time.Sleep(tempDelay)
	//	return err
	//}

	for {
		select {
		case <-conn.done:
			s.Infof("close message ack uid %v", conn.Uid)
			return
		default:
		}

		// 从队列中读取新的消息
		conn.messageMu.Lock()
		if len(conn.readMessage) == 0 {
			conn.messageMu.Unlock()
			// 添加睡眠,让任务切换
			time.Sleep(100 * time.Microsecond)
			continue
		}

		// 读取一条消息
		message := conn.readMessage[0]

		// 判断ack的方式
		switch s.opt.ack {
		case OnlyAck:
			// 一次应答，直接给客户端回复
			s.Send(&Message{
				FrameType: FrameAck,
				Id:        message.Id,
				AckSeq:    message.AckSeq + 1, // seq + 1
			}, conn)

			// 业务处理
			conn.readMessage = conn.readMessage[1:]
			conn.messageMu.Unlock()
			// 给send
			conn.message <- message
		case RigorAck:
			// 两次应答
			if message.AckSeq == 0 {
				// 未确认，进行确认
				conn.readMessage[0].AckSeq++
				conn.readMessage[0].ackTime = time.Now()
				s.Send(&Message{
					FrameType: FrameAck,
					Id:        message.Id,
					AckSeq:    message.AckSeq + 1, // seq + 1
				}, conn)

				s.Infof("message ack RigorAck send mid %v, seq %v, time %v", message.Id, message.AckSeq, message.ackTime)
				conn.messageMu.Unlock()
				continue
			}

			// 已经确认一次了，响应下一次结果

			// 1.客户端返回结果了，再次确认
			msqSeq := conn.readMessageSeq[message.Id] // 上一次客户端的序号
			if msqSeq.AckSeq > message.AckSeq {       // 确认是递增的seq
				// 业务处理
				conn.readMessage = conn.readMessage[1:]
				conn.messageMu.Unlock()
				s.Infof("message ack RigorAck success mid %v", message.Id)
				conn.message <- message
				continue
			}

			// 2.服务端发送一次确认之后，客户端没有再次确认，考虑ack的最大确认时间
			val := s.opt.ackTimeout - time.Since(message.ackTime) // 判断超时时间
			// 2.2 超过了，结束确认，删除所有东西
			if !message.ackTime.IsZero() && val <= 0 {
				delete(conn.readMessageSeq, message.Id)
				conn.readMessage = conn.readMessage[1:]
				conn.messageMu.Unlock()
				continue
			}
			// 2.1 未超过，重传ack
			conn.messageMu.Unlock()
			s.Send(&Message{
				FrameType: FrameAck,
				Id:        message.Id,
				AckSeq:    message.AckSeq, // 重传不加1
			}, conn)
			// 如果一直重传，3秒一次
			time.Sleep(3 * time.Second)
		}
	}
}

// handlerWrite
//
//	@Description:  进行message的处理，一中是没有ack直接写入message通道，一种是确认ack之后可以发送消息
//	@receiver s
func (s *Server) handlerWrite(conn *Conn) {
	for {
		select {
		case <-conn.done:
			// 连接关闭
			return
		case message := <-conn.message: // 处理消息，消息通道发送的
			// 根据消息类型进行处理
			switch message.FrameType {
			case FramePing: // 心跳检测
				s.Send(&Message{FrameType: FramePing}, conn)
			case FrameData:
				// 根据请求的method分发路由，执行
				if handler, ok := s.routes[message.Method]; ok {
					// 找到路由对应的处理方法，执行
					handler(s, conn, message)
				} else {
					// http连接会返回这个错误文本，返回统一的消息格式
					s.Send(&Message{
						FrameType: FrameData,
						Data:      fmt.Sprintf("不存在执行的方法 %v 请检查", message.Method),
					}, conn)
					// conn.WriteMessage(&Message{}, []byte(fmt.Sprintf("不存在执行的方法 %v 请检查", message.Method)))
				}
			}

			// 清理ack消息
			if s.isAck(message) {
				conn.messageMu.Lock()
				delete(conn.readMessageSeq, message.Id)
				conn.messageMu.Unlock()
			}
		}
	}
}

// GetConn
//
//	@Description: 根据用户的uid获取websocket连接对象
//	@receiver s
//	@param uid
//	@return *websocket.Conn
func (s *Server) GetConn(uid string) *Conn {
	s.RWMutex.RLock() // 读锁
	defer s.RWMutex.RUnlock()

	return s.userToConn[uid]
}

// GetConns
//
//	@Description: 根据用户uid组，获取conn组
//	@receiver s
//	@param uid
//	@return []*websocket.Conn
func (s *Server) GetConns(uids ...string) []*Conn {

	if len(uids) == 0 {
		return nil
	}

	s.RWMutex.RLock() // 读锁
	defer s.RWMutex.RUnlock()

	res := make([]*Conn, 0, len(uids))
	for _, uid := range uids {
		res = append(res, s.userToConn[uid])
	}
	return res
}

// GetUsers
//
//	@Description: 根据连接获取这个连接所有的用户uid
//	@receiver s
//	@param conn
//	@return string
func (s *Server) GetUsers(conns ...*Conn) []string {

	s.RWMutex.RLock() // 读锁
	defer s.RWMutex.RUnlock()

	var res []string
	if len(conns) == 0 {
		// 获取全部
		res = make([]string, 0, len(s.connToUser))
		for _, uid := range s.connToUser {
			res = append(res, uid)
		}
	} else {
		res = make([]string, 0, len(conns))
		for _, conn := range conns {
			res = append(res, s.connToUser[conn])
		}
	}

	return res
}

// Close
//
//	@Description: 关闭websocket连接，删除对连接的映射
//	@receiver s
//	@param conn 	自定义封装了websocket conn的连接对象
func (s *Server) Close(conn *Conn) {
	conn.Close()

	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	uid := s.connToUser[conn]
	// 防止重复关闭
	if uid == "" {
		// 已经被关闭
		return
	}

	delete(s.connToUser, conn)
	delete(s.userToConn, uid)
}

// SendByUserId
//
//	@Description: 根据指定的用户uid，获取到conn进行消息发送
//	@receiver s
//	@param msg
//	@param sendIds
//	@return error
func (s *Server) SendByUserId(msg interface{}, sendIds ...string) error {
	if len(sendIds) == 0 {
		return nil
	}
	// 直接找到所有对应的conn对象，调用发送
	return s.Send(msg, s.GetConns(sendIds...)...)
}

// Send
//
//	@Description: 根据传入的所有连接对象发送消息
//	@receiver s
//	@param msg
//	@param conns
//	@return error
func (s *Server) Send(msg interface{}, conns ...*Conn) error {
	if len(conns) == 0 {
		return nil
	}

	// 数据转为jsong发送
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return err
		}
	}

	return nil
}

// AddRoutes
//
//	@Description: 根据方法添加路由到路由组
//	@receiver s
//	@param rs
func (s *Server) AddRoutes(rs []Route) {
	for _, r := range rs {
		s.routes[r.Method] = r.Handler
	}
}

func (s *Server) Start() {
	http.HandleFunc(s.patten, s.ServerWs)
	s.Info(http.ListenAndServe(s.addr, nil))
}

func (s *Server) Stop() {
	fmt.Println("stop websocket Server")
}
