package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	// 从上到下打印，就是二叉树的广度优先搜索(BFS)
	// BFS 通常借助队列的先入先出特性来实现
	// 申请一个 队列，将头结点放入其中
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}
