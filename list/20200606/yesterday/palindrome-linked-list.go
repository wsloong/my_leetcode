// 请判断一个链表是否为回文链表。

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
先获取中间节点，如果有偶数个，使用右节点作为中间节点
通过中间节点获取右侧节点
反转右侧节点
依次对比左侧节点和右侧节点的值是否相同
将数组恢复到初始状态
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	middleNode := slow
	reversedAfterList := reverseList(middleNode)

	before := head
	after := reversedAfterList
	result := true

	for after != nil && result {
		if before.Val != after.Val {
			result = false
		}
		before = before.Next
		after = after.Next
	}
	slow.Next = reverseList(reversedAfterList)
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
