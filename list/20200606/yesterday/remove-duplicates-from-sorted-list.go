// 83. 删除排序链表中的重复元素
// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 这里我下意识的想到要使用双指针方法
// 仔细一想，使用当前指针完全可以
// 记录一下，以后多多思考在动手
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}