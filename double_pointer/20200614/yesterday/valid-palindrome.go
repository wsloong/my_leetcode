package main

import "strings"

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
