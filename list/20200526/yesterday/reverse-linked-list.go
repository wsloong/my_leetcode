// 反转链表
/*
1 -> 2 -> 3 -> 4 -> 5
1 <- 2 <- 3 <- 4 <- 5

需要一个变量存储上一个节点的值，定义为prev，默认是nil
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		temp := current.Next // 这时候相当于拿到了相邻的三个节点

		current.Next = prev
		prev = current
		current = temp
	}

	return prev
}