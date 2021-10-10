package main

import "fmt"

func main() {
    ret := buySellStocks([]int{7, 1, 5, 3, 6, 4})
    fmt.Printf("maxProfit:%v\n", ret)
}

/**
#122.买卖股票的最佳时机
持有1股
买卖无数次

例子：
[7,1,5,3,6,4]=>7
[1,2,3,4,5]=>4
[7,6,3,2]=>0

解题思路：
1.DFS  时间复杂度：O(2ⁿ)
a.买1股
b.卖1股
2.贪心算法。O(n)
3.DP动态规划。O(n)
*/
func buySellStocks(prices []int) int {
    var max int
    for i := 0; i < len(prices)-1; i++ {
        if prices[i+1] > prices[i] {
            max += prices[i+1] - prices[i]
        }
    }
    return max
}
