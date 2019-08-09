import "math"

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */
func minPathSum(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	var cache [][]int
	for i := 0; i < row; i++ {
		tmp := make([]int, column)
		cache = append(cache, tmp)
	}
	var helper func(i, j int) int
	helper = func(i, j int) int {
		if cache[i][j] != 0 {
			return cache[i][j]
		}
		if i == row-1 && j == column-1 {
			return grid[i][j]
		}
		if i == row-1 {
			return helper(i, j+1) + grid[i][j]
		}
		if j == column-1 {
			return helper(i+1, j) + grid[i][j]
		}

		ret1 := helper(i+1, j) + grid[i][j]
		ret2 := helper(i, j+1) + grid[i][j]
		result := int(math.Min(float64(ret1), float64(ret2)))
		cache[i][j] = result
		return result
	}
	return helper(0, 0)
}
