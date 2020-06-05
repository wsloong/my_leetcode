/*
83. 删除排序链表中的重复元素
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
因为是一个排序的链表，链表的数值非乱序的
可以使用双指针法
*/
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

// ===== 早上的 请判断一个链表是否为回文链表 =====

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	middleNode := slow

	//反转右侧链表
	reversedRight := reverseList(middleNode)

	result := true
	leftList := head
	rightList := reversedRight

	for rightList != nil && result {
		if rightList.Val != leftList.Val {
			result = false
		}
		rightList = rightList.Next
		leftList = leftList.Next
	}

	// 恢复之前的链表
	middleNode.Next = reverseList(reversedRight)
	return result
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}
	return prev
}
