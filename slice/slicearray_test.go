/*
@Time : 2020/11/13 14:53
@Author : ZhaoJunfeng
@File : slicearray_test
go test slicearray_test.go -bench . -benchmem -gcflags "-N -l"
*/
package main

import "testing"

func array() [1024]int {
	var a [1024]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	return a
}

func slice() []int {
	a := make([]int, 1024)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	//a:1024
	//a = append(a, 1)
	return a
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}
