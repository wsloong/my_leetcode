// 返回链表中 倒数第 k 个节点

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		head = head.Next
		fast = fast.Next
	}
	return head
}
