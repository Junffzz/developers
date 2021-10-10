package main

import (
	"fmt"
	"sync"
)

var list2 = make([]int, 10)
var wgList2 sync.WaitGroup = sync.WaitGroup{}

func main() {
	// 并发启动协程数
	max := 5
	fmt.Printf("list2 add num=%d\n", max)
	wgList2.Add(10)
	for i := 0; i < max; i++ {
		go notUseLockSetSlice(i)
	}

	for i := 0; i < max; i++ {
		go notUseLockSetSlice(i + max)
	}

	wgList2.Wait()
	fmt.Printf("---list2 len:%v; %v\n", len(list2), list2)

}

// 不加锁设置切片
func notUseLockSetSlice(i int) {
	list2[i] = i
	wgList2.Done()
}

// 加锁设置切片
func useLockSetSlice(i int) {
	list2[i] = i
	wgList2.Done()
}
