/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例：
	给定数组 nums = [-1, 0, 1, 2, -1, -4]，
	满足要求的三元组集合为：
	[
		[-1, 0, 1],
		[-1, -1, 2]
	]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int

	length := len(nums)

	// 如果nums的长度小于3，直接返回
	if length < 3 {
		return result
	}

	// 将数组按照从小到大排序，go使用优化过的快排: O(nlogn)
	sort.Ints(nums)

	for i, _ := range nums {
		// 如果nums[i]大于0，则排过序数组中，后面不会有相加等于0的情况
		if nums[i] > 0 {
			return result
		}

		// 重复的元素略过，避免出现重复的解
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 定义两个指针，l = i+1; r = length-1; 当 l < r,从两边向中间移动
		l := i + 1
		r := length - 1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				// 如果找到了，就把值放到结果中
				temp := []int{nums[i], nums[l], nums[r]}
				result = append(result, temp)

				// 为了排除重复的解，就需要判断l和l的下一位，r和r的下一位是否重复
				// 如果重复，将l右移，r左移
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if nums[i]+nums[l]+nums[r] > 0 {
				// 如果和大于0，说明nums[r]太大了，左移
				r--
			} else {
				// 如果和小于0，说明nums[l]太小了，右移
				l++
			}
		}
	}

	return result
}
