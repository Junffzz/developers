/*
@Time : 2020/10/15 11:26
@Author : ZhaoJunfeng
@File : slices
*/
package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])

	//s1:=arr[2:]
	//fmt.Println("s1=arr[2:]=",s1)
	//s2:=arr[:]
	//fmt.Println("s2=arr[:]=",s2)
	/***********************slice扩展******************************/
	s1 := arr[2:6]
	fmt.Println("s1:", s1)
	s2 := s1[3:5]
	fmt.Println("s4:", s2)
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	//s4 and s5 view different arr
	fmt.Println("s3:", s3, ";s4:", s4, ";s5:", s5, ";arr:", arr)
	/***********************更新******************************/
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)
}

func updateSlice(s []int) {
	s[0] = 100
}
