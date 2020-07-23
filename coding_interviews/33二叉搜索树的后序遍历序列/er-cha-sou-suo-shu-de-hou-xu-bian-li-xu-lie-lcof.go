package main

func verifyPostorder(postorder []int) bool {
	var recur func(i, j int) bool

	recur = func(i, j int) bool {
		// 说明此子树节点数量小于等于1，直接返回true
		if i >= j {
			return true
		}

		// 寻找第一个大于根节点(后序遍历，最后一个元素为根节点)
		p := i
		for postorder[p] < postorder[j] {
			p++
		}
		m := p // 将第一个大于根节点的元素索引记为 m

		// p 之后的节点都应该大于根节点
		for postorder[p] > postorder[j] {
			p++
		}

		// p不断后移，直到等于j，说明p之后的值都是大于根节点的值，符合二叉树
		// 然后递归判断左子树和右子树是否正确
		return p == j && recur(i, m-1) && recur(m, j-1)
	}
	return recur(0, len(postorder)-1)
}
