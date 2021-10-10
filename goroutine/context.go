package main

import (
	"context"
	"fmt"
	"time"
)

/**
文章：https://mp.weixin.qq.com/s/A03G3_kCvVFN3TxB-92GVw

context支持的方法：
WithCancel：基于父级 context，创建一个可以取消的新 context。
WithDeadline：基于父级 context，创建一个具有截止时间（Deadline）的新 context。
WithTimeout：基于父级 context，创建一个具有超时时间（Timeout）的新 context。
Background：创建一个空的 context，一般常用于作为根的父级 context。
TODO：创建一个空的 context，一般用于未确定时的声明使用。
WithValue：基于某个 context 创建并存储对应的上下文信息。

1.Context接口：
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
Deadline：获取当前 context 的截止时间。
Done：获取一个只读的 channel，类型为结构体。可用于识别当前 channel 是否已经被关闭，其原因可能是到期，也可能是被取消了。
Err：获取当前 context 被关闭的原因。
Value：获取当前 context 对应所存储的上下文信息。

2.Canceler接口：
type canceler interface {
 cancel(removeFromParent bool, err error)
 Done() <-chan struct{}
}
cancel：调用当前 context 的取消方法。
Done：与前面一致，可用于识别当前 channel 是否已经被关闭。
*/
func main() {
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 1*time.Millisecond)
	defer cancel()

	go test1(ctx)

	a, ok := ctx.Deadline()
	fmt.Printf("ctx.Deadline: %v,ok:%v\n", a, ok)
	time.Sleep(1 * time.Minute)
}

func test1(ctx context.Context) {
	select {
	case <-time.After(1 * time.Second):
		fmt.Printf("overslept\n")
	case <-ctx.Done():
		fmt.Printf("ctx.Done %v\n", ctx.Err())
	}
}
