/*
@Time : 2021/9/8 17:40
@Author : ZhaoJunfeng
@File : longestpalindrome
*/
package main

import "fmt"

func main() {
    s := "babad"
    t := longestPalindrome(s)
    fmt.Printf("t:%v\n", t)
}

func longestPalindrome(s string) string {
    n := len(s)
    if n < 2 {
        return s
    }

    maxLen := 1
    begin := 0

    var dp = make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }

    for L := 2; L <= n; L++ {
        for i := 0; i < n; i++ {
            j := L + i - 1
            if j >= n {
                break
            }

            if s[i] != s[j] {
                dp[i][j] = false
            } else {
                if j-i < 3 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }
            }

            // 只要dp[i][L]==true成立，就表示子串s[i..L]是回文，此时记录回文长度和起始位置
            if dp[i][j] && j-i+1 > maxLen {
                maxLen = j - i + 1
                begin = i
            }
        }
    }
    return s[begin : begin+maxLen]
}
