package main

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// map，存储访问过的节点
	visited, oldNode, newNode := make(map[*Node]*Node), head, &Node{Val: head.Val}
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