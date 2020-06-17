/*
面试题59 - I. 滑动窗口的最大值

给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:
	输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
	输出: [3,3,5,5,6,7]
解释:
  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

提示：
	你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
*/
package main

import (
	"container/list"
)

/*
思路回顾
	1，定义一个单调递减的双端队列，使用container/list实现，保存的是值对应的索引
	2，如果最右边索引对应的值，小于当前值i，则将右侧的值移除，然后放入当前值
	3，如果窗口已经形成， i >= k - 1,则将队列做大值放到结果列表中
*/

func maxSlidingWindow(nums []int, k int) []int {
	var result []int

	length := len(nums)
	if length <= 0 || k <= 0 || length < k {
		return result
	}

	deque := list.New()
	for i := 0; i < length; i++ {
		for deque.Len() != 0 {
			backItem := deque.Back()
			if backItem.Value.(int) < nums[i] {
				deque.Remove(backItem)
			} else {
				break
			}
		}

		// 经过for之后，当前值索引 i 就可以直接入队列了
		deque.PushBack(i)

		// 当滑块已经过了队列的头部，则删除头部元素
		frontIndex := deque.Front().Value.(int)
		if frontIndex == i-k {
			deque.Remove(deque.Front())
		}

		// 前几个元素依次入队，还没有形成大小为k的窗口
		// 是否形成了大小为k的窗口
		if i >= k-1 {
			frontIndex := deque.Front().Value.(int)
			result = append(result, nums[frontIndex])
		}
	}
	return result
}
