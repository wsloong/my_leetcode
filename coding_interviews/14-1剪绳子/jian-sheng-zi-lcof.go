package main

import (
	"fmt"
	"math"
)

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	// 除数和余数
	a, b := n/3, n%3

	if b == 0 {
		return int(math.Pow(3, float64(a)))
	}

	if b == 1 {
		return int(math.Pow(3, float64(a-1))) * 4
	}

	return int(math.Pow(3, float64(a))) * 2
}

func main() {
	fmt.Println(cuttingRope(10))
}
