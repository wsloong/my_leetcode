package main

import "fmt"

func isNumber(s string) bool {
	// 定义状态机
	states := []map[rune]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4}, // 0: start with 'blank'
		{'d': 2, '.': 4},                 // 1: 'sign' before 'e'
		{'d': 2, '.': 3, 'e': 5, ' ': 8}, // 2: 'dight' before 'dot'
		{'d': 3, 'e': 5, ' ': 8},         // 3: 'dight' after 'dot'
		{'d': 3},                         // 4: 'dight' after 'dot' ('blank' before 'dot')
		{'s': 6, 'd': 7},                 // 5: 'e'
		{'d': 7},                         // 6: 'sign' after 'e'
		{'d': 7, ' ': 8},                 // 7: 'digit' after 'e'
		{' ': 8},                         // 8: end with 'blank'
	}

	var p int

	for _, c := range s {
		t := '?'

		if c >= '0' && c <= '9' {
			t = 'd' // digit
		} else if c == '+' || c == '-' {
			t = 's' // sign
		} else if c == '.' || c == 'e' || c == 'E' {
			t = c // dot, e, blank
		}

		// 不存在 返回false
		if _, ok := states[p][t]; !ok {
			return false
		}

		// 重新赋值
		p = states[p][t]
	}

	return p == 2 || p == 3 || p == 7 || p == 8
}

func main() {
	fmt.Println(isNumber("1")) // 提交时候，官网显示这个输出为false???
}
