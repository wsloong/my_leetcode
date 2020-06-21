package main

// 二叉树定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：使用递归
func buildTreeWithRecursive(preorder []int, inorder []int) *TreeNode {
	length := len(preorder)

	// 特殊判断，如果前序遍历没有节点直接返回nil
	if length == 0 {
		return nil
	}

	// 使用一个 map 存储 中序遍历 的每个元素及其对应的下标
	// 目的是为了快速得到一个元素在 中序遍历中的位置
	// 递归方法中，对于前序遍历和中序遍历，下标范围都是 0 ~ n-1
	indexMap := map[int]int{}

	for i := 0; i < length; i++ {
		indexMap[inorder[i]] = i
	}

	root := buildTree(preorder, 0, length-1, inorder, 0, length-1, indexMap)
	return root
}

// 递归方法的定义
func buildTree(preorder []int, preorderStart, preorderEnd int, inorder []int, inorderStart, inorderEnd int, indexMap map[int]int) *TreeNode {
	if preorderStart > preorderEnd {
		return nil
	}

	// 树/子树 根节点的值，就是前序遍历开始的第一个值
	rootVal := preorder[preorderStart]
	root := &TreeNode{Val: rootVal}

	// 如果前序遍历的开始等于结束，说明只有一个节点，直接返回即可
	if preorderStart == preorderEnd {
		return root
	} else { // 有多个节点
		// 在中序遍历map中得到根节点的位置
		rootIndex := indexMap[rootVal]

		// 通过得到根节点的位置，可以得到左子树和右子树各自的下标范围和节点数量
		leftNodes, rightNodes := rootIndex-inorderStart, inorderEnd-rootIndex

		// 递归调用
		root.Left = buildTree(preorder, preorderStart+1, preorderStart+leftNodes, inorder, inorderStart, rootIndex-1, indexMap)
		root.Right = buildTree(preorder, preorderEnd-rightNodes+1, preorderEnd, inorder, rootIndex+1, inorderEnd, indexMap)
		return root
	}
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	buildTreeWithRecursive(preorder, inorder)
}
