/*
@Time : 2020/5/18 10:38
@Author : ZhaoJunfeng
@File : main
*/
package main

import (
    "fmt"
    "net/url"
)

type A1 struct {
    t1 string
}

func (receiver *A1) Start() {
    fmt.Printf("A1 start.%s\n", receiver.t1)
}

type A2 struct {
    A1
    t2 string
}

func (receiver *A2) Start() {
    fmt.Printf("A2 start,t1:%s, t2:%s\n", receiver.t1, receiver.t2)
}

func main() {
    a := &A2{
        A1: A1{
            t1: "This is T1.",
        },
        t2: "This is T2.",
    }

    fmt.Printf("-----debug a: %+v\n", a)
    a.Start()

    param := url.Values{"ids[]": []string{"1", "2", "3"}}
    param.Set("ids[]", "4,5")
    sparam := param.Encode()
    fmt.Printf("sparam:%v\n", sparam)
}
