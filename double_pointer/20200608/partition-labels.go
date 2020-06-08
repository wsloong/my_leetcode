/*
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。返回一个表示每个字符串片段的长度的列表。
示例 1:

输入: S = "ababcbacadefegdehijhklij"
输出: [9,7,8]
解释:
	划分结果为 "ababcbaca", "defegde", "hijhklij"。
	每个字母最多出现在一个片段中。
	像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。

注意:
	S的长度在[1, 500]之间。
	S只包含小写字母'a'到'z'。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-labels
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "sort"

func partitionLabels(S string) []int {
	// 定义一个map，用于记录每个字符最有一次出现的位置
	last := make(map[rune]int)
	for i, v := range S {
		last[v] = i
	}

	// start, end表示区间的首位
	var start, end int
	var result []int

	for i, v := range S {
		// 如果遇到的字符最后一次出现的位置下标大于 end， 就让 end=last[c] 来拓展当前的区间
		if last[v] > end {
			end = last[v]
		}

		// 当遍历到了当前区间的末尾时(即 i==end )，把当前区间加入答案，同时将 start 设为 i+1 去找下一个区间。
		if i == end {
			result = append(result, i-start+1)
			start = i + 1
		}
	}

	return result
}

// ====== 早上的 三数之和 =======
// 先排序数组，确定从小到大的顺序排列
// 循环数组：定义两个指针(l, r)，分别从当前位置和数组的尾部开始
// 在 l < r的条件下，一次对当前值 + l + r对应的值相加 = sum
// 如果sum=0，则依次将l和l的下一个，r和r的前一个比较，如果相同掠过
// 如果sum < 0; 说明l小了，右移
// 如果sum > 0; 说明r打了，左移

func threeSum(nums []int) [][]int {
	var result [][]int // 定义结果
	length := len(nums)
	// 如果数量小于3，直接返回
	if length < 3 {
		return result
	}

	// 先对数据进行排序，很关键
	sort.Ints(nums)

	for i := range nums {
		if nums[i] > 0 { // 如果当前值大于0，那么之后相加不会等于0
			return result
		}

		// 重复的元素略过，避免出现重复的解		// 错误1：这个没有写
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 定义两个指针，l：左侧，r:右侧，两个指针依次向中间移动
		l := i + 1
		r := length - 1

		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				temp := []int{nums[i], nums[l], nums[r]}
				result = append(result, temp)

				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return result
}
