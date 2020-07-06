package main

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// 匹配函数： 当 `y` 是 `.` 或者 `x` 和 `y` 本身相同时，这两个字符才会匹配
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}

		if p[j-1] == '.' {
			return true
		}

		return s[i-1] == p[j-1]
	}

	// 初始化 二维数组
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}

	// 当 s/p为空时候，为true
	f[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
