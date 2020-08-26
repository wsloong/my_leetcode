package main

import (
	"fmt"
	"math"
)

func twoSum(n int) []float64 {
	var dp [70]float64
	for i := 1; i <= 6; i++ {
		dp[i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 6 * i; j >= i; j-- {
			dp[j] = 0
			for k := 1; k <= 6; k++ {
				if j-k < i-1 {
					break
				}
				dp[j] += dp[j-k]
			}
		}
	}

	all := math.Pow(6, float64(n))

	var result []float64
	for i := n; i <= 6*n; i++ {
		result = append(result, dp[i]*1.0/all)
	}
	return result
}

func main() {
	fmt.Println(twoSum(2))
}
