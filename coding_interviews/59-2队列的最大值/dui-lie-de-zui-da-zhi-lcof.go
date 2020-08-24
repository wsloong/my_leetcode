package main

import (
	"container/list"
	"fmt"
)

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
		lastItem := this.deque.Back().Value.(int)
		if lastItem < value {
			this.deque.Remove(this.deque.Back())
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

func main() {
	obj := Constructor()
	param1 := obj.Max_value()
	obj.Push_back(3)
	param2 := obj.Pop_front()

	fmt.Println(param1, param2)
}
