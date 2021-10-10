/*
@Time : 2021/3/29 10:39
@Author : ZhaoJunfeng
@File : strings_test
*/
package basic

import (
    "fmt"
    "strings"
    "testing"
)

func TestIsPalindrome(t *testing.T) {
    s := "race a car"
    res := isPalindrome(s)
    fmt.Printf("isPalindrome %v\n", res)
}

// 验证回文串
func isPalindrome(s string) bool {
    var sgood string
    for i := range s {
        if isAlNum(s[i]) {
            sgood += string(s[i])
        }
    }

    sgood = strings.ToLower(sgood)
    n := len(sgood)
    // 双指针解法
    for i := 0; i < n/2; i++ {
        if sgood[i] != sgood[n-i-1] {
            return false
        }
    }
    return true
}

func isAlNum(b byte) bool {
    if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
        return true
    }
    return false
}

// 双指针，空间复杂度O(1)，效率也高
func isPalindrome2(s string) bool {
    s = strings.ToLower(s)
    left, right := 0, len(s)-1
    for left < right {
        for left < right && !isAlNum(s[left]) {
            left++
        }
        for left < right && !isAlNum(s[right]) {
            right--
        }
        if left < right {
            if s[left] != s[right] {
                return false
            }
            left++
            right--
        }
    }
    return true
}

func TestLengthOfLongestSubstring(t *testing.T) {
    tests := []struct {
        param   string
        wantRet int
    }{
        {"abcabcbb", 3},
        {"bbbbb", 1},
        {"pwwkew", 3},
        {"", 0},
    }

    for _, test := range tests {
        ret := lengthOfLongestSubstring(test.param)
        if ret == test.wantRet {
            t.Logf("success!s:%v,len:%v,wantRet:%v\n", test.param, ret, test.wantRet)
        } else {
            t.Errorf("fail!s:%v,len:%v,wantRet:%v\n", test.param, ret, test.wantRet)
        }
    }
}

// 无重复字符的最长子串
// 滑动窗口解法
func lengthOfLongestSubstring(s string) int {
    n := len(s)
    if n <= 1 {
        return n
    }
    charCountMap := make([]int, 256) // 字符词频表
    r := 0
    ans := 0
    for i := 0; i < n; i++ {
        for ; r < n && charCountMap[s[r]] == 0; r++ {
            charCountMap[s[r]]++
        }
        ans = max(ans, r - i)
        if r == n {
            break
        }
        charCountMap[s[i]]--
    }
    return ans
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}
