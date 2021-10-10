/*
@Time : 2020/11/25 16:44
@Author : ZhaoJunfeng
@File : list_test
*/
package linkedList

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	//初始化链表List
	list := List{}
	//1.向链表尾部追加元素1,2,3,4,5,a,b,c
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append("a")
	list.Append("b")
	list.Append("c")

	fmt.Printf("list len:%v\n", list.Length())
	fmt.Print("1.链表List当前数值为：\n")
	list.ShowList()
	fmt.Println()

	//2.从头部添加元素beforeHead
	list.Add("beforeHead")
	list.Add("beforeHead1")
	fmt.Print("2.链表List当前数值为：\n")
	list.ShowList()
	fmt.Println()

	//3.在指定位置index=5,value=b处插入元素five_insert_value
	list.Insert(5, "five_insert_value")
	fmt.Print("3.指定位置插入five_insert_value，链表List当前的数值为：\n")
	list.ShowList()
	fmt.Println()

	//4.删除元素five_insert_value
	list.Remove("five_insert_value")
	fmt.Print("4.删除five_insert_value后，链表List当前的数值为：\n")
	list.ShowList()
	fmt.Println()

	//5.从位置1删除元素index=1
	list.RemoveIndex(1)
	list.RemoveIndex(0)
	fmt.Print("5.删除位置1的元素后，链表List当前的数值为：\n")
	list.ShowList()
	fmt.Println()

}
