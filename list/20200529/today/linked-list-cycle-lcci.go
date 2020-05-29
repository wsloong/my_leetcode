// 面试题 02.08.环路检测
// 返回环的开头节点

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	intersectNode := getIntersect(head)
	// 如果没有相遇的节点，说明没有环
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

// 先找是否有环，如果有环返回快慢节点在环中相遇的节点
func getIntersect(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}
