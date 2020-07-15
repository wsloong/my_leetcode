package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var result []int

	// 特判
	if len(matrix) == 0 {
		return result
	}

	// 定义左右上下的边界
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1

	for {
		// 从左到右
		for i := l; i <= r; i++ {
			result = append(result, matrix[l][i])
		}
		t++
		if t > b {
			break
		}

		// 从上到下
		for i := t; i <= b; i++ {
			result = append(result, matrix[i][r])
		}
		r--
		if l > r {
			break
		}

		// 从右到左
		for i := r; i >= l; i-- {
			result = append(result, matrix[b][i])
		}
		b--
		if t > b {
			break
		}

		// 从下到上
		for i := b; i >= t; i-- {
			result = append(result, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}

	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println(spiralOrder(matrix))
}
