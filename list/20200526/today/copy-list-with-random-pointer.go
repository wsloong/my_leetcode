// 复制带随机指针的链表
package main

type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}

// 循环迭代
func copyRandomList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 定义一个map，key: 原list中的节点，value：新创建的节点
	visited := make(map[*ListNode]*ListNode)
	oldNode := head
	newNode := &ListNode{Val: oldNode.Val}
	visited[oldNode] = newNode

	// 定义一个方法，用户获取clone过的节点，如果没有找到，则返回一个新的节点
	getClonedNode := func(node *ListNode) *ListNode {
		if node == nil {
			return nil
		}

		n, ok := visited[node]
		if ok {
			return n
		}

		visited[node] = &ListNode{Val: n.Val}
		return visited[node]
	}

	for oldNode != nil {
		newNode.Random = getClonedNode(oldNode.Random)
		newNode.Next = getClonedNode(oldNode.Next)

		oldNode = oldNode.Next
		newNode = newNode.Next
	}

	return visited[head]
}

// ----------------------早上的两数之和复习-------------------------------
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := new(ListNode)
	p, q, current := l1, l2, head

	var x, y, carry int

	for p != nil || q != nil {
		x, y = 0, 0

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

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return head.Next
}
