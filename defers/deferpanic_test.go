/*
@Time : 2021/6/25 10:27
@Author : ZhaoJunfeng
@File : deferpanic_test
*/
package defers

import (
    "fmt"
    "testing"
)

func TestPanicDefer(t *testing.T) {
    // demo1()
    demo2()
}

func demo1()  {
    defer fmt.Println("1")
    defer func() {
        fmt.Println("2")
    }()
    panic("3")
}

func demo2()  {
    defer fmt.Println("1")
    defer func() {
        defer func() {
            panic("2")
        }()
        panic("3")
    }()
    panic("4")
}