// 两数之和

/*
先定义一个空节点指针 head 作为链表的头(带头链表)，返回结果时候，返回 head.Next
这样可以统一处理

定义一个变量用于存储进位数值
在最后如果有进位需要单独处理
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义一个新的节点为头节点，该节点的值为0，Next为返回结果
	head := new(ListNode)
	p, q, currentNode := l1, l2, head

	var carry, x, y int //carry进位值，x:p的值，y:q的值

	for p != nil || q != nil {
		x , y = 0, 0
		
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
		currentNode.Next = &ListNode{Val: sum % 10}
		currentNode = currentNode.Next
	}

	// 如果有进位，需要单独处理
	if carry > 0 {
		currentNode.Next = &ListNode{Val: carry}
	}
	return head.Next
}
