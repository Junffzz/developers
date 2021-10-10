/*
@Time : 2020/11/25 11:35
@Author : ZhaoJunfeng
@File : list
https://studygolang.com/articles/13795
实现链表的几种常见操作:
1.判断是否为空的单链表
2.单链表的长度
3.从头部添加元素
4.从尾部添加元素
5.在指定位置添加元素
6.删除指定元素
7.删除指定位置的元素
8.查看是否包含某个元素
9.遍历所有元素
*/
package linkedList

import "fmt"

type Object interface{}

type Node struct {
	Data Object //数据域
	Next *Node  //地址域（指向下一个节点的地址）
}

type List struct {
	headNode *Node //头结点
}

func (l *List) IsEmpty() bool {
	if l.headNode == nil { //只需要判断单链表的头部是否为空即可
		return true
	}
	return false
}

func (l *List) Length() int {
	cur := l.headNode
	var num = 0
	for cur.Next != nil {
		num++
		cur = cur.Next
	}
	return num
}

//头插
func (l *List) Add(data Object) *Node {
	var newNode = &Node{Data: data}
	newNode.Next = l.headNode
	l.headNode = newNode
	return newNode
}

//尾插
func (l *List) Append(data Object) {
	var newNode = &Node{Data: data}
	if l.IsEmpty() {
		l.headNode = newNode
	} else {
		cur := l.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newNode
	}
}

func (l *List) Insert(index int, data Object) {

	if index <= 0 {
		l.Add(data)
	} else if index > l.Length() {
		l.Append(data)
	} else {
		var pre = l.headNode
		for j := 0; j < index; j++ {
			pre = pre.Next
		}
		var newNode = &Node{Data: data}
		newNode.Next = pre.Next
		pre.Next = newNode
	}
}

//删除链表指定值的元素
func (l *List) Remove(data Object) {
	var cur = l.headNode
	if cur.Data == data {
		l.headNode = cur.Next
	} else {
		for cur.Next != nil {
			if cur.Next.Data == data {
				cur.Next = cur.Next.Next
			} else {
				cur = cur.Next
			}
		}
	}
}

func (l *List) RemoveIndex(index int) {
	var cur = l.headNode
	if index <= 0 {
		l.headNode = cur.Next
	} else if index > l.Length() {
		//beyond the scope of length
		fmt.Println("超出链表长度")
		return
	} else {
		var count = 0
		for count < (index-1) && cur.Next != nil {
			cur = cur.Next
			count++
		}
		cur.Next = cur.Next.Next
	}
}

func (l *List) Contain(data Object) bool {
	var cur = l.headNode
	if l.Length() > 0 {
		for cur.Next != nil {
			if cur.Data == data {
				return true
			}
			cur = cur.Next
		}
	}
	return false
}

func (l *List) ShowList() {
	if !l.IsEmpty() {
		cur := l.headNode
		for {
			fmt.Printf("\t%v", cur.Data)
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
	}
}

func (l *List) name()  {

}
