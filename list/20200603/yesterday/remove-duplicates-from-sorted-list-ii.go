// 82. 删除排序链表中的重复元素II

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := &ListNode{Next: head}

	prev := newHead
	current := head
	for current != nil && current.Next != nil {
		if prev.Next.Val != current.Next.Val {
			prev = prev.Next
			current = current.Next
			continue
		}

		// 如果相等，当前指针不停的后移
		for current != nil && current.Next != nil && prev.Next.Val == current.Next.Val {
			current = current.Next
		}

		prev.Next = current.Next
		// prev = current			// 错误1，prev不需要赋值，判断是使用 prev.Next 来判断的
		current = current.Next
	}
	return newHead.Next
}
