// 重排列表

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// 使用快慢指针查找中间的节点
	slow := head
	fast := head

	if fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 获取第二个链表
	newHead := slow.Next

	// 截断链表
	slow.Next = nil

	// 反转第二个链表
	newHead = reverseList(newHead)

	// 链接第一个链表和第二个链表
	// for newHead.Next != nil {   // 错误1：判断条件,因为在原链表上进行交换，这里使用了第二个链表进行判断
	for newHead != nil {
		// temp := head.Next	// 错误2：
		temp := newHead.Next

		newHead.Next = head.Next
		head.Next = newHead

		head = newHead.Next
		newHead = temp
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev := head
	curr := head.Next

	prev.Next = nil // 错误1：没写

	for curr != nil {
		temp := curr.Next

		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}
