/*
@Time : 2020/10/7 12:53
@Author : ZhaoJunfeng
@File : goruntine
*/
package main

import (
    "fmt"
    "os"
    "runtime/trace"
)

func main() {
    // 创建trace文件
    f, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // 启动trace goroutine
    err = trace.Start(f)
    if err != nil {
        panic(err)
    }
    defer trace.Stop()

    // main
    fmt.Println("Hello trace")

    i:=1
    defer fmt.Println(i)
    i = 2
    fmt.Println(i)

}
