package websocket

import (
	"math"
	"time"
)

// 配置的默认常量
const (
	defaultMaxConnectionIdle = time.Duration(math.MaxInt64)
	defaultAckTimeout        = 30 * time.Second
)
