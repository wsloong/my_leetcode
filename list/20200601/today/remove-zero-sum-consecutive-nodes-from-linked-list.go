//  1171. 从链表中删去总和值为零的连续节点

/*
这个题没有思路，本来是想用双指针游走或者两层嵌套循环的，并没有解出来这个题
下面按照 `philhsu` 的解题思路，进行编码

我们可以考虑如果给的入参不是链表是数组的话，只需要求出前缀和，对于前缀和相同的项，那他们中间的部分即是可以消除掉的，
比如以 [1, 2, 3, -3, 4] 为例，其前缀和数组为 [1, 3, 6, 3, 7] ，
我们发现有两项均为 3，则 6 和 第二个 3 所对应的原数组中的数字是可以消掉的。
换成链表其实也是一样的思路，把第一个 3 的 next 指向第二个 3 的 next 即可

作者：philhsu
链接：https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/solution/c-jian-ji-dai-si-lu-by-philhsu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := new(ListNode)
	newHead.Next = head

	m := make(map[int]*ListNode)

	var sum int
	for d := newHead; d != nil; d = d.Next {
		sum += d.Val
		m[sum] = d
	}

	sum = 0
	for d := newHead; d != nil; d = d.Next {
		sum += d.Val
		d.Next = m[sum].Next
	}
	return newHead.Next
}
