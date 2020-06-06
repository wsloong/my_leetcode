/*
21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := new(ListNode)
	current := newHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l2
	}

	return newHead.Next
}

// ====== 早上的 445. 两数相加 II  =======

type stack struct {
	items []int
}

func (s *stack) push(val int) {
	s.items = append(s.items, val)
}

func (s *stack) pop() int {
	len := len(s.items)
	if len == 0 {
		return 0
	}

	item := s.items[len-1]
	s.items = s.items[:len-1]
	return item
}

func (s *stack) isEmpty() bool {
	return len(s.items) == 0
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定于两个栈
	s1, s2 := new(stack), new(stack)

	// 将l1数据压入栈中
	for i := l1; i != nil; i = i.Next {
		s1.push(i.Val)
	}

	// 将l2数据压入栈中
	for j := l2; j != nil; j = j.Next {
		s2.push(j.Val)
	}

	// 这里定义一个空的节点，后续节点会依次将该值替换成当前值
	var ans *ListNode
	var carry, x, y int

	for !s1.isEmpty() || !s2.isEmpty() || carry != 0 {
		x, y = 0, 0 // 这里重制值

		if !s1.isEmpty() {
			x = s1.pop()
		}

		if !s2.isEmpty() {
			y = s2.pop()
		}

		sum := x + y + carry
		carry = sum / 10

		current := &ListNode{Val: sum % 10}
		current.Next = ans
		ans = current
	}
	return ans
}
