/*
@Time : 2020/11/13 18:07
@Author : ZhaoJunfeng
@File : arra
*/
package main

import "fmt"

// 数组的赋值和传参
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Printf("slice1:%v,slice2:%v;s3_2:%v\n",slice1,slice2,slice2[0:2])
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Printf("slice1:%v,slice2:%v\n",slice1,slice2)

	var a = make([]int,5,5)
	a[1]=2
	fmt.Printf("slice a1:%v\n",a)

	copy(a,slice2)
	fmt.Printf("slice a2:%v\n",a)
}

func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}