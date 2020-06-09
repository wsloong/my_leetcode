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
*/

/*
思路回顾
1 特殊判断，如果数组元素小于三个直接返回
2 对原数组进行排序，方便去重的同时，方便数据分组
3 循环数组
	* 如果当前 i 对应的数组大于 0，那么之后三个数之和必不为0，直接返回
	* 定义快慢双指针 l/r，如果 sum = nums[i] + nums[j] + nums[l]的值 == 0，存入结果
	* 如果 sum < 0,说明 l 小了，右移
	* 如果 sum > 0,说明 r 打了，左移
*/

package main

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	length := len(nums)
	if length < 3 {
		return result
	}

	sort.Ints(nums)

	for i := range nums {
		if nums[i] > 0 {
			return result
		}

		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := length - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
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
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return result
}
