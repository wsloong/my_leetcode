// 复杂链表的复制
// 同 [138.复制带随机指针的链表](https://leetcode-cn.com/problems/copy-list-with-random-pointer/)

package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	oldNode := head

	newNode := &Node{Val: oldNode.Val}
	visited[oldNode] = newNode

	getClonedNode := func(node *Node) *Node {
		if node == nil {
			return nil
		}

		n, ok := visited[node]
		if ok {
			return n
		}

		visited[node] = &Node{Val: node.Val}
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
