package main

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	length := len(s)
	i, j := length-1, length-1

	var res []string
	for i >= 0 {
		// 从后往前搜索第一个空格
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res = append(res, s[i+1:j+1])

		// 过滤掉多余的空格
		for i >= 0 && s[i] == ' ' {
			i--
		}

		// 调整j的顺序，重新开始查找下一个单词
		j = i
	}

	return strings.Join(res, " ")
}

// 内置函数法
func reverseWordsWithIn(s string) string {
	s = strings.Trim(s, " ")
	strs := strings.Split(s, " ")

	var tmp []string
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] == "" {
			continue
		}
		tmp = append(tmp, strs[i])
	}

	return strings.Join(tmp, " ")
}
