// 面试题 02.03. 删除中间节点

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// ==========早上的 1367. 二叉树中的列表 ======

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

	// 这一步的 isSubPath(head, root.Left) || isSubPath(head, root.Right) 要注意
	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

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
