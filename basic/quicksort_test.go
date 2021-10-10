/*
@Time : 2021/3/25 19:17
@Author : ZhaoJunfeng
@File : quicksort_test
*/
package basic

import (
    "fmt"
    "math/rand"
    "testing"
    "time"
)

// 生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
    // 范围检查
    if end < start || (end-start) < count {
        return nil
    }

    // 存放结果的slice
    nums := make([]int, 0)
    // 随机数生成器，加入时间戳保证每次生成的随机数不一样
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for len(nums) < count {
        // 生成随机数
        num := r.Intn((end - start)) + start

        // 查重
        exist := false
        for _, v := range nums {
            if v == num {
                exist = true
                break
            }
        }

        if !exist {
            nums = append(nums, num)
        }
    }

    return nums
}

func TestQuickSort(t *testing.T) {
    // arr := generateRandomNumber(0, 99999, 99999)
    start := time.Now()
     arr := []int{8, 4, 3, 2, 5, 1, 6, 7, 9, 0, 89, 23, 45, 65, 7, 0}
    quickSort(arr)
    latency := time.Since(start)
    fmt.Printf("quicksort %v time cost:%vs\n",arr, latency.Seconds())
}

func quickSort(arr []int) {
    quickSortRecursion(arr, 0, len(arr)-1)
    return
}

func quickSortRecursion(arr []int, low, high int) {
    if low >= high {
        return
    }
    i := low
    j := high
    index := arr[i]
    for i < j {
        for i < j && arr[j] >= index {
            j--
        }
        if i<j {
            arr[i]=arr[j]
            i++
        }

        for i < j && arr[i] < index {
            i++
        }
        if i<j {
            arr[j]=arr[i]
            j--
        }
    }

    arr[i] = index
    quickSortRecursion(arr, low, i-1)
    quickSortRecursion(arr, i+1, high)
    return
}
