/*
@Time : 2021/3/31 14:04
@Author : ZhaoJunfeng
@File : head
*/
package basic

import (
    "container/heap"
    "fmt"
    "testing"
)

type IntHeap []int

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func (h IntHeap) Len() int {
    return len(h)
}

// 根据这个逻辑适配大顶堆/小顶堆
func (h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Peek() int {
    return h[0]
}

func TestHeapSort(t *testing.T) {
    heapSort()
}

func heapSort() {
    // arr := []int{1, 2, 3, 4, 5, 6}

    intHeap := &IntHeap{}
    intHeap.Push(3)
    intHeap.Push(5)
    intHeap.Push(1)
    intHeap.Push(2)
    intHeap.Push(4)
    intHeap.Push(6)
    heap.Init(intHeap)
    // heap.Remove(intHeap,0)
    t := heap.Pop(intHeap)
    fmt.Printf("%v\n", t)
}

type KthLargest struct {
    minHeap *IntHeap
    k       int
}

func KthConstructor(arr []int, k int) KthLargest {
    h := &IntHeap{}
    heap.Init(h)
    kth := KthLargest{
        minHeap: h,
        k:       k,
    }
    for i := range arr {
        kth.Add(arr[i])
    }

    return kth
}

func (this *KthLargest) Add(val int) {
    if this.minHeap.Len() < this.k {
        this.minHeap.Push(val)
    } else if this.minHeap.Peek() < val { // 堆满了
        heap.Pop(this.minHeap)
        this.minHeap.Push(val)
    }
}


