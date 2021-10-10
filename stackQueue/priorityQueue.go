/*
@Time : 2020/12/23 15:17
@Author : ZhaoJunfeng
@File : priority
*/
package main

import (
    "container/heap" // http://cngolib.com/container-heap.html
    "fmt"
)

/**
leetcode 703 Kth Largest Element in a Stream
题目：实时判断数据流中第K大元素。
示例：
int k = 3;
int[] arr = [4,5,8,2];


*/

type IntHeap []int

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
    // Push 和 Pop 使用 pointer receiver 作为参数，
    // 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func (h *IntHeap) Peek() interface{} {
    return (*h)[0]
}

type KthLargest struct {
    minHeap *IntHeap
    k       int
}

func Constructor(k int, nums []int) KthLargest {
    h := &IntHeap{}
    heap.Init(h)
    kth := KthLargest{minHeap: h, k: k}

    for _, v := range nums {
        kth.Add(v)
    }
    return kth
}

func (this *KthLargest) Add(val int) int {
    if this.minHeap.Len() < this.k { // 堆未满，插入元素
        this.minHeap.Push(val)
    } else if this.minHeap.Peek().(int) < val { // 判断插入元素大于当前元素里的最小值，则插入
        heap.Remove(this.minHeap, 0) // 等同于heap.Pop(this.minHeap) ，数组在最小堆中会有序
        this.minHeap.Push(val)
    }
    return this.minHeap.Peek().(int)
}

/**
Your KthLargest object will be instantiated and called as such:
obj:=Constructor(k,nums);
param_1:=obj.Add(val);
*/
func main() {
    k := 3
    nums := []int{4, 5, 8, 2}
    obj := Constructor(k, nums)
    fmt.Println(obj.Add(3))
    fmt.Println(obj.Add(5))
    fmt.Println(obj.Add(10))
    fmt.Println(obj.Add(9))
    fmt.Println(obj.Add(4))
}
