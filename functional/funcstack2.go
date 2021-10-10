/*
@Time : 2021/6/24 12:59
@Author : ZhaoJunfeng
@File : funcstack2
*/
package main

import "fmt"

/****************fdsfdsa ******************/
func foo1(x *int) func() {
    return func() {
        *x = *x+1
        fmt.Printf("foo1 val = %d\n",*x)
    }
}

func foo2(x int) func() {
    return func() {
        x = x+1
        fmt.Printf("foo2 val =%d\n",x)
    }
}

func main()  {
    // Q1第一组实验
    x := 133
    f1 := foo1(&x)
    f2 := foo2(x)
    f1()
    f2()
    f1()
    f2()

    // Q1第二组
    x = 233
    f1()
    f2()
    f1()
    f2()

    // Q1第三组
    foo1(&x)
    foo2(x)
    foo1(&x)()
    foo2(x)()
}
