package main

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	var dfs func(row, col, index int) bool

	dfs = func(row, col, index int) bool {
		if !(row >= 0 && row < len(board)) || !(col >= 0 && col < len(board[0])) || board[row][col] != word[index] {
			return false
		}

		if index == len(word)-1 {
			return true
		}

		tmp := board[row][col]
		board[row][col] = '/'

		res := dfs(row+1, col, index+1) || dfs(row-1, col, index+1) || dfs(row, col+1, index+1) || dfs(row, col-1, index+1)

		board[row][col] = tmp

		return res
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if dfs(row, col, 0) {
				return true
			}
		}
	}
	return false
}
