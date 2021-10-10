/*
@Time : 2020/12/1 20:42
@Author : ZhaoJunfeng
@File : channel2
https://www.bookstack.cn/read/qcrao-Go-Questions/channel-channel%20%E5%8F%91%E9%80%81%E5%92%8C%E6%8E%A5%E6%94%B6%E5%85%83%E7%B4%A0%E7%9A%84%E6%9C%AC%E8%B4%A8%E6%98%AF%E4%BB%80%E4%B9%88.md
*/
package main

import (
    "fmt"
    "time"
)

type user struct {
    name string
    age int8
}
var u = user{name: "Ankur", age: 25}
var g = &u
func modifyUser(pu *user) {
    fmt.Println("modifyUser Received Vaule", pu)
    pu.name = "Anand"
}
func printUser(u <-chan *user) {
    time.Sleep(2 * time.Second)
    fmt.Println("printUser goRoutine called", <-u)
    fmt.Println("printUser goRoutine called2", <-u)
}
func main() {
    c := make(chan *user, 5)
    c <- g
    fmt.Println(g)
    // modify g
    g = &user{name: "Ankur Anand", age: 100}
    go printUser(c)
    go modifyUser(g)
    time.Sleep(5 * time.Second)
    fmt.Println(g)
}