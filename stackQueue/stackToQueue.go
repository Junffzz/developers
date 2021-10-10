/*
@Time : 2020/12/23 14:10
@Author : ZhaoJunfeng
@File : stackToQueue
*/
package main

/**
* 232、225
stack => queue 用堆栈实现队列效果
queue => stack 用队列实现堆栈效果

解题思路：使用两个堆栈
s1：Input输入栈
s2: Output输出栈
*/

type MyQueue struct {
	InStack  []int
	OutStack []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		InStack:  make([]int, 0),
		OutStack: make([]int, 0),
	}
}

// push element x to the back of queue.
func (q *MyQueue) Push(x int) {
	q.InStack = append(q.InStack, x)
}

// remove the element from in front of queue and returns that element.
func (q *MyQueue) Pop() int {
	for _, v := range q.InStack {
		q.OutStack = append(q.OutStack, v)
	}
	q.InStack = make([]int, 0)
	v := q.OutStack[0]
	q.OutStack = q.OutStack[1 : len(q.OutStack)-1]
	return v
}

// Get the front element.
func (q *MyQueue) Peek() int {
	for _, v := range q.InStack {
		q.OutStack = append(q.OutStack,v)
	}
	q.InStack = make([]int,0)
	return q.OutStack[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.InStack)!=0 || len(q.OutStack)!=0
}

/**
Your MyQueue object will be instantiated and called as such:
obj:=NewMyQueue()
obj.Push(x);
param_2:=obj.Pop();
param_3:=obj.Peek();
param_4:=obj.Empty();
 */