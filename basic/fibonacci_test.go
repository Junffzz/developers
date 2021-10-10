/*
@Time : 2021/4/12 18:21
@Author : ZhaoJunfeng
@File : fibonaqi_test
*/
package basic

import (
    "fmt"
    "testing"
)

var count int

func TestFibonacci(t *testing.T) {
    tests := []struct {
        name   string
        method uint8 // 1:递归，2：正推
        n      int
    }{
        {name: "递归法", method: 1, n: 10},
        {name: "正推法", method: 2, n: 10},
    }

    for _, test := range tests {
        var ret int
        switch test.method {
        case 1:
            ret = Fibonacci(test.n)
        case 2:
            ret = Fibonacci2(test.n)
        }

        fmt.Printf("%v:%v\n", test.name, ret)
    }

}

var rest = make(map[int]int)

// 递归
func Fibonacci(n int) int {
    if ret, ok := rest[n]; ok {
        return ret
    }
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci2(n int) int {
    var arr = make([]int, n+1)
    arr[0] = 0
    arr[1] = 1
    for i := 2; i <= n; i++ {
        arr[i] = arr[i-1] + arr[i-2]
    }
    return arr[n]
}
