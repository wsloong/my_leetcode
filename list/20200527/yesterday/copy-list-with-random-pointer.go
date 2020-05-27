// 复制带随机指针的链表

/*
可以使用递归方式，代码量更少，但是我在leetcode提交的代码超时
这里还是用迭代方式
官方还有一种解法，不好理解
*/
package main

type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}

// 定义一个map，用于存储已经访问过的节点
var visited = make(map[*ListNode]*ListNode)

// 定义一个方法，用于查找节点
// 如果节点已经访问过，直接返回
// 如果没有访问过，用该节点的值创建一个新节点返回
func getClonedNode(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	if n, ok := visited[node]; ok {
		return n
	}

	visited[node] = &ListNode{Val: node.Val}
	return visited[node]
}

func copyRandomList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 这次默写的代码，错误连连，使用的还是原来的节点指针，错错错
	// visited[head] = head
	// current := head
	// if current != nil {
	// 	current.Next = getClonedNode(current.Next)
	// 	current.Random = getClonedNode(current.Random)
	// 	current = current.Next
	// }

	// 正确的代码
	oldNode := head
	newNode := &ListNode{Val: oldNode.Val}
	visited[oldNode] = newNode

	for oldNode != nil {
		newNode.Next = getClonedNode(oldNode.Next)
		newNode.Random = getClonedNode(oldNode.Random)

		oldNode = oldNode.Next
		newNode = newNode.Next
	}

	return visited[head]
}
