/*
930. 和相同的二元子数组
在由若干 0 和 1  组成的数组 A 中，有多少个和为 S 的非空子数组。

示例：
	输入：A = [1,0,1,0,1], S = 2
	输出：4
	解释：
		如下面黑体所示，有 4 个满足题目要求的子数组：
		[1,0,1,0,1]
		[1,0,1,0,1]
		[1,0,1,0,1]
		[1,0,1,0,1]
提示：
	A.length <= 30000
	0 <= S <= A.length
	A[i] 为 0 或 1
*/
package main

func numSubarraysWithSum(A []int, S int) int {
	length := len(A)
	if length < 1 {
		return 0
	}

	var n, sum, current int
	for i := 0; i < length; i++ {
		sum += A[i]
		if sum == S {
			for j := i + 1; j < length; j++ {
				sum += A[j]
				if sum == S {
					n++
				}
			}
			n++
			i = current
			current++
			sum = 0
		}
	}
	return n
}
