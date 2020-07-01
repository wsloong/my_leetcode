package main

func hammingWeightMe(num uint32) int {
	var result int
	for num != 0 {
		t := num % 2
		if t == 1 {
			result++
		}
		num = num / 2
	}

	return result
}

// 逐位判断法
func hammingWeightLoop(num uint32) int {
	var result uint32
	for num != 0 {
		result += num & 1
		num >>= 1
	}
	return int(result)
}

// 最优解，巧用 n&(n-1)
func hammingWeight(num uint32) int {
	var result int
	for num != 0 {
		result++
		num &= num - 1
	}
	return result
}
