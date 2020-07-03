package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{Next: head}

	current := newHead
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
			break // 因为链表中节点的值互不相同
		}
		current = current.Next
	}

	return newHead.Next
}
