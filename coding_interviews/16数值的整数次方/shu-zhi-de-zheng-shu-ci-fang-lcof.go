package main

import "fmt"

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	// 初始化结果
	var res float64 = 1

	// 如果 n < 0  x的n次方 等价于 1/x的-n次方
	if n < 0 {
		x, n = 1/x, -n
	}

	for n > 0 {
		// 如果n是奇数
		if n&1 == 1 {
			res *= x
		}
		// 因为在进行 x *= x 时，生成的就是 x^1, x^2, x^4 的序列
		x *= x

		n >>= 1	// 右移一位，相当于n//2
	}

	return res
}

func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
}