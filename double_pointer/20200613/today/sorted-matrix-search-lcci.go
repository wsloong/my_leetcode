/*
面试题 10.09. 排序矩阵查找

给定M×N矩阵，每一行、每一列都按升序排列，请编写代码找出某元素。

示例:
	现有矩阵 matrix 如下：
	[
		[1, 4, 7, 11, 15],
		[2, 5, 8, 12, 19],
		[3, 6, 9, 16, 22],
		[10, 13, 14, 17, 24],
		[18, 21, 23, 26, 30]
	]

	给定 target = 5，返回 true。
	给定 target = 20，返回 false。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sorted-matrix-search-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// 从右上角的元素出发
	// 如果当前元素>目标值，说明这一列都大于目标值，向左移一列
	// 如果当前元素<目标值，说明这一列其他的元素有可能是目标值，向下移动一列
	// 找到目标元素或者行/列越界了
	row, col := 0, len(matrix[0]) - 1
	for row != len(matrix) && col != -1 {
		if matrix[row][col] > target{
			col--
		} else if matrix[row][col] < target {
			row++
		} else {
			return true
		}
	}
	return false
}	