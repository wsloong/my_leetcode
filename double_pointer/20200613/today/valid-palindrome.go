/*
125. 验证回文串

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
	输入: "A man, a plan, a canal: Panama"
	输出: true

示例 2:
	输入: "race a car"
	输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	lowerS := strings.ToLower(s)

	var temp string
	for i := range lowerS {
		if isNum(lowerS[i]) || isAlpha(lowerS[i]) {
			temp += string(lowerS[i])
		}
	}

	if temp == "" {
		return true
	}

	left, right := 0, len(temp)-1
	for left < right {
		leftVal, rightVal := temp[left], temp[right]
		if leftVal != rightVal {
			return false
		}

		left++
		right--
	}
	return true
}

func isNum(b byte) bool {
	if b < 48 || b > 57 {
		return false
	}

	return true
}

func isAlpha(b byte) bool {
	if b < 97 || b > 122 {
		return false
	}
	return true
}

// ======= 早上的 面试题 10.09. 排序矩阵查找 =====
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1
	for row != len(matrix) && col != -1 {
		if matrix[row][col] > target {
			col--
		} else if matrix[row][col] < target {
			row++
		} else {
			return true
		}
	}
	return false
}
