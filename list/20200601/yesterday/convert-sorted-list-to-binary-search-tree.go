// 109. 有序链表转换二叉搜索树

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

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	// 查找中心节点作为数的根
	middleNode := getMiddleNode(head)

	tree := &TreeNode{Val: middleNode.Val}

	// 处理只有一个节点时候
	if head == middleNode {
		return tree
	}
	tree.Left = sortedListToBST(head)
	tree.Right = sortedListToBST(middleNode)
	return tree
}

func getMiddleNode(head *ListNode) *ListNode {
	// 用来从中间断开左边的节点  错误1：没写
	var prevNode *ListNode

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		prevNode = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if prevNode != nil {
		prevNode.Next = nil
	}

	return slow
}
