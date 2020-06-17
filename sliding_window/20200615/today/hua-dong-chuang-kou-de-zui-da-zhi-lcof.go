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

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"container/list"
)

// 最优解 单调的双向队列
// 双向队列：既能从头部进行插入、删除；也可以从尾部进行杀入删除
// 单调: 队列的从头到尾所有的元素都是依次递减(递增)
// 这里使用 container/list 来维护一个递减的队列
// 特别注意: 队列中维护的是值的索引，因为通过索引可以很容易找到值，而且不重复

// 提交后发现，使用双向队列相比暴力法，用了较多的时间和内存
// 这是因为 list 的声明、节点的操作需要内存和时间
func maxSlidingWindow(nums []int, k int) []int {
	var result []int

	length := len(nums)
	if length == 0 || k < 1 || length < k {
		return result
	}

	queue := list.New()
	for i := 0; i < length; i++ {
		// 当队列不为空，如果队列尾部的元素<=当前元素
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

		// 经过前面的for，说明队列已经维持了递减原则，可以直接将该值push进去
		queue.PushBack(i)

		// 如果滑块窗口已经略过了队列中的头部元素
		// 则将头部元素弹出
		frontIndex := queue.Front().Value.(int)
		if frontIndex == i-k {
			queue.Remove(queue.Front())
		}

		// 只有形成大小为k的窗口，才能收集窗口的最大值，即最头部的元素
		if i >= k-1 {
			// 这里要重新获取一下头部的队列的值
			// 之前的值在`frontIndex == i-k`会被删除
			frontIndex = queue.Front().Value.(int)
			result = append(result, nums[frontIndex])
		}
	}
	return result
}

/*
没有关于滑块的思路，使用暴力解法
这里要维护两个变量:最大值maxVal和最大值的索引maxIndex
*/
func maxSlidingWindow2(nums []int, k int) []int {
	var result []int
	length := len(nums)
	if length == 0 || length < k {
		return result
	}

	maxIndex, maxVal := -1, nums[0]
	// 注意这里循环退出的条件 i < length-k+1
	for i := 0; i < length-k+1; i++ {
		if maxIndex >= i {
			if nums[i+k-1] > maxVal {
				maxVal = nums[i+k-1]
				maxIndex = i + k - 1
			}
		} else {
			maxVal = nums[i]
			for j := i; j < i+k; j++ {
				if maxVal < nums[j] {
					maxVal = nums[j]
					maxIndex = j
				}
			}
		}
		result = append(result, maxVal)
	}

	return result
}
