package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 特判
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 定义带头空的带头节点
	newHead := new(ListNode)
	current1, current2, current := l1, l2, newHead

	// 执行循环
	for current1 != nil && current2 != nil {
		if current1.Val <= current2.Val {
			current.Next, current1 = current1, current1.Next
		} else {
			current.Next, current2 = current2, current2.Next
		}
		current = current.Next
	}

	if current1 != nil {
		current.Next = current1
	} else {
		current.Next = current2
	}
	return newHead.Next
}
