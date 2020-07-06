package main

// 使用临时空间的暴力法
func exchangeWithTemp(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1

	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 1 { // 奇数
			result[left] = nums[i]
			left++
		} else {
			result[right] = nums[i]
			right--
		}
	}

	return result
}

// 使用头尾双指针法
func exchangeWithLeftRightPoint(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1

	for left < right {
		for (left < right) && (nums[left]&1) == 1 {
			left++
		}

		for (left < right) && (nums[right]&1) == 0 {
			right--
		}

		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return nums
}

// 使用快慢双指针法
func exchangeWithSlowFast(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast]&1 == 1 {
			if slow != fast {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow++
		}
		fast++
	}
	return nums
}
