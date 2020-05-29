// 环形链表 II
// 找到环入口节点

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	intersertNode := getIntersect(head)
	if intersertNode == nil {
		return nil
	}

	// 链表头到环入口 和 环入口到相遇节点的距离相同
	ptr1 := head
	ptr2 := intersertNode
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1
}

// 检查是否有环，返回环内快慢指针的相遇节点
func getIntersect(head *ListNode) *ListNode {
	if head == nil {
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
