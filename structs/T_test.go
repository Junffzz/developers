/*
@Time : 2021/6/11 10:22
@Author : ZhaoJunfeng
@File : T_test
*/
package structs

import (
    "fmt"
    "testing"
)

type T struct {
    Val int
}

func (receiver T) A()  {
    receiver.Val = 1
    fmt.Printf("%v method.\n","A")
}

func (receiver *T) B()  {
    receiver.Val = 2
    fmt.Printf("%v method.\n","B")
}

func TestTMethod(t *testing.T) {
    var p T
    p.A()
    fmt.Printf("pA:%v\n",p)
    p.B()
    fmt.Printf("pB:%v\n",p)
    (&p).B()

    f:=T.A
    f(p)

    pt:=&T{}
    pt.A()
    fmt.Printf("ptA:%v\n",p)
    pt.B()
    fmt.Printf("ptB:%v\n",p)
}