package main

func findRepeatNumber(nums []int) int {
	// 定义一个map，用于存储数组中数字
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		// 在map中出现过，说明重复了，直接返回该值
		if m[nums[i]] > 0 {
			return nums[i]
		}
		// 没有出现过，就将该值放到map中
		m[nums[i]]++
	}

	return -1
}
