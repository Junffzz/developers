/*
@Time : 2021/4/3 11:53
@Author : ZhaoJunfeng
@File : heapify_test
*/
package basic

import (
    "fmt"
    "testing"
)

func TestHeapify(t *testing.T) {
    tests := []struct {
        name     string
        heapSize int
        intHeap  []int
        wantData []int
    }{
        {"maxHeapify", 9, []int{3, 5, 7, 2, 0, 1, 4, 6, 8}, []int{3, 5, 7, 2, 0, 1, 4, 6, 8}},
        {"maxHeapify", 6, []int{3, 2, 1, 5, 6, 4}, []int{1, 2, 3, 4, 5, 6}},
    }

    for _, test := range tests {
        if test.name == "maxHeapify" {
            minHeapSort(test.intHeap, test.heapSize)
            t.Logf("maxHeapify %v\n", test.intHeap)
        }
    }

    // 第K大的元素
    nums := []int{2, 3, 1, 6, 0, 5}
    ret := findKthLargest1(nums, 2)
    fmt.Printf("findKthLargest1:%v\n", ret)
}

// 堆排
func heapSorts(tree []int, n int) {
    buildHeapify(tree, n)
    for i := n - 1; i >= 0; i-- {
        tree[i], tree[0] = tree[0], tree[i] // 第i个节点和第0个节点进行交换
        maxHeapify(tree, 0, i)              // heapSize在减小，所以是由i来决定的
    }
}

// 建堆
func buildHeapify(tree []int, heapSize int) {
    lastNode := heapSize - 1
    parent := (lastNode - 1) / 2
    for i := parent; i >= 0; i-- {
        maxHeapify(tree, i, heapSize)
    }
}

// 大顶堆化
func maxHeapify(tree []int, i, heapSize int) {
    if i >= heapSize {
        return
    }
    l, r := 2*i+1, 2*i+2
    max := i
    if l < heapSize && tree[l] > tree[max] {
        max = l
    }

    if r < heapSize && tree[r] > tree[max] {
        max = r
    }

    if i != max {
        tree[i], tree[max] = tree[max], tree[i]
        maxHeapify(tree, max, heapSize)
    }
}

func minHeapSort(tree []int, n int) {
    buildMinHeapify(tree, n)
    for i := n - 1; i >= 0; i-- {
        tree[0], tree[i] = tree[i], tree[0]
        minHeapify(tree, 0, i)
    }
}

func buildMinHeapify(tree []int, n int) {
    lastNode := n - 1
    parent := (lastNode - 1) / 2
    for i := parent; i >= 0; i-- {
        minHeapify(tree, i, n)
    }
}

func minHeapify(tree []int, i, n int) {
    if i >= n {
        return
    }
    l, r := 2*i+1, 2*i+2
    min := i
    if l < n && tree[l] < tree[min] {
        min = l
    }
    if r < n && tree[r] < tree[min] {
        min = r
    }
    if i != min {
        tree[i], tree[min] = tree[min], tree[i]
        minHeapify(tree, min, n)
    }
}

func findKthLargest1(nums []int, k int) int {
    if len(nums) == 0 || k < 0 {
        return -1
    }
    start, end := 0, len(nums)-1
    m := len(nums) - k
    pos := partition1(nums, start, end)
    for pos != m {
        if pos > m {
            end = pos - 1
        } else {
            start = pos + 1
        }
        pos = partition1(nums, start, end)
    }
    return nums[m]
}

func partition1(nums []int, l, r int) int {

    i, j := l, r
    index := nums[i]
    for i < j {
        for i < j && nums[j] >= index {
            j--
        }
        if i < j {
            nums[i] = nums[j]
            i++
        }
        for i < j && nums[i] < index {
            i++
        }
        if i < j {
            nums[j] = nums[i]
            j--
        }
    }
    nums[i] = index
    return i
}
