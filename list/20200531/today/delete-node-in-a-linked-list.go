// 237. 删除链表中的节点

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// ===================早上的 有序链表转换二叉搜索树 ====================

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	middleNode := findMiddleNode(head)

	tree := &TreeNode{Val: middleNode.Val}

	// 如果只有一个节点
	if head == middleNode {
		return tree
	}

	tree.Left = sortedListToBST(head.Next)
	tree.Right = sortedListToBST(middleNode.Next)
	return tree

}

func findMiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if prev != nil {
		prev.Next = nil
	}

	return slow
}
