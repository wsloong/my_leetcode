package main

func firstUniqChar(s string) byte {
	m := make(map[byte]bool)
	length := len(s)

	for i := 0; i < length; i++ {
		v := s[i]
		if _, ok := m[v]; !ok {
			m[v] = true
		} else {
			m[v] = false
		}
	}

	for i := 0; i < length; i++ {
		if m[s[i]] {
			return s[i]
		}
	}
	return ' '
}
