/*
@Time : 2021/3/26 17:13
@Author : ZhaoJunfeng
@File : bubbleSort_test
*/
package basic

import (
    "fmt"
    "testing"
    "time"
)

func TestBubbleSort(t *testing.T) {
    arr := generateRandomNumber(0, 99999, 99999)
    start := time.Now()
    // arr := []int{3, 4, 2, 6, 8, 0, 1, 5, 7, 9}
    bubbleSort(arr)
    // insertionSort(arr)
    latency := time.Since(start)
    fmt.Printf("bubbleSort time cost:%vs\n", latency.Seconds())
    // fmt.Printf("bubbleSort arr:%v\n", arr)
}

func bubbleSort(arr []int) {
    if len(arr) == 0 {
        return
    }
    l := len(arr)
    for i := 0; i < l; i++ {
        for j := 0; j < l-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

// 插入排序
func insertionSort(arr []int) []int {
    for i:= range arr {
        preIndex:=i-1
        current := arr[i]
        for preIndex >= 0 && arr[preIndex] > current {
            arr[preIndex+1] = arr[preIndex]
            preIndex -=1
        }
        arr[preIndex+1] = current
    }
    return arr
}