/*
面试题22. 链表中倒数第k个节点

思路回顾
	定义slow, fast两个节点，将fast移动到k的位置
	然后依次移动slow,fast,当fast==nil的时候，即fast移动了链表尾部了
	这时候返回slow即可
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
