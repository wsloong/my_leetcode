/*
143. 重排链表

给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:
给定链表 1->2->3->4, 重新排列为 1->4->2->3.

示例 2:
给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

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

	for fast.Next != nil && fast.Next.Next != nil{
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

	tail := head
	head = head.Next
	tail.Next = nil

	for head != nil {
		temp := head.Next
		head.Next = tail
		tail = head
		head = temp
	}
	return tail
}
