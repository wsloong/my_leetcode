package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一： 使用深度优先算法
func maxDepth(root *TreeNode) int {
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		return max(dfs(root.Left), dfs(root.Right)) + 1
	}

	return dfs(root)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 方法二： 使用层序遍历方法
func maxDepthWithQueue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue, res := []*TreeNode{root}, 0
	for len(queue) > 0 {
		var tmp []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}

			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
		res++
	}
	return res

}
