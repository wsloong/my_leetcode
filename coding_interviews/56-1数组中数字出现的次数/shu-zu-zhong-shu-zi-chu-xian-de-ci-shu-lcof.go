package main

func singleNumbers(nums []int) []int {
	// result: 所有数字按位异或的结果
	// 这个结果等于 2个出现一次的数字的异或值
	var result int
	for _, num := range nums {
		result ^= num
	}

	// div: 找到任意为1的位
	div := 1
	for div&result == 0 {
		div <<= 1
	}

	// 根据 div的位，对数字进行分组
	// 每个组内进行异或操作，得到2个数字
	a, b := 0, 0
	for _, num := range nums {
		if div&num == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}

}
