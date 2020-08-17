package main

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		res := nums[l] + nums[r]
		if res == target {
			break
		}

		if res < target {
			l++
		} else {
			r--
		}
	}
	return []int{nums[l], nums[r]}
}
