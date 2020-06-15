/*
76. 最小覆盖子串
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。
*/
package main

import (
	"math"
)

// 再抄一遍代码 ^-^
// 难哦~
func minWindow(s string, t string) string {
	// ori:保存t字符和字符出现的次数
	// cnt: 保存包含t字符和字符出现的次数
	ori, cnt := map[byte]int{}, map[byte]int{}

	// 将t中的字符放到ori中，为什么要保存次数，这是因为t中可能出现重复的字符
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	// 结果为s[ansL: ansR];sLen:s的长度
	ansL, ansR, sLen, length := -1, -1, len(s), math.MaxInt32

	// 定义一个函数，用于验证cnt中是否包含了全部的s
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}

		// cnt中包含了所有的ori，就依次右移l
		for check() && l <= r {
			if r-l+1 < length {
				length = r - l + 1
				ansL, ansR = l, l+length
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
