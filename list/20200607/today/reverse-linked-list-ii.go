/*
92. 反转链表 II
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
	1 ≤ m ≤ n ≤ 链表长度。

示例:
	输入: 1->2->3->4->5->NULL, m = 2, n = 4
	输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	current := head

	// 根据m递减，可以找到 m的前一个节点和m的节点
	for m > 1 {
		prev = current
		current = current.Next
		m, n = m-1, n-1
	}

	// 定义一个tail节点和con节点
	tail, con := current, prev

	for n > 0 {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
		n -= 1
	}

	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}

	tail.Next = current
	return head
}
