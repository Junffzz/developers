package main

import (
    "fmt"
    "reflect"
    "sort"
)

/**
映射Map & 集合Set
*/

func main() {

    // 两数之和
    ret := TwoSum([]int{2, 7, 11, 15}, 9)
    fmt.Printf("TwoSum result:%v\n", ret)

    // 三数之和
    ret3 := ThreeSum([]int{-4, -1, -1, 0, 1, 2})
    fmt.Printf("ThreeSum:%v\n", ret3)
}

/**
#242：有效的字母异位词
valid
Anagram
示例1：
输入：s = "anagram",t = "nagaram"
输出：true

示例2：
输入：s = "rat",t="car"
输出：false

1.sort 快排
O(NlogN)
2.Map计数
O(N)
*/
func isAnagram(s string, t string) bool {
    m1 := make(map[string]int)
    m2 := make(map[string]int)

    for _, v := range s {
        m1[string(v)] += 1
    }
    for _, v := range t {
        m2[string(v)] += 1
    }

    return reflect.DeepEqual(m1, m2)
}

/**
#1：两数之和
[2,7,11,15]   两数之和为9
1.暴力 [x,y]=>x+y=9
O(N的2平方)

2.set x+y=9  =>y=9-x
for: x:o->length   O(N)
	 set: 9-x	O(1)
*/
// target是两数之和的结果
func TwoSum(arr []int, target int) [2]int {
    tmpMap := map[int]int{}
    for k, v := range arr {
        if index, ok := tmpMap[target-v]; ok {
            return [2]int{index, k}
        }
        tmpMap[v] = k
    }
    return [2]int{-1, -1}
}

/**
#15 题目：三数之和
示例：
给定数组nums = [-1,0,1.2.-1.-4]
满足要求的三元组集合为：目标数是0  a+b+c=0
[
	[-1,0,1],
	[-1,-1,2]
]
解法：
1.暴力枚举。a,b,c->3Loops O(N3次方)
2.c = -(a+b) => set O(1)
a,b=>2 Loops
3.sortFind 先排序(快排)，再找 O(NlogN)
[-4,-1,-1,0,1,2]
*/
func ThreeSum(nums []int) [][]int {
    // 根据题目中的要求使用双指针来解决问题。https://leetcode-cn.com/problems/3sum/solution/15-san-shu-zhi-he-ha-xi-fa-shuang-zhi-zhen-fa-xi-2/
    sort.Ints(nums) // 判重比较方便
    n := len(nums)
    num := make([][]int, 0)
    // a=nums[i]
    for i := 0; i < n; i++ {
        if i >= 1 && nums[i-1] == nums[i] {
            continue
        }
        avg := -1 * nums[i]
        k := n - 1 // k下标指向数组结尾的位置作为right,c=nums[right]
        // j作为定义下标left, b=nums[left]
        for j := i + 1; j < n; j++ {
            if j >= i+2 && nums[j-1] == nums[j] {
                continue
            }
            for k > j && nums[k]+nums[j] > avg {
                // a+b+c>0说明此时三数之和大了，nums是排序后了，所以right下标应该往左移动，这样三数之和小一些。
                k--
            }
            if k == j {
                break
            }
            if nums[j]+nums[k] == avg {
                num = append(num, []int{nums[i], nums[j], nums[k]})
            }

            // 上述条件都不满足，说明a+b+c<0,left就向右移动
        }
    }
    return num
}
