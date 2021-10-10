package main

import "fmt"

func main() {
    //res := mySqrt(8)
    res := _mySqrt(8)
    fmt.Printf("%v\n", res)
}

/**
#69.实现int sqrt(int x)函数。
计算并返回x的平方根，其中x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
    由于返回类型是整数，小数部分将被舍去。

解法：
1.二分法
2.牛顿迭代法
*/

// 二分法
func mySqrt(x int) int {
    if x ==0 || x ==1 {
        return x
    }
    left,right:=0,x
    for left<=right{
        mid:=left+(right-left)>>1
        if mid == x/mid {
            return mid
        } else if mid < x/mid {
            left = mid+1
        } else{
            right = mid-1
        }
    }

    if left>x/left {
        return left-1
    }
    return left
}

// 二分法1
func mySqrt1(x int) int {
    l, r := 0, x
    for l < r {
        mid := (l + r + 1) >> 1
        if mid <= x/mid {
            l = mid
        } else {
            r = mid - 1
        }
    }
    return l // r
}

// 二分法2
func mySqrt2(x int) int {
    l, r := 0, x
    ans := -1
    for l <= r {
        mid := l + (r-l)>>1
        if mid*mid <= x {
            ans = mid
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return ans
}

//作者：xilepeng
//链接：https://leetcode-cn.com/problems/sqrtx/solution/golang-er-fen-by-xilepeng/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

//func mySqrt1(x int) float64{
//    // 抄一下 牛顿迭代法
//    var r float64
//    if(x<=1) {
//        return x
//    }
//    r = x
//    for(r > x/r){
//    r = (int)((r + x/r) / 2)
//    }
//    return r
//}
