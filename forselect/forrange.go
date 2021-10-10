/*
@Time : 2021/5/18 10:36
@Author : ZhaoJunfeng
@File : forrange
*/
package main

import (
    "fmt"
)

func main() {
    arr := []int{1, 2, 3}
    // 在遍历切片时追加的元素不会增加循环的执行次数，所以循环最终还是停了下来。
    for _, v := range arr {
        arr = append(arr, v)
    }
    fmt.Println(arr)
    c:=make(chan bool,1)
    for {
        select {
        case <-c:
            fmt.Println("hello c!")
        default:
            fmt.Println("default!")
        }
    }


}
