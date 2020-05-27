/*
141. 环形链表

使用快慢双指针
慢指针步进为1
快指针步进为2
如果有环，则快慢指针一定会相遇
如果没有，快指针会先到达尾部，其next为nil则返回false
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

// ----------------------------重排链表-------------------
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1 使用快慢指针，找到中间节点，获取第二个链表
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 获取到第二个链表
	newHead := slow.Next
	slow.Next = nil // 截断第一个链表，这时候已经将之前的链表等分成2个了

	// 2 反转链表
	newHead = reverseList(newHead)

	// 重新链接起这两个链表
	for newHead != nil {
		temp := newHead.Next // 错误1： 没有将newHead.Next赋值给临时变量，到时后面赋值丢失了。注意
		newHead.Next = head.Next
		head.Next = newHead

		// head = head.Next		// 错误2：赋值错误，第二个节点没有找到队位置
		head = newHead.Next
		// newHead = newHead.Next // 这次默写的错误地方
		newHead = temp
	}

}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev := head
	curr := head.Next
	prev.Next = nil

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
