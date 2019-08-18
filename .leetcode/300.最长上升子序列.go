/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 *
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (42.02%)
 * Likes:    250
 * Dislikes: 0
 * Total Accepted:    21.5K
 * Total Submissions: 50.2K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 *
 * 示例:
 *
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 *
 * 说明:
 *
 *
 * 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
 * 你算法的时间复杂度应该为 O(n^2) 。
 *
 *
 * 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
 *
 */
func lengthOfLIS(nums []int) int {
	var length = len(nums)
	if length < 2 {
		return length
	}
	var Max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dp = make([]int, length)
	dp[0] = 1
	var result = 1
	for i := 1; i < length; i++ {
		maxCount := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxCount = Max(maxCount, dp[j]+1)
			}
		}
		dp[i] = maxCount
		result = Max(result, maxCount)
	}
	return result
}

