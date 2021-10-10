/*
@Time : 2020/12/24 23:00
@Author : ZhaoJunfeng
@File : slidingWindow
*/
package main

import "fmt"

/**
#239：返回滑动窗口中的最大值
Array max sliding window
窗口大小k=3

1,3,-1,-3,5,3,6
解题思路：
1.MaxHeap
    a.维护Heap  logk
    b.Max->top o(1)
总时间复杂度：N*logk

2.Queue => deque 双端队列
    a.入队列
    b.维护
*/
func main() {
    var nums []int
    nums = []int{1, 3, -1, -3, 5, 3, 6}
    res := maxSlidingWindow(nums, 3)
    fmt.Printf("%v\n", res)
}

func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 {
        return nil
    }

    var window []int //窗口只存储nums的下标
    var res []int
    for idx, val := range nums {
        if idx >= k && window[0] <= idx-k {
            window = window[1:]
        }

        for {
            if len(window) > 0 && nums[window[len(window)-1]] <= val {
                window = window[:len(window)-1]
            } else {
                break
            }
        }

        window = append(window, idx)
        if idx >= k-1 {
            res = append(res, nums[window[0]])
        }
    }
    return res
}
