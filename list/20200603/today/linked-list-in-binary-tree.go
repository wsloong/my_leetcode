// 1367. 二叉树中的列表
/*
给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。

如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。

一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
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

// 一直连续向下的路径,不存在链表跨越树的左右节点
func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil || root == nil {
		return false
	}

	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(head *ListNode, treeRoot *TreeNode) bool {
	// 链表到达了尾部 => true
	if head == nil {
		return true
	}

	// 二叉树访问到了空节点 => false
	if treeRoot == nil {
		return false
	}

	// 当前匹配的二叉树上节点的值与链表节点的值不相等 => false
	if treeRoot.Val != head.Val {
		return false
	}

	return dfs(head.Next, treeRoot.Left) || dfs(head.Next, treeRoot.Right)
}
