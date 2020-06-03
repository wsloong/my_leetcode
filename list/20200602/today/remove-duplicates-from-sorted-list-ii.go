// 82. 删除排序链表中的重复元素II

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 因为是排序过的节点，可以使用双指针方法解题
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newList := &ListNode{Next: head}

	prev := newList
	current := head

	for current != nil && current.Next != nil {
		// 注意这里的判断方式，很巧妙
		if prev.Next.Val != current.Next.Val {
			prev = prev.Next
			current = current.Next
			continue
		}
		// 如果相等，就不停的向后移动current指针
		for current != nil && current.Next != nil && prev.Next.Val == current.Next.Val {
			current = current.Next
		}

		prev.Next = current.Next
		current = current.Next

	}
	return newList.Next
}

// ===============早上的 K个一组翻转链表 =============
func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := &ListNode{Next: head}

	prev := newHead
	groupHead := newHead

	// 这里没有使用死循环是兼容了head == nil的情况
	for head != nil {
		groupTail := prev
		for i := 0; i < k; i++ {
			groupTail = groupTail.Next
			if groupTail == nil { // 错误1: 这个判断没有写
				return newHead.Next
			}
		}

		// temp := groupHead.Next		// 错误2:这个错误不应该
		temp := groupTail.Next
		groupHead, groupTail = reverseNode(groupHead, groupTail)

		prev.Next = groupHead
		groupTail.Next = temp
		prev = groupTail

		groupHead = groupTail.Next
	}
	return newHead.Next

}

func reverseNode(head, tail *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode
	current := head 
	for prev != tail {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}

	return tail, head
}