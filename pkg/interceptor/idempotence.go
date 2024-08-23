package interceptor

import (
	"context"
	"easy-chat/pkg/xerr"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// 幂等拦截器

type Idempotent interface {
	// context是幂等请求id携带，method携带请求的方法，判断支不支持幂等
	Identify(ctx context.Context, method string) string
	// 是否支持幂等性
	IsIdempotentMethod(fullMethod string) bool
	// 幂等性的验证
	TryAcquire(ctx context.Context, id string) (resp interface{}, isAcquire bool)
	// 执行之后结果的保存
	SaveResp(ctx context.Context, id string, resp interface{}, respErr error) error
}

var (
	TKey = "easy-chat-idempotence-task-id"      // 请求id的标识
	DKey = "easy-chat-idempotence-dispatch-key" // 设置rpc调度中，rpc请求的标识
)

func ContextWithVal(ctx context.Context) context.Context {
	// 设置请求的id
	return context.WithValue(ctx, TKey, utils.NewUuid())
}

// NewIdempotenceClient
//
//	@Description: 客户端拦截器，设置请求id
func NewIdempotenceClient(idempotent Idempotent) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 获取唯一的请求key
		identify := idempotent.Identify(ctx, method)
		// grpc的头部信息
		ctx = metadata.NewOutgoingContext(ctx, map[string][]string{
			DKey: {identify},
		})

		// 调用实际请求
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// NewIdempotenceServer
//
//	@Description: 服务端拦截器
//	@param idempotent
//	@return grpc.UnaryServerInterceptor
func NewIdempotenceServer(idempotent Idempotent) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any,
		err error) {
		// 客户端请求的时候把key设置在了头部，从头部获取就行
		identify := metadata.ValueFromIncomingContext(ctx, DKey)

		// 过滤需要实现幂等的方法
		if len(identify) == 0 || !idempotent.IsIdempotentMethod(info.FullMethod) {
			// 不进行幂等请求
			return handler(ctx, req)
		}

		fmt.Println("----", "请求进入， 幂等性处理", identify)

		// 验证幂等
		r, isAcquire := idempotent.TryAcquire(ctx, identify[0])
		// 通过幂等，执行
		if isAcquire {
			resp, err = handler(ctx, req)
			fmt.Println("---- 幂等性通过，未执行过，执行任务", identify)

			// 保存执行之后的结果
			if err := idempotent.SaveResp(ctx, identify[0], resp, err); err != nil {
				return resp, err
			}

			return resp, err
		}

		fmt.Println("---- 幂等性未通过，任务已经在执行了", identify)
		if r != nil {
			fmt.Println("--- 任务已经执行完了")
			return r, nil
		}

		// 任务还在执行，出错了
		return nil, errors.WithStack(xerr.New(int(codes.DeadlineExceeded), fmt.Sprintf("存在其他任务执行")))
	}
}

var (
	DefaultIdempotent       = new(defaultIdempotent)
	DefaultIdempotentClient = NewIdempotenceClient(defaultIdempotent)
)

// 实现接口
type defaultIdempotent struct {
	// redis获取和设置请求id，设置定时清理
	*redis.Redis
	// 存储
	*collection.Cache
	// 设置支持幂等的方法
	method map[string]bool
}

func NewDefaultIdempotent(c redis.RedisConf) Idempotent {
	cache, err := collection.NewCache(60 * 60)
	if err != nil {
		panic(err)
	}

	return &defaultIdempotent{
		Redis: redis.MustNewRedis(c),
		Cache: cache,
		method: map[string]bool{
			"/social.social/GrouipCreate": true,
		},
	}
}

// context是幂等请求id携带，method携带请求的方法，判断支不支持幂等
func (d *defaultIdempotent) Identify(ctx context.Context, method string) string {
	id := ctx.Value(TKey)
	// 生成请求id
	rpcId := fmt.Sprintf("%v.%s", id, method)
	return rpcId
}

// 是否支持幂等性
func (d *defaultIdempotent) IsIdempotentMethod(fullMethod string) bool {
	return d.method[fullMethod]
}

// 幂等性的验证
func (d *defaultIdempotent) TryAcquire(ctx context.Context, id string) (resp interface{}, isAcquire bool) {
	// 基于redis锁实现， 分布式过期锁
	retry, err := d.SetnxEx(id, "1", 60*60)
	// 已经有请求拿到锁了，这里拿不到，说明不是幂等操作
	if err != nil {
		return nil, false
	}
	// 拿到锁了
	if retry {
		return nil, true
	}

	// 没拿到锁，也没有错误，在缓存中获取数据
	resp, _ = d.Cache.Get(id)
	return resp, false
}

// 执行之后结果的保存
func (d *defaultIdempotent) SaveResp(ctx context.Context, id string, resp interface{}, respErr error) error {
	d.Cache.Set(id, resp)
	return nil
}
