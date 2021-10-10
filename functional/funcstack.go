/*
@Time : 2021/6/23 10:51
@Author : ZhaoJunfeng
@File : funcstack
*/
package main

import "fmt"

func main() {
    var a, b int
    b = add2(a)
    fmt.Println(a, b)

}

func add(a int) int {
    var b int
    defer func() {
        a++
        b++
    }()
    a++
    b = a
    return b
}

func add2(a int) (b int) {
    defer func() {
        a++
        b++
    }()
    b++
    return b
}




