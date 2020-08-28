package main

func add(a int, b int) int {
	for b != 0 { // 当进位为0时候跳出
		c := (a & b) << 1 // c： 进位
		a ^= b            // a: 非进位和
		b = c             // b = 进位
	}
	return a
}
