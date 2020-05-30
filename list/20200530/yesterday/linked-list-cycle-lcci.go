// 面试题 02.08. 环路检测

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 判断是否有环，如有环则返回快慢指针在环中相遇的节点
	getIntersect := func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}

		slow := head
		fast := head
		for fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return slow
			}
		}
		return nil
	}

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
