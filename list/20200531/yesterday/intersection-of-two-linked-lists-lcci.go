// 链表相交

package main

type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}

	current1 := head1
	current2 := head2

	for current1 != nil || current2 != nil {
		if current1 == current2 {
			return current1
		}

		if current1 == nil {
			current1 = head2
		} else {
			current1 = current1.Next
		}

		if current2 == nil {
			current2 = head1
		} else {
			current2 = current2.Next
		}
	}
	return nil
}