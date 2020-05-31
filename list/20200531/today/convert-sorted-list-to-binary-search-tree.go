// 109. 有序链表转换二叉搜索树

// 没思路啊！！！，就想到找中间节点，然后就没思路了。后期复习多思考下
// 下面是官方的第一个思路

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

	// 找中间节点
	middleNode := findMiddleNode(head)

	// mid成为了BST的根
	tree := &TreeNode{Val: middleNode.Val}

	if head == middleNode {
		return tree
	}

	tree.Left = sortedListToBST(head)
	tree.Right = sortedListToBST(middleNode.Next)
	return tree
}

func findMiddleNode(head *ListNode) *ListNode {
	var prevNode *ListNode // 用来从中间断开左边的节点

	slow := head
	fast := head

	for head != nil && head.Next != nil {
		prevNode = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Handling the case when slowPtr was equal to head.
	if prevNode != nil {
		prevNode.Next = nil
	}
	return slow
}
