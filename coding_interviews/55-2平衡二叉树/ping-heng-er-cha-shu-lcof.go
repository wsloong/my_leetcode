package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		if left == -1 {
			return -1
		}

		right := dfs(root.Right)
		if right == -1 {
			return -1
		}

		if abs(left, right) <= 1 {
			return max(left, right) + 1
		}
		return -1
	}
	return dfs(root) != -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
