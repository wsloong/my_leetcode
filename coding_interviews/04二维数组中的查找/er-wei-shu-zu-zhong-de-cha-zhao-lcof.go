package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 特殊判断
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] > target {
			col--
		} else if matrix[row][col] < target {
			row++
		} else {
			return true
		}
	}
	return false
}

// 从左下角开始判断
func findNumberIn2DArrayFromLeftBottom(matrix [][]int, target int) bool {
	// 特殊判断
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := len(matrix)-1, 0
	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] < target {
			col++
		} else if matrix[row][col] > target {
			row--
		} else {
			return true
		}
	}
	return false
}
