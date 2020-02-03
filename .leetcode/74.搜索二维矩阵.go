/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	column, row := len(matrix[0]), len(matrix)
	start, end := 0, column*row-1

	var getVal = func(idx int) int {
		return matrix[idx/column][idx-(idx/column)*column]
	}
	var idx, val int
	for start <= end {
		idx = start + (end-start)/2
		val = getVal(idx)
		if val == target {
			return true
		} else if val > target {
			end = idx - 1
		} else {
			start = idx + 1
		}
	}
	return false
}

// @lc code=end

