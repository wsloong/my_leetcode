/*
面试题 16.06. 最小差
给定两个整数数组a和b，计算具有最小差绝对值的一对数值（每个数组中取一个值），并返回该对数值的差

示例：

	输入：{1, 3, 15, 11, 2}, {23, 127, 235, 19, 8}
	输出： 3，即数值对(11, 8)

提示：
1 <= a.length, b.length <= 100000
-2147483648 <= a[i], b[i] <= 2147483647
正确结果在区间[-2147483648, 2147483647]内

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/smallest-difference-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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

func min(a, b int) {
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
