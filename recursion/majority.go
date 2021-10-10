package main

func main() {
	
}

/**
#169：求众数
示例 1：
输入：[1,3,3,2,3]
输出：3

示例 2：
输入：[1,1,1,0,2]
输出：1

解题思路：
count(x)>n/2
1.暴力循环。loop:x嵌套loop count(x)  O(n²)
2.map:{x:count_x} loop=>map.count比较数最大。 O(n)
3.sort:[1,2,3,3,3]
O(NlogN)
4.分治算法 Divide & conquer
一个数组二分为Left,right
left==right:=<left
return count(left)>count(right)
 */
func Majority()  {
    
}