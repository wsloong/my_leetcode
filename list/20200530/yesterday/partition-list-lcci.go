// 面试题 02.04. 分割链表
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	lessHead, largeHead := new(ListNode), new(ListNode)
	less, large, current := lessHead, largeHead, head

	for current != nil {
		if current.Val < x {
			less.Next = current
			less = less.Next
		} else {
			large.Next = current
			large = large.Next
		}
		current = current.Next
	}

	less.Next = largeHead.Next
	large.Next = nil
	return lessHead.Next
}
