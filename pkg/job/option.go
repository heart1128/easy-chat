package job

import "time"

type (
	RetryOptions func(opts *retryOptions)

	retryOptions struct {
		timeout     time.Duration   // 重试间隔
		retryNums   int             // 次数
		isRetryFunc IsRetryFunc     // 是否重试
		retryJetLag RetryJetLagFunc // 时间策略，指数，线性等
	}
)

func newOptions(opts ...RetryOptions) *retryOptions {
	opt := &retryOptions{
		timeout:     DefaultRetryTimeout,
		retryNums:   DefaultRetryNums,
		isRetryFunc: RetryAlways,
		retryJetLag: RetryJetLagAlways,
	}

	for _, options := range opts {
		options(opt)
	}
	return opt
}

// 单独设置超时时间
func WithRetryTime(timeout time.Duration) RetryOptions {
	return func(opts *retryOptions) {
		if timeout > 0 {
			opts.timeout = timeout
		}
	}
}
