/*
@Time : 2021/5/21 10:32
@Author : ZhaoJunfeng
@File : defer_test
*/
package defers

import (
    "fmt"
    "testing"
)

/**
defer是在return之前执行的。这个在 官方文档中是明确说明了的。要使用defer时不踩坑，最重要的一点就是要明白，return xxx这一条语句并不是一条原子指令!
函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。
defer表达式可能会在设置函数返回值之后，在返回到调用函数之前，修改返回值，使最终的函数返回值与你想象的不一致。
return xxx会被改写成:
    返回值 = xxx
    调用defer函数
    空的return
 */
func TestDefers(t *testing.T) {
    res := tf2()
    fmt.Printf("%v\n", res)
}

func tf() (result int) {
    defer func() {
        result++
    }()
    return 0
}

func tf2() (r int) {
    t := 5
    defer func() {
        t = t + 5
    }()
    return t
}

func tf3() (r int)  {
    defer func(r int) {
        r = r+6
    }(r)
    return 1
}

func ef4() int {
    a:=1
    defer func(a int) {
        a = a+6
    }(a)

    return tmpReturn(a)
}

func tmpReturn(x int) int {
    x = x+5
    return x
}