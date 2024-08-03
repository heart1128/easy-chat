package handler

import (
	"easy-chat/apps/task/mq/internal/handler/msgTransfer"
	"easy-chat/apps/task/mq/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

type Listen struct {
	svc *svc.ServiceContext
}

func NewListen(svc *svc.ServiceContext) *Listen {
	return &Listen{svc: svc}
}

// Services
//
//	@Description: 返回多个服务对象（消费者），消费者是使用go zero的Service接口包装的,接口只是实现了所有对象的start(), stop方法
//	@receiver l
//	@return []service.Service
func (l *Listen) Services() []service.Service {
	return []service.Service{
		kq.MustNewQueue(l.svc.Config.MsgReadTransfer, msgTransfer.NewMsgReadTransfer(l.svc)),
		// TODO: 可以加载多个消费者
		// 加载kafka消费者对象, 就是实现了Consume接口的对象
		kq.MustNewQueue(l.svc.Config.MsgChatTransfer, msgTransfer.NewMsgChatTransfer(l.svc)),
	}
}
