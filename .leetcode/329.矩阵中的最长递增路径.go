/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 *
 * https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/description/
 *
 * algorithms
 * Hard (38.53%)
 * Likes:    117
 * Dislikes: 0
 * Total Accepted:    8.7K
 * Total Submissions: 21.9K
 * Testcase Example:  '[[9,9,4],[6,6,8],[2,1,1]]'
 *
 * 给定一个整数矩阵，找出最长递增路径的长度。
 *
 * 对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。
 *
 * 示例 1:
 *
 * 输入: nums =
 * [
 * ⁠ [9,9,4],
 * ⁠ [6,6,8],
 * ⁠ [2,1,1]
 * ]
 * 输出: 4
 * 解释: 最长递增路径为 [1, 2, 6, 9]。
 *
 * 示例 2:
 *
 * 输入: nums =
 * [
 * ⁠ [3,4,5],
 * ⁠ [3,2,6],
 * ⁠ [2,2,1]
 * ]
 * 输出: 4
 * 解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
 *
 *
 */

// @lc code=start
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bfs(matrix [][]int, i, j, rows, cols int, dp [][]int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	val, ret := matrix[i][j], 0
	if i > 0 && matrix[i-1][j] > val {
		ret = Max(ret, bfs(matrix, i-1, j, rows, cols, dp)+1)
	}
	if i < rows-1 && matrix[i+1][j] > val {
		ret = Max(ret, bfs(matrix, i+1, j, rows, cols, dp)+1)
	}
	if j > 0 && matrix[i][j-1] > val {
		ret = Max(ret, bfs(matrix, i, j-1, rows, cols, dp)+1)
	}
	if j < cols-1 && matrix[i][j+1] > val {
		ret = Max(ret, bfs(matrix, i, j+1, rows, cols, dp)+1)
	}
	if ret == 0 {
		ret = 1
	}
	dp[i][j] = ret
	return ret
}

func longestIncreasingPath(matrix [][]int) int {
	var rows = len(matrix)
	if rows == 0 {
		return 0
	}
	var cols = len(matrix[0])
	//dp[i][j]代表以其为开头的最长路径
	var dp = make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	var ret int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ret = Max(ret, bfs(matrix, i, j, rows, cols, dp))
		}
	}
	return ret
}

// @lc code=end

