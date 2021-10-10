package basic

import (
    "fmt"
    "testing"
)

/**
* lc:136 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
* @date: 2021/3/8
*/
func SingleNumber(nums []int) int {
    // if nums == nil || len(nums) == 0 {
    //     return -1
    // }
    // var tmpCount = make(map[int]int)
    // for i := range nums {
    //     if count, ok := tmpCount[nums[i]]; ok {
    //         tmpCount[nums[i]] = count + 1
    //         continue
    //     }
    //     tmpCount[nums[i]] = 1
    // }
    //
    // for i := range tmpCount {
    //     if tmpCount[i] == 1 {
    //         return i
    //     }
    // }
    // return -1
    for i := 1; i < len(nums); i++ {
        nums[0] ^= nums[i]
    }
    // 时间复杂度O（n）,空间复杂度O（1）
    return nums[0]
}

func TestSingleNumber(t *testing.T) {
    tests := []struct {
        args []int
    }{
        {
            args: []int{4, 1, 2, 1, 2},
        },
    }

    for _, test := range tests {
        ret := SingleNumber(test.args)
        t.Logf("ret:%v\n", ret)
    }
}

// 计算 a ^ b
func poww(a, b int) {
    // base := a
    // ans := 1
    // for b>0 {
    //     if b & 1 {
    //         ans *=base
    //     }
    //     base *=base
    //     b >>=1
    // }
    // return ans
}

/**
lc.169
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
题解：摩尔投票
*/
func majorityElement(nums []int) int {
    var count = 0 // 计数
    var result = 0
    for _, num := range nums {
        if count == 0 {
            result = num // 假定num为众数
        }
        if num == result {
            count++
        } else {
            count--
        }
    }
    return result
}

func searchMatrix(matrix [][]int, target int) bool {
    if matrix == nil || len(matrix) == 0 {
        return false
    }
    for _, m := range matrix {
        if ok := binarySearch(m, target); ok {
            return true
        }
    }
    return false
}

func binarySearch(arr []int, target int) bool {
    left, right := 0, len(arr)-1
    for left <= right {
        mid := left + (right-left)>>1
        if arr[mid] == target {
            return true
        } else if arr[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return false
}

// lc.88 合并两个有序数组
// 三指针思想，Go语言实现
func twoArrMerge(nums1 []int, m int, nums2 []int, n int) {
    if m == 0 && n > 0 {
        copy(nums1, nums2)
        return
    }

    size := m + n // 代表存放nums1的位置，最大m+n-1
    p := size // 第三个指针，用于指向nums1，从后往前赋值
    m--
    n--
    for i := 0; i < size; i++ {
        if (m < 0 && n < 0) || p == 0 {
            return
        }
        p--
        if m >= 0 && (n < 0 || nums1[m] > nums2[n]) {
            nums1[p] = nums1[m]
            m--
        } else if n >= 0 && (m < 0 || nums1[m] < nums2[n]) {
            nums1[p] = nums2[n]
            n--
        } else {
            nums1[p] = nums1[m]
            m--
            p--
            nums1[p] = nums2[n]
            n--
        }
    }

}

func TestMerge(t *testing.T) {
    tests := []struct {
        NumsOne []int
        M       int
        NumsTwo []int
        N       int
    }{
        {[]int{2, 0}, 1, []int{1}, 1},
        {[]int{0}, 0, []int{1}, 1},
        {[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
    }
    for _, test := range tests {
        twoArrMerge(test.NumsOne, test.M, test.NumsTwo, test.N)
        fmt.Printf("nums1:%v\n", test.NumsOne)
    }

}
