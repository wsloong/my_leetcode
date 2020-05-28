// 142. 环形链表 II

// 使用map保存已经访问过的节点

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]*ListNode)

	curr := head
	for curr != nil {
		if n, ok := visited[curr]; ok {
			return n
		}

		visited[curr] = curr
		curr = curr.Next
	}
	return nil
}

//----------------------最优解--------------------

/*
利用了出发点到环入口 和 环入口到环中快慢节点相遇处路径相同
*/
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