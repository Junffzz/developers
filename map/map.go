/*
@Time : 2020/11/27 18:15
@Author : ZhaoJunfeng
@File : map
*/
package main

import "fmt"

func main()  {
	//type Person struct {
	//	Age int
	//}
	//
	//var m = make(map[int]Person)
	//m[1] = Person{Age: 0}

	testPanic()
}

func testPanic()  {
	var result = make(map[int]int)
	defer func() {
		if r:=recover();r!=nil {
			fmt.Println("====1",r)
		}
	}()
	go func() {
		for  {
			_=result[1]

		}
	}()

	for  {
		result[1]=1
	}
}
