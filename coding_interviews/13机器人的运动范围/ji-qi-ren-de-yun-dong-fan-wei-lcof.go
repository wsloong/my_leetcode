package main

import "fmt"

// 辅助函数，求数位和
func Solution(x int) int {
	var result int
	for x != 0 {
		result += x % 10
		x /= 10
	}
	return result
}

// 深度优先算法
func movingCountWithDFS(m int, n int, k int) int {
	var dfs func(i, j int) int

	visited := make(map[string]bool)
	dfs = func(i, j int) int {
		key := fmt.Sprintf("%d-%d", i, j)

		if i >= m || j >= n || visited[key] || Solution(i)+Solution(j) > k {
			return 0
		}
		visited[key] = true

		// 从(0, 0)开始，只向下，向右搜索
		return 1 + dfs(i+1, j) + dfs(i, j+1)
	}

	return dfs(0, 0)
}

// 广度优先算法
func movingCountWithBFS(m int, n int, k int) int {
	// 从 (0, 0) 开始, 将 (0, 0)放到队列中
	var queue []point
	queue = append(queue, point{0, 0})

	// 记录访问过的元素
	visited := make(map[point]bool)

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] || current.row >= m || current.col >= n || Solution(current.row)+Solution(current.col) > k {
			continue
		}
		visited[current] = true
		queue = append(queue, point{current.row + 1, current.col}, point{current.row, current.col + 1})
	}

	return len(visited)
}

type point struct {
	row, col int
}
