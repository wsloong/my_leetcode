/*
445. 两数相加 II

给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶：
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例：

输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 先反转两个数组然后按照之前的两数之和进行求解
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	reversedL1 := reverseList(l1)
	reversedL2 := reverseList(l2)

	newHead := new(ListNode)
	current := newHead
	var sum, carry int

	for reversedL1 != nil && reversedL2 != nil {
		sum = reversedL1.Val + reversedL2.Val + carry
		current.Next = &ListNode{Val: sum % 10}
		carry = sum / 10

		reversedL1 = reversedL1.Next
		reversedL2 = reversedL2.Next
		current = current.Next
	}

	for reversedL1 != nil {
		sum = reversedL1.Val + carry
		current.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		reversedL1 = reversedL1.Next
		current = current.Next
	}

	for reversedL2 != nil {
		sum = reversedL2.Val + carry
		current.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		reversedL2 = reversedL2.Next
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	// 最后反转新链表
	return reverseList(newHead.Next)
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

// 如果不反转链表。可以借用栈这个数据结构
// go语言没有栈这个结构，这里简单进行构造一个

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

func addTwoNumbersWithStack(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := new(stack), new(stack)

	for i := l1; i != nil; i = i.Next {
		s1.push(i.Val)
	}

	for j := l2; j != nil; j = j.Next {
		s2.push(j.Val)
	}

	var ans *ListNode
	var carry, x, y int

	for !s1.isEmpty() || !s2.isEmpty() || carry != 0 {
		x, y = 0, 0

		if !s1.isEmpty() {
			x = s1.pop()
		}

		if !s2.isEmpty() {
			y = s2.pop()
		}

		sum := x + y + carry
		carry = sum / 10

		// 注意这里的交换
		current := &ListNode{Val: sum % 10}
		current.Next = ans
		ans = current
	}
	return ans
}
