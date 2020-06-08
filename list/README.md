## 链表专题的总结

### 为什么开始于链表
    链表是一个很有意思的数据结构；除此之外，相比于字符串/数组/排序，我在开发中很少使用这个数据结构，链表中有很多的技巧

### 链表的技巧总结

#### 反转链表

```
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		temp := curr.Next       // 这里使用临时变量是，curr.next会重新赋值，如不就会丢失这个值
		curr.Next = prev
		prev = curr
		curr = temp 
	}
	return prev
}
```

代表题目:
    [206.链表的反转](https://leetcode-cn.com/problems/reverse-linked-list/)

#### 双指针
```
// 这样结束，如果是偶数个节点，slow=右节点
// slow, fast := head, head.Next; 这样slow=左节点
slow, fast := head, head        

for fast.Next != nil && fast.Next.Next != nil{
    slow, fast = slow.Next, fast.Next.Next
}
```
代表题目:
    [876.链表的中间结点](https://leetcode-cn.com/problems/middle-of-the-linked-list/)
    [143.重排链表](https://leetcode-cn.com/problems/reorder-list/)
    [141.环形链表](https://leetcode-cn.com/problems/copy-list-with-random-pointer/)
    [面试题 02.02.返回倒数第K个节点](https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/)

#### 数学知识的应用

代表题目:
    [142.环形链表II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)
    [面试题 02.02.返回倒数第K个节点](https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/)
    [1171.从链表中删去总和值为零的连续节点](https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/)


### 需要多刷几次的题目

除了上面的代表题目还有:
    [25.K个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)
    [82.删除排序链表中的重复元素II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)
    [1367.二叉树中的列表](https://leetcode-cn.com/problems/linked-list-in-binary-tree/)
    [109.有序链表转换二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/)
    [148.排序链表](https://leetcode-cn.com/problems/sort-list/)
    [92.反转链表II](https://leetcode-cn.com/problems/reverse-linked-list-ii/)