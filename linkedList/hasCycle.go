/*
@Time : 2020/11/25 10:13
@Author : ZhaoJunfeng
@File : hasCycle
*/
package linkedList

func hasCycle(self,head *ListNode) bool{
	var fast,slow = head,head
	for slow!=nil && fast!=nil && fast.Next!=nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow==fast {
			return true
		}
	}
	return false
}