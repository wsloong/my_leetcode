package main

func lastRemaining(n int, m int) int {
	var result int
	for i := 2; i <= n; i++ {
		result = (result + m) % i
	}
	return result
}
