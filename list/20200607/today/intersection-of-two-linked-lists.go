/*
编写一个程序，找到两个单链表相交的起始节点。

注意：

	如果两个链表没有交点，返回 null.
	在返回结果后，两个链表仍须保持原有的结构。
	可假定整个链表结构中没有循环。
	程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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

	currentA, currentB := headA, headB

	for currentA != currentB {
		if currentA != nil {
			currentA = currentA.Next
		} else {
			currentA = headB
		}

		if currentB != nil {
			currentB = currentB.Next
		} else {
			currentB = headA
		}
	}

	return currentA
}
