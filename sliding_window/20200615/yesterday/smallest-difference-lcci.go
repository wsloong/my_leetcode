/*
面试题 16.06. 最小差
给定两个整数数组a和b，计算具有最小差绝对值的一对数值（每个数组中取一个值），并返回该对数值的差

*/
package main

import (
	"math"
	"sort"
)

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	n, m := len(a), len(b)

	i, j, result := 0, 0, math.MaxInt64
	for i < n && j < m {
		if a[i] == b[j] {
			return 0
		}

		result = min(result, abs(a[i], b[j]))
		if a[i] > b[j] {
			j++
		} else {
			i++
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
