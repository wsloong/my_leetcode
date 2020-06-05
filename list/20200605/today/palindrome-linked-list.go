/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
将链表从中间截断，获得左，右两个链表
反转右链表
依次比对两个链表的值
*/
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// 获得右侧链表
	right := slow

	// 反转右链表
	reversedRight := reverseList(right)

	// 返回的结果
	result := true

	leftHead := head
	rightHead := reversedRight
	// 依次比对
	for rightHead != nil && result {
		if rightHead.Val != leftHead.Val {
			result = false
		}

		rightHead = rightHead.Next
		leftHead = leftHead.Next
	}

	// 最后恢复链表
	slow.Next = reverseList(reversedRight)
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
