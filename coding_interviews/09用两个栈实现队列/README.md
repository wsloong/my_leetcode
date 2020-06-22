## 07用两个栈实现队列

```
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1)

示例 1：
    输入：
        ["CQueue","appendTail","deleteHead","deleteHead"]
        [[],[3],[],[]]
    输出：
        [null,null,3,-1]

示例 2：
    输入：
        ["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
        [[],[],[5],[2],[],[]] 
    输出：
        [null,-1,null,null,5,2]

提示：
    1 <= values <= 10000
    最多会对 appendTail、deleteHead 进行 10000 次调用

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

### 思路

```
栈：先入后出
队列：先入先出

1，维护两个栈 A/B; A栈存储原数值，比如 [1, 2, 3]; B栈是将A栈的数据弹出在依次入B栈，则为 [3, 2, 1]
2, 加入队尾函数 appendTail(): 将数值直接如A栈即可
3，删除队首函数 deleteHead():
    如果B栈不为空，说明还有先入的数值，直接返回B栈顶的元素
    当A为空，两个栈都为空，返回-1
    否则：将A元素一次pop，然后入B栈中，实现元素倒序，并返回B栈顶元素

由于 golang 中没有实现官方的栈，这里使用slice简单模拟一下
```