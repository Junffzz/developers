/*
@Time : 2021/5/24 10:19
@Author : ZhaoJunfeng
@File : panics
*/
package defers

import (
    "fmt"
    "testing"
    "time"
)

func TestPanicRecover(t *testing.T) {
    defer println("in mian.")
    tpr2()
}

func tpr1()  {
    go func() {
        defer func() {
            if err := recover(); err != nil {
                println("in goroutine.")
                fmt.Println(err)
            }
        }()
        panic("")

    }()

    time.Sleep(1 * time.Second)
}

func tpr2()  {
    defer func() {
        defer func() {
            panic("panic again and again")
        }()
        panic("panic again")
    }()

    panic("panic once")
}