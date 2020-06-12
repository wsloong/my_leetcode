// 19. 删除链表的倒数第N个节点

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{Next: head}

	slow, fast := head, head
	// for i := 0; i < n; i++ {	// 错误1： 这里增加了一个带头节点，应该多移动一位
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}

	slow.Next = slow.Next.Next

	return newHead.Next
}
