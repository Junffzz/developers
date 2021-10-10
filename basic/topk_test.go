/*
@Time : 2021/3/29 20:07
@Author : ZhaoJunfeng
@File : topk_test
*/
package basic

import (
    "math/rand"
    "testing"
)

func TestGetLeastNumbers(t *testing.T) {
    tests := []struct {
        arr []int
        k   int
    }{
        {[]int{3, 2, 1}, 2},
        {[]int{0, 1, 2, 1}, 2},
    }
    for _, test := range tests {
        tmpArr := make([]int, len(test.arr))
        copy(tmpArr, test.arr)

        res := getLeastNumbers2(test.arr, test.k)
        t.Logf("getLeastNumbers Arr:%v K:%v =%v\n", tmpArr, test.k, res)
    }

}

func getLeastNumbers(arr []int, k int) []int {
    if len(arr) <= 0 || k == 0 {
        return nil
    }
    start, end := 0, len(arr)-1
    pos := partition(arr, start, end) // 先找一个基准数小的下标位置
    for pos != k-1 {                  // 一旦index==k-1，即可停止。
        if pos > k-1 {
            end = pos - 1
        } else {
            start = pos + 1
        }

        pos = partition(arr, start, end)
    }
    return arr[:k]
}

func partition(arr []int, l, r int) int {
    i := l
    j := r
    index := arr[i]
    for i < j {
        for i < j && arr[j] >= index {
            j--
        }
        if i < j {
            arr[i] = arr[j]
            i++
        }
        for i < j && arr[i] < index {
            i++
        }
        if i < j {
            arr[j] = arr[i]
            j++
        }
    }

    arr[i] = index
    return i
}

func RandInt64(min, max int) int {
    if min >= max || min == 0 || max == 0 {
        return max
    }
    return rand.Intn(max-min) + min
}

// 练习
func getLeastNumbers2(arr []int,k int) []int{
    if len(arr)<=0 || k == 0 {
        return nil
    }
    start,end := 0, len(arr)-1
    pos:= partition2(arr,start,end)
    for pos!=k-1{
        if pos>k-1{
            end = pos-1
        } else{
            start = pos + 1
        }
        pos = partition2(arr,start,end)
    }
    return arr[:k]
}

func partition2(arr []int,l, r int) int {
    if l>=r {
        return l
    }
    i:=l
    j:=r
    index:=arr[i]
    for i<j {
        for i<j && arr[j]>=index{
            j--
        }
        if i<j {
            arr[i] = arr[j]
            i++
        }
        for i<j && arr[i]<index{
            i++
        }
        if i<j {
            arr[j] = arr[i]
            j--
        }
    }
    arr[i] = index
    return i
}