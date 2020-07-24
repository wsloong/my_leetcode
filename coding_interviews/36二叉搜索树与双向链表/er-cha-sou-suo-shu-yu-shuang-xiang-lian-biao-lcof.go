package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	// 定义一个深度优先的算法函数
	// 遍历顺序为中序遍历：左 -> 根 -> 右

	// 需要定义下面几个节点
	// head: 头节点，用于链接尾节点
	// prev: 上一个节点(会根据递归更改) prev.right = current, current.left = prev.
	var head *Node
	var prev *Node

	var dfs func(current *Node)
	dfs = func(current *Node) {
		if current == nil {
			return
		}

		// 中序遍历之左
		dfs(current.Left)

		if prev == nil {
			head = current
		} else {
			prev.Right, current.Left = current, prev // 牵手
		}

		// 中序遍历之根
		prev = current

		// 中序遍历之右
		dfs(current.Right)
	}

	// 从根开始
	dfs(root)

	// 首尾牵手
	head.Left, prev.Right = prev, head
	return head
}
