package basic

import (
    "bytes"
    "fmt"
    "sync"
    "testing"
)

func TestWgMutext(t *testing.T) {
    buf := bytes.NewBufferString("a")
    fmt.Printf("len:%v cap:%v\n", buf.Len(), buf.Cap())

    // 位运算
    a:=12&11 // i & (i-1) 该运算将 xx 的二进制表示的最后一个 11 变成 00
    fmt.Printf("a: %v\n",a)
}

func wgMutext() {
    var wg sync.WaitGroup
    var mutex sync.Mutex
    total, sum := 0, 0
    for i := 1; i <= 10; i++ {
        wg.Add(1)
        sum += i
        go func(i int) {
            defer wg.Done()
            mutex.Lock()
            total += i
            mutex.Unlock()
        }(i)
    }
    wg.Wait()

    fmt.Printf("total:%d sum %d\n", total, sum)

}
