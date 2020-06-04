/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// partition 定义两个哑节点，作为左右链表的头
// 依次迭代链表，小于x的放到左链表，大于x的放到有里阿尼报
// 链接左侧链表和右侧链表
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	beforeHead, afterHead := new(ListNode), new(ListNode)
	before, after := beforeHead, afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}

	// 这一步很重要，第一次没有写。特别注意
	after.Next = nil

	before.Next = afterHead.Next
	return beforeHead.Next
}

// ======== 早上的 排序链表
/*
1 查找中间节点
2 按照中间节点对链表进行分割,如果有偶数个节点，使用右节点
3 递归执行前2步，直到分割后两边只有一个节点(或者一个有一个节点，一个没有节点:nil)
4 按照大小合并节点
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next // 如果要使用偶数个节点的左节点(slow, fast := head, head)

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// 中间节点
	mid := slow.Next
	slow.Next = nil

	// 分而治之
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

	if left != nil {
		current.Next = left
	}

	if right != nil {
		current.Next = right
	}

	return newHead.Next
}
