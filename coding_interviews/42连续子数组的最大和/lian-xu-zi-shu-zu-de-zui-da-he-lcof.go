package main

func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += max(nums[i-1], 0)
	}
	return maxSlice(nums)
}

func max(num1, num2 int) int {
	if num1 < num2 {
		return num2
	}
	return num1
}

func maxSlice(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = max(res, nums[i])
	}
	return res
}

// 使用额外空间
func maxSubArrayWithExtendSpace(nums []int) int {
	length, max := len(nums), nums[0]

	dp := make([]int, length)
	dp[0] = nums[0]

	for i := 1; i < length; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
