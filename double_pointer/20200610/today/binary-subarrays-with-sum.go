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

// 没有思路
// 官方的方法很多关于数学，奈何能看明白，但是写不出来
// 先抄代码
package main

// 三指针方法
func numSubarraysWithSum(A []int, S int) int {
	iLow, iHight, sumLow, sumHight, result := 0, 0, 0, 0, 0

	for i, v := range A {
		sumLow += v
		for iLow < i && sumLow > S {
			sumLow -= A[iLow]
			iLow++
		}

		sumHight += v
		for iHight < i && (sumHight > S || sumHight == S && A[iHight] == 0) {
			sumHight -= A[iHight]
			iHight++
		}

		if sumLow == S {
			result += iHight - iLow + 1
		}
	}
	return result
}

// ======= 早上的 面试题 17.11. 单词距离 =======

func findClosest2(words []string, word1 string, word2 string) int {
	t1, t2, res := -1, -1, len(words)

	for i := range words {
		if words[i] == word1 {
			t1 = i
		} else if words[i] == word2 {
			t2 = i
		}

		if t1 != -1 && t2 != -1 {
			res = min(res, abs(t1, t2))
		}

		if res == 1 {
			return res
		}
	}

	return res
}
