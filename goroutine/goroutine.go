/*
@Time : 2020/11/27 11:48
@Author : ZhaoJunfeng
@File : goroutine
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//fmt.Printf("Hello from goroutine %d\n",i)'

				//抢占了，没有交出控制权，导致cpu很高
				a[i]++

				//主动交出控制权
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}
