package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
