package main

func numWays(n int) int {
	a, b, sum := 1, 1, 0

	for i := 0; i < n; i++ {
		sum = (a + b) % 1000000007
		a, b = b, sum
	}
	return a
}

// 动态规划
func numWaysWithDP(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}
