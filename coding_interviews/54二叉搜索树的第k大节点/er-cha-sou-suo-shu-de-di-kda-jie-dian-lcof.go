package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var result, count int

	// 定义深度优先算法
	var dfs func(root *TreeNode)

	// 实现算法
	// 中序遍历是有序的，
	// 但是因为要求第K大的值，这里的使用中序遍历的倒序，即 右-> 根 -> 左
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Right)
		count++

		// 当count=k的时候，说明找到了
		if count == k {
			result = root.Val
			return
		}

		dfs(root.Left)

	}

	dfs(root)
	return result
}
