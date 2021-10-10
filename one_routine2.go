/*
@Time : 2020/10/8 11:53
@Author : ZhaoJunfeng
@File : one_routine2
*/
package main

import (
	"fmt"
	"time"
)

func main()  {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello scheduler")
	}
}
