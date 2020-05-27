// 重排链表

/*
1：先找到中间，拿到第二个链表
	使用快慢指针，
	快指针步进为2
	慢指针步进为1
	当快指针到达尾部时候，慢指针在中间
2: 反转第二个链表
3: 链接2个链表
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// 找中间的节点，将链表分成2个部分
	// 如果链表有偶数个，slow走到的左端点，所以不需要分开考虑奇偶
	slow := head
	fast := head

	// 这里的判断保证了可以获取第二个链表的前一个节点
	// 用于截断第一个链表
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 第二个链表
	newHead := slow.Next
	slow.Next = nil

	// 翻转第二个链表
	newHead = reverseList(newHead)

	// 链表的节点依次链接
	for newHead != nil {
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
	head = head.Next

	// 这里开始为原链表的头部(循环到最后将编程尾部)，其next重置为nil
	prev.Next = nil

	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
