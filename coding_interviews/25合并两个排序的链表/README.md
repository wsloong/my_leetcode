### 25合并两个排序的链表

```
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

#### 思路

主要使用`带头空节点: newHead`，作为结果的头节点，返回时候只需要返回`newHead.next`即可
* 执行循环，由于循环中需要不短的后移节点，定义三个临时变量`current1, current2, cuurrent = l1, l2, newHead`; 
* 循环条件 `current1 != nil && current2 != nil`
* 循环过称依次比较；当`current1.val <= current2.val`时候，`current.next = current1`；反之 `current.next = current2`
* 最后将`current1`或者`current2`遗留的节点，放到`current`后面即可

T: O(m+n): 其中 `n` 和 `m` 分别为两个链表的长度
S: O(1)