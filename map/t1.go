/*
@Time : 2020/11/27 20:38
@Author : ZhaoJunfeng
@File : t1
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type T struct {
		b int32
		a int16
		c [2]string
	}
	size := unsafe.Sizeof(T{})
	fmt.Println("size:", size)
}
