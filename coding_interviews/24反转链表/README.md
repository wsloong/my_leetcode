### 24反转链表

```
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点

示例:
    输入: 1->2->3->4->5->NULL
    输出: 5->4->3->2->1->NULL

限制：
0 <= 节点个数 <= 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```


#### 思路

* 使用双指针: `prev`: 标记上一个节点，默认为`nil`； `current`: 标记当前节点
* 执行循环，条件：`current != nil`
    * 循环中，可以获取当前节点`current`, 下一个节点`temp = current.next`和上一个节点`prev`
    * 修改节点指向
    ```
    current.next = prev
    prev = current
    current = temp
    ```

T: O(n)
S: O(1)