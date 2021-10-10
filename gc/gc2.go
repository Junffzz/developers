package main

import (
    "fmt"
    "runtime/debug"
    "time"
)

func allocate() {
    _ = make([]byte, 1<<20)
}

func printGCStats() {
    t := time.NewTicker(time.Second)
    s := debug.GCStats{}
    for {
        select {
        case <-t.C:
            debug.ReadGCStats(&s)
            fmt.Printf("gc %d last@%v, PauseTotal %v\n",s.NumGC,s.LastGC,s.PauseTotal)
        }
    }
}

func main() {
    // go printGCStats()
    // gc文档：https://mp.weixin.qq.com/s/o2oMMh0PF5ZSoYD0XOBY2Q
    for n := 1; n < 100000; n++ {
        allocate()
    }
}
