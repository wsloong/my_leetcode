package main

func minArray(numbers []int) int {
	i, j := 0, len(numbers)-1

	for i < j {
		m = (i + j) / 2
		// 说明旋转节点在右边
		if numbers[m] > numbers[j] {
			i = m + 1
		} else if numbers[m] < numbers[j] { // 说明旋转节点在左边
			j = m
		} else { // 中间节点 = 右侧节点，则让右侧左移
			j--
		}
	}
	return numbers[i]
}
