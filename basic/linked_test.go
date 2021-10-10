/*
@Time : 2021/3/29 15:57
@Author : ZhaoJunfeng
@File : linked_test
*/
package basic

type ListNode struct {
    Val  int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

    var tail *ListNode
    tail = head
    for {
        tail = tail.Next
        if tail.Next == nil {
            break
        }
    }
    return head
}
