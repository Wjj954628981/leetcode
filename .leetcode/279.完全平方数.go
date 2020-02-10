/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 *
 * https://leetcode-cn.com/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (49.91%)
 * Likes:    301
 * Dislikes: 0
 * Total Accepted:    35.2K
 * Total Submissions: 66.1K
 * Testcase Example:  '12'
 *
 * 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
 *
 * 示例 1:
 *
 * 输入: n = 12
 * 输出: 3
 * 解释: 12 = 4 + 4 + 4.
 *
 * 示例 2:
 *
 * 输入: n = 13
 * 输出: 2
 * 解释: 13 = 4 + 9.
 *
 */
import "math"

// @lc code=start
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	cnt := int(math.Sqrt(float64(n)))
	var l = make([]int, 0, cnt)
	for i := 1; i <= cnt; i++ {
		l = append(l, i*i)
	}
	return a(n, l)
}

var m = make(map[int]int)

func a(n int, l []int) int {
	if val, ok := m[n]; ok {
		return val
	}
	var ret = math.MaxInt8
	for _, val := range l {
		if val <= n {
			if n-val == 0 {
				m[n] = 1
				return 1
			} else {
				ret = Min(ret, a(n-val, l)+1)
			}
		}
	}
	m[n] = ret
	return ret
}

// @lc code=end

