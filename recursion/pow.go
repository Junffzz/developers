package main

import "fmt"

func main() {
    x := float64(2)
    n := 9
    ret := myPow(x, n)
    fmt.Printf("pow(%v,%v)=%v\n", x, n, ret)

    n>>=1
    fmt.Printf("n>>=1 n:%v\n",n)

    n=n&1
    fmt.Printf("n&1 n:%v\n",n)
}

/**
* #50.x的n次方
x=2,n=10,2E10=1024
* pow(x,n)
1.调库函数O(1)
2.暴力循环：O(N)
3.x*x*x*x*x*x
分治，递归。O(logN)
n偶数：y=x*E[n/2];result=y*y
n奇数：y=x*E[(n-1)/2];result=y*y*x
*/
// 递归分治版本
func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }

    if n < 0 {
        return 1 / myPow(x, -n)
    }

    if n%2 > 0 {
        return x * myPow(x, n-1)
    }
    return myPow(x*x, n/2)
}

// 快速幂法
func myPow2(x float64, n int) float64 {
    if n < 0 {
        x = 1 / x
        n = -n
    }
    var pow float64 = 1
    for n > 0 {
        if n&1 == 1 { //判断奇偶性,为1则是奇数
            pow *= x
        }
        x *= x
        n >>= 1 // 等价于向下取整n//2
    }
    return pow
}
