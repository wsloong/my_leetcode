package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var tmp []int

		for range queue {
			node := queue[0]
			queue = queue[1:]

			// 偶数层：将值放到队列头部
			if len(result)%2 == 0 {
				tmp = append(tmp, node.Val)
			} else {
				tmp = append([]int{node.Val}, tmp...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, tmp)
	}
	return result
}
