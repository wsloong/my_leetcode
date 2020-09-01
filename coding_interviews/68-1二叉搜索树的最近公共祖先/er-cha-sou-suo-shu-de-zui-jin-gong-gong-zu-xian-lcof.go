package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		// p, q 都在 root 的右子树中
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
			// p, q 都在 root 的左子树中
		} else if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
}


