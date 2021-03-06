/*
面试题59 - II. 队列的最大值
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：
	输入:
		["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
		[[],[1],[2],[],[],[]]
	输出: [null,null,null,2,1,2]

示例 2：
	输入:
		["MaxQueue","pop_front","max_value"]
		[[],[],[]]
	输出: [null,-1,-1]

限制：
1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "container/list"

/*
思路：
	1,将数据队列Queue交给list去维护，这个队列只是简单的进行push和pop即可
	2,另外一个需要维护单调双端递减的队列

*/
type MaxQueue struct {
	deque *list.List
	Queue *list.List
}

func Constructor() MaxQueue {
	return MaxQueue{
		deque: list.New(),
		Queue: list.New(),
	}
}

func (this *MaxQueue) Max_value() int {
	if this.deque.Len() == 0 {
		return -1
	}
	result := this.deque.Front().Value.(int)
	return result
}

func (this *MaxQueue) Push_back(value int) {
	for this.deque.Len() != 0 {
		lastItem := this.deque.Back()
		lastValue := lastItem.Value.(int)
		if lastItem < value {
			this.deque.Remove(lastItem)
		} else {
			break
		}
	}
	this.deque.PushBack(value)
	this.Queue.PushBack(value)
}

func (this *MaxQueue) Pop_front() int {
	if this.deque.Len() == 0 {
		return -1
	}
	ans := this.Queue.Front().Value.(int)
	if ans == this.deque.Front().Value.(int) {
		this.deque.Remove(this.deque.Front())
	}
	this.Queue.Remove(this.Queue.Front())
	return ans
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
