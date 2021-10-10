/*
@Time : 2020/11/18 11:00
@Author : ZhaoJunfeng
@File : reverseLinkedList
*/
package linkedList

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*迭代法
* 定义两个指针： pre 和 cur ；pre 在前 cur 在后。
每次让 pre的 next 指向 cur ，实现一次局部反转
局部反转完成之后，pre 和 cur 同时往前移动一个位置
循环上述过程，直至 pre  到达链表尾部
* 1 -> 2 -> 3 -> 4 -> null
* null <- 1 <- 2 <- 3 <- 4
时间复杂度O(n)
空间复杂度O(1)
*/
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode //前指针节点（指针）
	var curr = head //当前指针节点（指针）
	//每次循环，都将当前节点指向它前面的节点，然后当前节点和前节点后移
	for curr != nil {
		var tmpNext = curr.Next //临时节点，暂存当前节点的下一节点，用于后移
		curr.Next = pre //将当前节点指向它前面的节点
		pre = curr //前指针后移
		curr = tmpNext //当前指针后移
	}
	return pre
}

/**
简洁的递归
使用递归函数，一直递归到链表的最后一个结点，该结点就是反转后的头结点，记作 ret .
此后，每次函数在返回的过程中，让当前结点的下一个结点的 next 指针指向当前节点。
同时让当前结点的 next 指针指向 NULL ，从而实现从链表尾部开始的局部反转
当递归函数全部出栈后，链表反转完成。
* 1 -> 2 -> 3 -> 4 -> null
* null <- 1 <- 2 <- 3 <- 4
*/
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList2(head)
	ret.Next.Next = head
	ret.Next = nil
	return ret
}

/**
lian
 */
func reverseList3(self,head *ListNode)  {

}

func createNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func main() {
	var listNodes []ListNode
	var linkInt = [5]int{1, 2, 3, 4, 5}
	for k, v := range linkInt {
		var nextNode *ListNode
		if k+1 < len(linkInt) {
			nextNode = createNode(linkInt[k+1])
		}
		listNodes = append(listNodes, ListNode{v, nextNode})
	}

	fmt.Printf("linkedList1:%+v\n", listNodes)

	listNodes2 := f1(listNodes)
	fmt.Printf("linkedList2:%+v", listNodes2)
}

func f1(listNodes []ListNode) []ListNode {
	var listNodes2 []ListNode
	listNodes2 = make([]ListNode, 5, 5)
	len := len(listNodes)
	for k := range listNodes {
		l := len - k - 1
		listNodes2[k].Val = listNodes[l : l+1][0].Val
		listNodes2[k].Next = nil
		if l > 0 {
			listNodes2[k].Next = &listNodes[k+1]
		}

		fmt.Printf("k:%v,len:%v,l:%v,node:%v\n", k, len, l, listNodes[l:l+1])
	}
	return listNodes2
}

func f2(listNodes []ListNode) []ListNode {

	return nil
}
