package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	// result 结果
	// path 存储的是路径，回溯时候要删除最后的元素
	result, path := [][]int{}, []int{}

	// 定义递归函数
	var recur func(root *TreeNode, tar int)

	// 实现递归函数
	recur = func(root *TreeNode, tar int) {
		if root == nil {
			return
		}

		path = append(path, root.Val)
		tar -= root.Val

		// 找到完整的路径了
		if tar == 0 && root.Left == nil && root.Right == nil {
			// 这里需要copy一份快照，不能直接使用path
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
		}

		// 递归左右节点
		recur(root.Left, tar)
		recur(root.Right, tar)

		path = path[:len(path)-1]
	}

	recur(root, sum)
	return result
}
