package main

func lengthOfLongestSubstring(s string) int {
	var res, tmp int

	m := make(map[rune]int)

	for j, v := range s {
		i, ok := m[v]
		if !ok {
			i = -1
		}

		m[v] = j

		if tmp < j-i {
			tmp++
		} else {
			tmp = j - i
		}

		res = max(res, tmp)

	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
