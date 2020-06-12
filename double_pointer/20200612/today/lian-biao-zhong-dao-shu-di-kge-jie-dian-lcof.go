/*
面试题22. 链表中倒数第k个节点

输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。

示例：
	给定一个链表: 1->2->3->4->5, 和 k = 2.
	返回链表 4->5.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}

	return slow
}

// ======== 早上的 合并两个有序数组 =============

/*
思路回顾
	申请两个指针:p1, p2;分别执行nums1，nums2的尾部
	指针 p 表示 nums1和nums2的总长度
	依次从后往前对比p1,p2的值，将较大的值放到p的位置，然后按照条件操作
	如果 nums1[p1] < nums2[p2], nums1[p] = nums2[p2] p2--, p--
	最后为了防止遗漏nums2的元素，将nums2[:p2+1]的元素copy到nums1[:p2+1]位置
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			// 由于是从后往前，这里将较大的数据放到p的位置
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}

	copy(nums1[:p2+1], nums2[:p2+1])
}
