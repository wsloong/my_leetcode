// 链表的中心
/*
使用快慢指针方式
慢指针步进为1
快指针步进为2
当快指针到达链表尾部时候，慢指针正好在指针的中间位置
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	// 我这次默写忘记 `fast.Next.Next != nil` 的判断
	for fast != nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
