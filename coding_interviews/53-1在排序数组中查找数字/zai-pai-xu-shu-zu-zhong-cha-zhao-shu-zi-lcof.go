package main

func search(nums []int, target int) int {
	// 搜索右边界 right
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] <= target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	right := i

	// 若数组中无 target ，则提前返回
	if j >= 0 && nums[j] != target {
		return 0
	}

	// 搜索左边界 left
	i = 0
	for i <= j {
		m := (i + j) / 2
		if nums[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	left := j
	return right - left - 1
}
