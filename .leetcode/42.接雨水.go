/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (44.38%)
 * Likes:    529
 * Dislikes: 0
 * Total Accepted:    22.4K
 * Total Submissions: 48.9K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 *
 * 示例:
 *
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 *
 */
func trap(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	var max = func(x int, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var min = func(x int, y int) int {
		if x < y {
			return x
		}
		return y
	}
	leftMaxArr, rightMaxArr := make([]int, length), make([]int, length)
	leftMaxArr[0] = height[0]
	rightMaxArr[length-1] = height[length-1]
	for i := 1; i < length-1; i++ {
		leftMaxArr[i] = max(height[i], leftMaxArr[i-1])
	}
	for i := length - 2; i >= 0; i-- {
		rightMaxArr[i] = max(height[i], rightMaxArr[i+1])
	}
	var result int
	for i := 1; i < length-1; i++ {
		result += min(leftMaxArr[i], rightMaxArr[i]) - height[i]
	}
	return result
}

