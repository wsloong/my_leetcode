// 分割链表

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 定义2个带头数组，代表小于x和大于x的2个数组的头
	lessHead, largeHead := new(ListNode), new(ListNode)

	// 这3个值分别用于迭代时候进行替换的值
	less, large, current := lessHead, largeHead, head

	for current != nil {
		if current.Val < x {
			less.Next = current
			less = less.Next
		} else {
			large.Next = current
			large = large.Next
		}
		current = current.Next
	}

	// 这2不要特别注意
	less.Next = largeHead.Next
	large.Next = nil

	return lessHead.Next
}

// =============早上的 环路检测 ==========================
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 检查是否有环，如果有返回环中相遇的节点
	intersectNode := getIntersect(head)
	if intersectNode == nil {
		return nil
	}

	ptr1 := head
	ptr2 := intersectNode
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return ptr1
}

func getIntersect(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}

	return nil
}
