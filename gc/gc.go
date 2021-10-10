package main

import (
    "runtime"
    "time"
)

func main() {
    go func() {
        for  {

        }
    }()
     time.Sleep(time.Microsecond)

    runtime.GC()
    /**
    GC在进入STW时，需要等待让所有的用户态代码停止，但是for{}所在的goroutine永远都不会被中断，从而停留在STW阶段。
    实际实践中也是如此，当程序的某个goroutine长时间得不到停止，强行拖慢STW，这种情况下造成的影响（卡死）是非常可怕的。
    在GO1.14之后，这类goroutine能够被异步地抢占，从而使得STW的时间如同普通程序那样，不会超过半毫秒，程序也不会因为仅仅等待一个goroutine的停止而停顿在STW阶段。
     */
    println("ok")
}
