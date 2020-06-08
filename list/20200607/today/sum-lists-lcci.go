/*
给定两个用链表表示的整数，每个节点包含一个数位。
这些数位是反向存放的，也就是个位排在链表首部。
编写函数对这两个整数求和，并用链表形式返回结果。

示例：
	输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
	输出：2 -> 1 -> 9，即912

进阶：假设这些数位是正向存放的，请再做一遍。

示例：

输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
输出：9 -> 1 -> 2，即912

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-lists-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 该题和 20200526 的两数之和一样，相比上次，这次代码更简洁了
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	newHead := new(ListNode)
	current := newHead

	var x, y, carry int
	for l1 != nil || l2 != nil || carry != 0 {
		x, y = 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return newHead.Next
}

// 进阶题和 20200606的一样，要注意的是 使用栈 和 中间交换
