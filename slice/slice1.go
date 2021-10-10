/*
@Time : 2020/10/19 17:22
@Author : ZhaoJunfeng
@File : slice1
*/
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	x int
}

func (t T) String() string {
	return "boo"
}

//func main() {
//	//var a []int
//	//a = make([]int, 0,1)
//	//a = append(a, 2)
//	//fmt.Printf("%+v,len:%v\n", a, len(a))
//
//	t:=T{123}
//	fmt.Printf("%v\n",t)
//	fmt.Printf("%#v\n",t)
//
//}

func main() {
	var arr = make([]int,0,10)
	fmt.Println(arr, "before return ")
	//doSomeHappyThings(arr)
	doSomeReflectThings(&arr)
	fmt.Println(arr, len(arr), cap(arr), "after return")
}

// 如果父方法同时改变切片值，切片需要指针传递
func doSomeHappyThings(arr []int) {
	arr = append(arr, 1)
	fmt.Println(arr, "after append")
}

// 如果父方法同时改变切片值，切片需要指针传递
// https://segmentfault.com/q/1010000019651690
func doSomeReflectThings(arrPtr interface{}) {
	// reflect.Append()
	rt:=reflect.ValueOf(arrPtr)
	arrVal:=rt.Elem()
	arrVal = reflect.Append(arrVal,reflect.ValueOf(1))
	//rt:=reflect.ValueOf(arrPtr)
	rt.Elem().Set(arrVal)
	fmt.Printf("after append, arrPtr:%v, arrVal:%v\n",arrPtr,arrVal)
}