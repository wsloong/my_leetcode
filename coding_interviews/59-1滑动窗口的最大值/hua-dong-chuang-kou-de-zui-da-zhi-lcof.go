package main

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	length := len(nums)

	if length == 0 || k < 1 || k > length {
		return result
	}

	queue := list.New()
	for i := 0; i < length; i++ {
		// 当队列不为空，如果队列尾部的元素 <= 当前元素
		// 为了维护递减原则，将尾部元素弹出
		for queue.Len() != 0 {
			// 获取最后一个元素
			lastIndex := queue.Back().Value.(int)
			if nums[lastIndex] <= nums[i] {
				queue.Remove(queue.Back())
			} else {
				break
			}
		}

		// 经过前面的循环，说明队列已经是递减了，直接把该值 push 进去
		queue.PushBack(i)

		// 如果滑块窗口已经略过了队列的头部元素
		// 将头部元素弹出
		frontIndex := queue.Front().Value.(int)
		if frontIndex == i-k {
			queue.Remove(queue.Front())
		}

		// 只要形成了大小为k的窗口，才能手机窗口的最大值，queue的头部元素
		if i >= k-1 {
			// 这里要重新获取一下头部的队列的值
			// 之前的值在`frontIndex == i - k ` 会被删除
			frontIndex = queue.Front().Value.(int)
			result = append(result, nums[frontIndex])
		}
	}
	return result
}
