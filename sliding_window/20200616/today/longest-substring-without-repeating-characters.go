/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
	输入: "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
	输入: "bbbbb"
	输出: 1
	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:
	输入: "pwwkew"
	输出: 3
	解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
	     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"log"
)

func lengthOfLongestSubstring(s string) int {
	var result int // 结果
	length := len(s)
	if length == 0 {
		return result
	}
	// 记录字符是否出现过
	visited := make(map[byte]int)
	right := -1 // 右指针，初始为-1,

	for i := 0; i < length; i++ {
		if i != 0 { // 左指针右移一次，我们就删除一个字符
			delete(visited, s[i-1])
		}

		for right+1 < length && visited[s[right+1]] == 0 {
			// 移动右指针
			visited[s[right+1]]++
			right++
		}
		result = max(result, right-i+1)
	}
	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	test := []struct {
		src    string
		result int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, item := range test {
		if res := lengthOfLongestSubstring(item.src); res != item.result {
			log.Printf("src:%s, want:%d, get:%d", item.src, item.result, res)
		}
	}
}
