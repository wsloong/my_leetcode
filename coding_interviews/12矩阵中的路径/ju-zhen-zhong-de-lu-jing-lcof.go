package main

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	var dfs func(row, col, index int) bool

	// 深度优先算法
	dfs = func(row, col, index int) bool {
		// 排除越界和值不等的情况
		if !(row >= 0 && row < len(board)) || !(col >= 0 && col < len(board[0])) || board[row][col] != word[index] {
			return false
		}

		// 如果已经全部找到了就返回true
		if index == len(word)-1 {
			return true
		}

		// 当前矩阵中的值修改为 '/'
		tmp := board[row][col]
		board[row][col] = '/'

		// 递归调用
		res := dfs(row+1, col, index+1) || dfs(row-1, col, index+1) || dfs(row, col+1, index+1) || dfs(row, col-1, index+1)

		// 恢复矩阵中的值
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
