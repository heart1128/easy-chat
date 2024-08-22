package job

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

// 定义重试的时间策略
type RetryJetLagFunc func(ctx context.Context, retryCount int, lastTime time.Duration) time.Duration

func RetryJetLagAlways(ctx context.Context, retryCount int, lastTime time.Duration) time.Duration {
	return DefaultRetryJetLag
}

// 是否进行重试
type IsRetryFunc func(ctx context.Context, retryCount int, err error) bool

func RetryAlways(ctx context.Context, retryCount int, err error) bool {
	return true
}

// 实际重试的操作
func WithRetry(ctx context.Context, handler func(ctx context.Context) error, opts ...RetryOptions) error {
	opt := newOptions(opts...)

	// 判断程序本身是否设置了超时
	_, ok := ctx.Deadline()
	// 如果没有设置，就用上下文设置超时
	if !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opt.timeout)
		defer cancel()
	}

	var (
		herr        error
		retryJetLag time.Duration
		ch          = make(chan error, 1)
	)

	// 不断重试，直到重试超次数或者执行完成
	for i := 0; i < opt.retryNums; i++ {
		// 协程执行
		go func() {
			ch <- handler(ctx)
		}()

		select {
		case herr = <-ch:
			// 执行之后没问题，就跳出
			if herr == nil {
				return nil
			}
			// 有问题，判断是否要重试
			if !opt.isRetryFunc(ctx, i, herr) {
				return herr
			}
			// 重试时间间隔
			retryJetLag = opt.retryJetLag(ctx, i, retryJetLag)
			time.Sleep(retryJetLag)
		case <-ctx.Done(): // 协程超时了
			return errors.New("任务超时")
		}
	}
	return herr
}
