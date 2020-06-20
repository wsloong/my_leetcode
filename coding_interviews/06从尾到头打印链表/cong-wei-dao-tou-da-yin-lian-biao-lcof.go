package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var result []int
	if head == nil {
		return result
	}

	var prev *ListNode
	current := head

	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}

	for prev != nil {
		result = append(result, prev.Val)
		prev = prev.Next
	}
	return result
}
