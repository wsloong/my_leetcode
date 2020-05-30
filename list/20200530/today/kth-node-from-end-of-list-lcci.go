//  返回倒数第 k 个节点
/*
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：

给定的 k 保证是有效的。
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用双指针游走，妙妙妙！！！！
/*
先将快指针游走到k的位置
然后两个指针以此往后，快指针到达尾部时候，慢指针正好在倒数k的位置
简直妙妙妙啊！！！
*/
func kthToLast(head *ListNode, k int) int {
	fast := head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		head = head.Next
	}
	return head.Val
}

// 使用了额外的空间，非最优解
// func kthToLast(head *ListNode, k int) int {
// 	NodeMap := make(map[int]*ListNode)

// 	var count int
// 	current := head
// 	for current != nil {
// 		NodeMap[count] = current
// 		current = current.Next
// 		count++
// 	}

// 	return NodeMap[count-k].Val
// }

// 先反转链表，在循环，非最优解
// func kthToLast(head *ListNode, k int) int {
// if head == nil {
// 	return -1
// }

// reversedList := reverseList(head)

// var count int
// for reversedList != nil {
// 	count++
// 	if count == k {
// 		return reversedList.Val
// 	}
// 	reversedList = reversedList.Next
// }
// return -1
// }

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	var prev *ListNode
// 	current := head

// 	for current != nil {
// 		temp := current.Next
// 		current.Next = prev
// 		prev = current
// 		current = temp
// 	}
// 	return prev
// }

// ==============早上的 // 链表相交 ==================

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	currentA, currentB := headA, headB

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
