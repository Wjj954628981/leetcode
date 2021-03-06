/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 *
 * https://leetcode-cn.com/problems/rectangle-area/description/
 *
 * algorithms
 * Medium (42.35%)
 * Likes:    52
 * Dislikes: 0
 * Total Accepted:    6.5K
 * Total Submissions: 15.3K
 * Testcase Example:  '-3\n0\n3\n4\n0\n-1\n9\n2'
 *
 * 在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。
 *
 * 每个矩形由其左下顶点和右上顶点坐标表示，如图所示。
 *
 *
 *
 * 示例:
 *
 * 输入: -3, 0, 3, 4, 0, -1, 9, 2
 * 输出: 45
 *
 * 说明: 假设矩形面积不会超出 int 的范围。
 *
 */

// @lc code=start
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	lx, rx, by, ty := max(A, E), min(C, G), max(B, F), min(D, H)
	area1, area2 := (C-A)*(D-B), (G-E)*(H-F)
	if lx > rx || by > ty {
		return area1 + area2
	}
	return area1 + area2 - (rx-lx)*(ty-by)
}

// @lc code=end

