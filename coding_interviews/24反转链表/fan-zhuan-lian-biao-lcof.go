package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	current := head

	for current != nil {
		temp := current.Next // 下一个节点
		current.Next = prev  // 当前节点指向上一个节点
		prev = current
		current = temp
	}
	return prev
}
