// 328. 奇偶链表

/*
定义奇偶两个链表
将奇数节点放到奇链表
偶数节点放到偶链表
把偶数节点链接到奇数节点后面
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd, even := head, head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}

// ----------------早上的[环形链表 II]---------------------

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 找到相遇节点
	intersectNode := getIntersect(head)
	if intersectNode == nil {
		return nil
	}

	// 从链表开头到环入口 = 相遇节点到环入口
	ptr1 := head
	ptr2 := intersectNode

	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1
}

// 找相遇节点,如果为空表示没有相遇节点，没有环，直接返回即可
func getIntersect(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

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
