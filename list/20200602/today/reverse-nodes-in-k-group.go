// 25. K个一组翻转链表

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	newHead := &ListNode{Next: head}

	// 分组的头
	groupHead := head

	prev := newHead

	for head != nil {
		// 分组的尾部
		groupTail := prev
		// 对链表进行分组
		for i := 0; i < k; i++ {
			groupTail = groupTail.Next
			if groupTail == nil {
				return newHead.Next
			}
		}

		// 保存尾部的下一个节点
		temp := groupTail.Next
		groupHead, groupTail = reverseList(groupHead, groupTail)

		// 头尾链接 很重要，很容易出错
		prev.Next = groupHead
		groupTail.Next = temp
		prev = groupTail
		groupHead = groupTail.Next
	}
	return newHead.Next
}

// 反转头尾两个节点中间的节点
// 头结点依次后移到尾节点，则退出
// 反转后之前的尾部和头部互换
func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode
	current := head
	for prev != tail {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}
	return tail, head
}
