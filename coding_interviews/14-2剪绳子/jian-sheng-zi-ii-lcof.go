package main

import (
	"fmt"
)

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	a, b, p := n/3, n%3, 1000000007

	if b == 0 {
		return remainder(3, a, p)
	}

	if b == 1 {
		return (remainder(3, a-1, p) * 4) % p
	}

	return (remainder(3, a, p) * 2) % p
}

// 循环求余: 求（x ^ 3） % p 的值
func remainder(x, a, p int) int {
	result := 1
	for i := 0; i < a; i++ {
		result = (result * x) % p
	}
	return result
}

func main() {
	fmt.Println(cuttingRope(5))
}
