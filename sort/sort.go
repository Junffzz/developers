package main

import "fmt"

func main() {
    var arr = []int{4, 5, 6, 3, 2, 1}
    result := BubbleSort(arr)
    fmt.Printf("%v\n", result)
}

/**
 * 冒泡排序
 * @Description:
 * @Params:
 * @date: 2021/2/8
 */
func BubbleSort(arr []int) []int {
    length := len(arr)
    if length == 0 {
        return arr
    }

    for i := 0; i < length-1; i++ {
        for j := 0; j < length-i-1; j++ {
            if arr[j] > arr[j+1] {
                tmp := arr[j]
                arr[j] = arr[j+1]
                arr[j+1] = tmp
                // arr[j],arr[j+1] = arr[j+1],arr[j]
            }
        }

    }

    return arr
}

/**
 * 插入排序
 * @Description:
 * @Params:
 * @date: 2021/2/8
 */
func InsertSort(arr []int) []int {
    for i := range arr {
        preIndex := i - 1
        current := arr[i]
        for preIndex >= 0 && arr[preIndex] > current {
            arr[preIndex+1] = arr[preIndex]
            preIndex -= 1
        }
        arr[preIndex+1] = current
    }

    return arr
}
