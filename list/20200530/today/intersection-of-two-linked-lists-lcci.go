// 链表相交

/*

两个指针分别从 headA 和 headB 出发，如果指向 headA 的指针到达了末尾，就从 headB 出发，另一个也一样，这样可以保证能够在交点处相遇。
为什么呢？因为假设
headA 的结构是 12345 | 678
headB 的结构是 ab | 678

第二圈到达交点处 A 运动了 5 + 3 + 2 = 10
B 运动了 2 + 3 + 5 = 10
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	currentA := headA
	currentB := headB

	for currentA != nil || currentB != nil {
		if currentA == currentB {
			return currentA
		}

		if currentA == nil {
			currentA = headB
		} else {
			currentA = currentA.Next
		}

		if currentB == nil {
			currentB = headA
		} else {
			currentB = currentB.Next
		}
	}

	return nil
}
