package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmpLeft := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(tmpLeft)
	return root
}
