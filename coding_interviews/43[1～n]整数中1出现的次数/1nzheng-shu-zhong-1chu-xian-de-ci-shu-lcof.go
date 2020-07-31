package main

func countDigitOne(n int) int {
	digit, res := 1, 0
	hight, cur, low := n/10, n%10, 0

	for hight != 0 || cur != 0 {
		if cur == 0 {
			res += hight * digit
		} else if cur == 1 {
			res += hight*digit + low + 1
		} else {
			res += (hight + 1) * digit
		}

		low += cur * digit
		cur = hight % 10
		hight /= 10
		digit *= 10
	}
	return res
}
