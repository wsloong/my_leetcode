// 148. 排序链表
/*
在 O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 使用双指针查找中间节点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	mid := slow.Next // 中间节点
	slow.Next = nil

	// 递归
	// 最后left, right只剩下一个节点了
	left, right := sortList(head), sortList(mid)

	newHead := new(ListNode)
	current := newHead
	for left != nil && right != nil {
		if left.Val < right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}

	// 如果是原链表有奇数个节点
	// 最后分割的左右链表中，会有一个链表有节点，另一个没有节点
	if left != nil {
		current.Next = left
	}

	if right != nil {
		current.Next = right
	}

	return newHead.Next
}
