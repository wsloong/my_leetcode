/*
19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：
	给定一个链表: 1->2->3->4->5, 和 n = 2.
	当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：
	给定的 n 保证是有效的。

进阶：
	你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
这个题之前刷过，这里再刷一遍思路如下
定义一个2个指针，slow, fast := head, head
遍历链表将fast移动到n的位置
遍历链表,将slow，fast一次后移一个位置，当fast到达尾部时候，slow正好在倒数n的位置
这时候执行删除操作即可
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	// 定义一个新的带头节点
	newHead := &ListNode{Next: head}

	slow, fast := newHead, newHead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}

	// 删除slow即可
	slow.Next = slow.Next.Next
	return newHead.Next
}

// ========== 早上的 反转字符串 ============

func reverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
