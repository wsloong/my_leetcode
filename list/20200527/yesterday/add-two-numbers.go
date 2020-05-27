// 两数之和

/*
技巧点
1，带头链表
2，处理进位
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)

	p, q, curr := l1, l2, head

	var carry, x, y int

	for p != nil || q != nil {
		x, y = 0, 0 // 这里忘记写了，这里很重要，每次都要重置x，y的值
		
		if p != nil {
			x = p.Val
			p = p.Next
		}

		if q != nil {
			y = q.Val
			q = q.Next
		}

		sum := x + y + carry
		carry = sum / 10

		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return head.Next
}
