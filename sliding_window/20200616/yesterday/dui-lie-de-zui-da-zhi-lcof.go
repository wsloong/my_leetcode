/*
面试题59 - II. 队列的最大值
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
若队列为空，pop_front 和 max_value 需要返回 -1
*/
package main

import "container/list"

type MaxQueue struct {
	deque *list.List
	queue *list.List
}

func Constructor() MaxQueue {
	return MaxQueue{
		deque: list.New(),
		queue: list.New(),
	}
}

func (this *MaxQueue) Max_value() int {
	if this.deque.Len() == 0 {
		return -1
	}

	value := this.deque.Front().Value.(int)
	return value

}

func (this *MaxQueue) Push_back(value int) {
	for this.deque.Len() != 0 {
		// 获取队列中最后一个值，如果比当前值小，则删除
		lastItem := this.deque.Back()
		if lastItem.Value.(int) < value {
			this.deque.Remove(lastItem)
		} else {
			break
		}
	}
	// 经过上面的for，deque中的所有值都是单调递增的
	this.deque.PushBack(value)
	this.queue.PushBack(value)
}

func (this *MaxQueue) Pop_front() int {
	if this.deque.Len() == 0 {
		return -1
	}

	ans := this.queue.Front().Value.(int)
	if ans == this.deque.Front().Value.(int) {
		this.deque.Remove(this.deque.Front())
	}
	this.queue.Remove(this.queue.Front())
	return ans
}
