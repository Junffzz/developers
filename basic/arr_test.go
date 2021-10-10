/*
@Time : 2021/4/7 19:08
@Author : ZhaoJunfeng
@File : arr_test
*/
package basic

import (
    "fmt"
    "testing"
)

func TestTwoSum(t *testing.T) {
    tests := []struct {
        name     string
        arr      []int
        target   int
        wantData []int
    }{
        {"twoSum", []int{3, 2, 4}, 6, []int{1, 2}},
    }

    for _, test := range tests {
        ret := twoSum(test.arr, test.target)
        fmt.Printf("test.arr:%v\n", ret)
    }

    defer fmt.Printf("defer1\n")
    defer fmt.Printf("defer2\n")
    defer fmt.Printf("defer3\n")
}

func twoSum(nums []int, target int) []int {
    var ret = make([]int, 0, 2)
    tmpMap := make(map[int]int)
    for i := range nums {
        tmpMap[nums[i]] = i
    }

    for j := range nums {
        a := nums[j]
        if a > target {
            continue
        }
        b := target - a
        if v, ok := tmpMap[b]; ok {
            ret = append(ret, j, v)
            return ret
        }
    }
    return ret
}

func TestFindUniqInt(t *testing.T) {
    ret := tArr([]int{1, 2, 1, 2, 3,3,5})
    fmt.Printf("%d\n", ret)
}

// []int{1,1,2}
func tArr(arr []int) int {
    var ret int
    for _, v := range arr {
        ret ^= v
    }
    return ret
}
