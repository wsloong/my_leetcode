package main

func findContinuousSequence(target int) [][]int {
	// 定义滑动窗口的左右边界，这里的滑动窗口使用左闭右开的范围，即[)
	i, j, sum := 1, 1, 0

	var result [][]int

	// 因为至少含有两个数，
	// 所以最大的边界应该是一半
	for i <= target/2 {
		// 值小于target，右边界向右+1
		if sum < target {
			sum += j
			j++
		} else if sum > target { // 左边界向右+1
			sum -= i
			i++
		} else {
			// 将值记录到结果中
			var tmp []int
			for k := i; k < j; k++ {
				tmp = append(tmp, k)
			}
			result = append(result, tmp)

			// 将结果放入后，需要将左边界的值去掉，继续往下找结果
			sum -= i
			i++
		}
	}
	return result
}
