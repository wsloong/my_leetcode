// 1367. 二叉树中的列表
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)

}

// 定义一个 深度优先搜索 算法
func dfs(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	if head.Val != root.Val {
		return false
	}

	return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)

}
