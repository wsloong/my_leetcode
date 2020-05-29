// 328. 奇偶链表

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 定义奇数和偶数的节点
	odd, even := head, head.Next

	// 定义偶数链表的头
	evenHead := even

	// for even != nil {		// 错误1：这里需要判断next不能为空
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		// even.Next = odd.Next.Next		// 这里不够简洁
		// odd = odd.Next
		odd = odd.Next
		even.Next = odd.Next

		even = even.Next
	}

	odd.Next = evenHead
	return head
}
