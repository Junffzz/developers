package main

func main() {
	
}

/**
#191.位1的个数
编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。

提示：

请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的示例 3中，输入表示有符号整数 -3。

进阶：

如果多次调用这个函数，你将如何优化你的算法？

示例 1：

输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011中，共有三位为 '1'。
示例 2：

输入：00000000000000000000000010000000
输出：1
解释：输入的二进制串 00000000000000000000000010000000中，共有一位为 '1'。
示例 3：

输入：11111111111111111111111111111101
输出：31
解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。

提示：

输入必须是长度为 32 的 二进制串 。
 */

/**
https://leetcode-cn.com/problems/number-of-1-bits/solution/golang-yi-wei-wei-cao-zuo-fa-by-bloodborne/
移位法
解题思路
1、首先，判断这个数的最后一位是否为1，如果为1，则计数器加1
2、然后，通过右移丢弃掉最后一位，循环执行该操作直到这个数等于0为止。
 */
func hammingWeight(num uint32) int {
    count := 0
    for num != 0 {
        if num&1 == 1 {
            count++
        }
        num = num >> 1
    }
    return count
}

/**
位操作法
解题思路
1、给定一个数n，每进行一次n&(n-1)计算，其结果中都会少了一位1，而且是最后一位。
2、可以通过不断地用n&(n-1)操作去掉n中最后一位1的方法求出n中1的个数
 */
func hammingWeight2(num uint32) int {
    count := 0
    for num > 0 {
        num &= num -1
        count++
    }
    return count
}