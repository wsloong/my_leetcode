// 1171. 从链表中删去总和值为零的连续节点

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用空的抬头节点作为链表的头，连接上原链表的头
func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{Next: head}

	prevSumMap := make(map[int]*ListNode)

	var sum int
	for d := newHead; d != nil; d = d.Next {
		sum += d.Val
		prevSumMap[sum] = d
	}

	sum = 0
	for d := newHead; d != nil; d = d.Next {
		sum += d.Val
		d.Next = prevSumMap[sum].Next // 这里不需要担心找不到节点问题
	}

	return newHead.Next
}
