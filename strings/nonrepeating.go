/*
@Time : 2020/10/18 21:08
@Author : ZhaoJunfeng
@File : nonrepeating
*/
package main

import "fmt"

func main() {
	/*案例：寻找最长不含有重复字符的子串，来源leetcode
	* abcabcbb->abc
	* bbbbb->b
	* pwwkew->wke
	 */
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"),
		)
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))

	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	//for i, ch := range []byte(s) {
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch];ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
