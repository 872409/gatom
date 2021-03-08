package gcom

import (
	"context"
	"time"
)

// Context 业务上下文
type Context struct {
	// context.WithTimeout派生的子上下文
	TimeoutCtx context.Context
	// 超时函数
	context.CancelFunc
}

// GetContext 获取业务上下文实例
// d 超时时间
func GetContext(d time.Duration) *Context {
	c := &Context{}
	c.TimeoutCtx, c.CancelFunc = context.WithTimeout(context.Background(), d)
	return c
}