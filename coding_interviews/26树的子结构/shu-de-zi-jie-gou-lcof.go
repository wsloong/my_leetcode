package main

type ListNode struct {
	Val int
	Left *ListNode
	Right *ListNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var recur func(a, b *TreeNode) bool 
	recur = func(a, b *TreeNode) bool {
		// 如果b为空，说明全部找到了。直接返回true就可以了
		if b == nil {
			return true
		}
		// 如果a为空，或者a的值不等于b的值，返回false
		if a == nil || a.Val != b.Val {
			return false
		}

		return recur(a.Left, b.Left) && recur(a.Right, b.Right)
	}

	return (A != nil && B != nil) && (isSubStructure(A.Left, B) || isSubStructure(A.Right, B) || recur(A, B))
}