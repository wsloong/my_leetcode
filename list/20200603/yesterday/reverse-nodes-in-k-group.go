// 25. K个一组翻转链表

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := &ListNode{Next: head}

	prev := newHead

	// 分组的头，在接下来的循环中会依次后移k个节点
	groupHead := head

	for head != nil {
		// 获取分组的尾
		groupTail := prev
		for i := 0; i < k; i++ {
			groupTail = groupTail.Next
			// 如果到达了尾部，说明不满足k个元素，直接返回
			if groupTail == nil {
				return newHead.Next
			}
		}

		tailNext := groupTail.Next

		// 经过k次后移，已经可以拿到分组的 前一个元素:prev, 头部:groupHead, 尾部:groupTail, 尾部节点的下一个节点:tailNext
		// 反转头部和尾部
		groupHead, groupTail = reverseList(groupHead, groupTail)

		prev.Next = groupHead
		groupTail.Next = tailNext

		prev = groupTail
		groupHead = tailNext
	}
	return newHead.Next
}

func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	// 定义上一个元素，通过不停的后移，直到和尾部相同
	// 这个元素开始定义为尾部节点的下一个元素
	// 在反转后他会和头部节点链接
	prev := tail.Next

	current := head

	for prev != tail {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}

	return tail, head
}
