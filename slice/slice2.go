/*
@Time : 2020/11/12 16:55
@Author : ZhaoJunfeng
@File : slice2
*/
package main

import "fmt"

func main() {
	//s1遍历是否能结束
	s1 := []int{1, 2, 3}
	for _, v := range s1 {
		s1 = append(s1, v)
	}

	fmt.Println("s1:", s1)

	//s2遍历是否能结束
	s2:=[]int{1,2,3}
	for i := 0; i < len(s2); i++ {
		fmt.Println("s2:", len(s2))
		s2 = append(s2, s2[i])
	}
	fmt.Println("s2:", s2)
}
