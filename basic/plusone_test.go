/*
@Time : 2021/4/8 00:33
@Author : ZhaoJunfeng
@File : plusone_test
*/
package basic

import (
    "fmt"
    "testing"
)

func TestPlusOne(t *testing.T) {
    tests := []struct {
        paramArr []int
        wantArr  []int
    }{
        {[]int{1, 2, 3}, []int{1, 2, 4}},
        {[]int{1, 2, 3, 4, 9}, []int{1, 2, 3, 5, 0}},
        {[]int{9}, []int{1, 0}},
        {[]int{9, 9}, []int{1, 0, 0}},
    }

    for _, test := range tests {
        var param = make([]int, len(test.paramArr))
        copy(param, test.paramArr)
        ret := plusOne(test.paramArr)
        fmt.Printf("paramArr:%v => rest:%v\n", param, ret)
    }
}

func plusOne(digits []int) []int {
    n := len(digits) - 1
    for {
        digits[n]++
        if digits[n] > 9 {
            if n == 0 {
                digits[n] = 1
                digits = append(digits, 0)
                break
            } else {
                digits[n] = 0
                n--
                continue
            }
        } else {
            break
        }
    }
    return digits
}
