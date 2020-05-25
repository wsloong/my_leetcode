/*
142. 环形链表 II

给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
说明：不允许修改给定的链表。


示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 哈希表 时间复杂度：O(n); 空间复杂度：O(n)
func detectCycle(head *ListNode) *ListNode {
	// 创建一个map，用于保存记录过的节点
	// 如果这个节点再次出现，则一定形成了环，这个点就是环的入口
	visited := make(map[*ListNode]struct{})

	currentNode := head
	for currentNode != nil {
		if _, ok := visited[currentNode]; ok {
			return currentNode
		}

		visited[currentNode] = struct{}{}
		currentNode = currentNode.Next
	}
	return nil
}

// Floyd 算法
func detectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 定义一个方法用于查询是否有环,返回相遇时候的节点
	getIntersect := func(head *ListNode) *ListNode {
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

	// 找到相遇的节点
	intersectNode := getIntersect(head)
	if intersectNode == nil {
		return nil
	}

	// 找到环的入口,定义2个同速节点，第一个是链表的开头，第二个是相遇点
	ptr1 := head          // 链表的开头
	ptr2 := intersectNode // 相遇节点
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return ptr1
}
