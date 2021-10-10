package syncs

import (
    "fmt"
    "runtime"
    "runtime/debug"
    "sync"
    "sync/atomic"
)

/**
 * 临时对象池
https://www.cnblogs.com/qcrao-2018/p/12736031.html
对于很多需要重复分配、回收内存的地方，sync.Pool 是一个很好的选择。频繁地分配、回收内存会给 GC 带来一定的负担，严重的时候会引起 CPU 的毛刺，而 sync.Pool 可以将暂时不用的对象缓存起来，待下次需要的时候直接使用，不用再次经过内存分配，复用对象的内存，减轻 GC 的压力，提升系统的性能。
临时对象池的实例不应该被复制

1.关键思想是对象的复用，避免重复创建、销毁。将暂时不用的对象缓存起来，待下次需要的时候直接使用，不用再次经过内存分配，复用对象的内存，减轻 GC 的压力。
2.sync.Pool 是协程安全的，使用起来非常方便。设置好 New 函数后，调用 Get 获取，调用 Put 归还对象。
3.Go 语言内置的 fmt 包，encoding/json 包都可以看到 sync.Pool 的身影；gin，Echo 等框架也都使用了 sync.Pool。
4.不要对 Get 得到的对象有任何假设，更好的做法是归还对象时，将对象“清空”。
5.Pool 里对象的生命周期受 GC 影响，不适合于做连接池，因为连接池需要自己管理对象的生命周期。
6.Pool 不可以指定⼤⼩，⼤⼩只受制于 GC 临界值。
7.procPin 将 G 和 P 绑定，防止 G 被抢占。在绑定期间，GC 无法清理缓存的对象。
8.在加入 victim 机制前，sync.Pool 里对象的最⼤缓存时间是一个 GC 周期，当 GC 开始时，没有被引⽤的对象都会被清理掉；加入 victim 机制后，最大缓存时间为两个 GC 周期。
9.Victim Cache 本来是计算机架构里面的一个概念，是 CPU 硬件处理缓存的一种技术，sync.Pool 引入的意图在于降低 GC 压力的同时提高命中率。
10.sync.Pool 的最底层使用切片加链表来实现双端队列，并将缓存的对象存储在切片中。

 * @Description:
 * @Params:
 * @date: 2021/2/13
 */
func main() {
    // 禁用GC，并保证在main函数执行结束前恢复GC
    defer debug.SetGCPercent(debug.SetGCPercent(-1))
    var count int32

    newFunc := func() interface{} {
        return atomic.AddInt32(&count, 1)
    }
    pool := sync.Pool{New: newFunc}

    // New字段值的作用
    v1 := pool.Get()
    fmt.Printf("Value 1:%v\n", v1)

    // 临时对象池的存取
    pool.Put(10)
    pool.Put(11)
    pool.Put(12)
    v2 := pool.Get()
    fmt.Printf("Value 2:%v\n", v2)

    // 垃圾回收对临时对象池的影响
    debug.SetGCPercent(100)
    runtime.GC()
    v3 := pool.Get()
    fmt.Printf("Value 3:%v\n", v3)
    pool.New = nil
    v4 := pool.Get()
    fmt.Printf("Value 4:%v\n", v4)
}


var pool *sync.Pool

type Person struct {
    Name string
}

func initPool()  {
    pool =  &sync.Pool{
        New: func() interface{} {
            fmt.Println("Creating a new Person")
            return new(Person)
        },
    }
}

func case1()  {
    initPool()

    p:=pool.Get().(*Person)
    fmt.Println("首次从pool里获取：",p)

    p.Name = "first"
    fmt.Printf("设置 p.Name = %s\n",p.Name)

    pool.Put(p)

    fmt.Println("Pool 里已有一个对象：&{first},调用 Get: ",pool.Get().(*Person))
    fmt.Println("Pool 没有对象了，调用Get: ",pool.Get().(*Person))
}