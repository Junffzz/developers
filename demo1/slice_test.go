/*
@Time : 2021/3/18 16:33
@Author : ZhaoJunfeng
@File : slice
*/
package demo1

import (
    "fmt"
    "github.com/spf13/cast"
    "testing"
)

func TestSliceConvert(t *testing.T) {
    a := []byte("123")
    b := cast.ToInt(cast.ToString(a))
    t.Logf("b:%v", b)

    var c = 4 << (^uintptr(0) >> 63)
    fmt.Printf("c:%v\n",c)
}


