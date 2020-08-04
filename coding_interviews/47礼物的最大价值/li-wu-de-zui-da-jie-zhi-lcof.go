package main

import "fmt"

func maxValue(grid [][]int) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}
	n := len(grid[0])

	// 初始化第一行
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}

	// 初始化第一列
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += max(grid[i][j-1], grid[i-1][j])
		}
	}

	return grid[m-1][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var grid [][]int
	row1 := []int{1, 2, 5}
	row2 := []int{3, 2, 1}
	grid = append(grid, row1, row2)

	fmt.Println(grid)
	fmt.Println(maxValue(grid))
	fmt.Println(grid)
}
