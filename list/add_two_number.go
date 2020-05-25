/*
2. 两数相加

给出两个`非空`的链表用来表示两个非负的整数。其中，它们各自的位数是按照`逆序`的方式存储的，并且它们的每个节点只能存储`一位`数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义一个新的节点为头节点，该节点的值为0，Next为返回结果
	head := new(ListNode)

	p, q, currNode := l1, l2, head

	// 记录进位的值
	var carry int
	// 记录当前位上的l1， l2的值
	var x, y int

	for p != nil || q != nil {
		if p != nil {
			x = p.Val
			p = p.Next
		}

		if q != nil {
			y = q.Val
			q = q.Next
		}

		// 当前位的值
		sum := x + y + carry
		// 进位的值
		carry = sum / 10

		currNode.Next = &ListNode{Val: sum % 10}
		currNode = currNode.Next
	}

	if carry > 0 {
		currNode.Next = &ListNode{Val: carry}
	}

	return head.Next
}
