package main

import (
	"fmt"
	"strings"
)

func permutation(s string) []string {
	var result []string
	c := strings.Split(s, "")

	var dfs func(x int)
	dfs = func(x int) {
		if x == len(c)-1 {
			result = append(result, strings.Join(c, ""))
			return
		}

		var m = make(map[string]bool)
		for i := x; i < len(c); i++ {
			if m[c[i]] {
				continue
			}

			m[c[i]] = true

			c[i], c[x] = c[x], c[i]
			dfs(x + 1)
			c[i], c[x] = c[x], c[i]
		}

	}
	dfs(0)
	return result
}

func main() {
	fmt.Println(permutation("abc"))
}
