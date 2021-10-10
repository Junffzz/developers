/*
@Time : 2021/4/16 10:05
@Author : ZhaoJunfeng
@File : dp_test
*/
package basic

import (
    "math"
    "testing"
)

func TestCoinChange(t *testing.T) {
    tests := []struct {
        coins  []int
        amount int
    }{
        {[]int{1,2,5},11},
        {[]int{2},3},
        {[]int{1},0},
    }
    for _, test := range tests {
        ret := coinChange(test.coins, test.amount)
        t.Logf("coinChange result:%v\n", ret)
    }
}

func coinChange(coins []int, amount int) int {
    if len(coins) == 0 {
        return -1
    }
    var dp = make([]int, amount+1)
    dp[0] = 0
    for i := 1; i <= amount; i++ {
        var min = math.MaxInt64
        for j := 0; j < len(coins); j++ {
            // 所有小于等于i的面值coins[j]，并且最优解小于默认最大值
            if coins[j] <=i && dp[i-coins[j]] < min {
                min = dp[i-coins[j]] + 1 // 更新dp
            }
        }
        dp[i] = min
    }
    if dp[amount] == math.MaxInt64 {
        return -1
    }
    return dp[amount]
}
