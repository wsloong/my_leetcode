package main

func reverseLeftWords1(s string, n int) string {
	return s[n:] + s[:n]
}

func reverseLeftWords2(s string, n int) string {
	var res []byte

	for i := n; i < n+len(s); i++ {
		res = append(res, s[i%len(s)])
	}

	return string(res)
}
