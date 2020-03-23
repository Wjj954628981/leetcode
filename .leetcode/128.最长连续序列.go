/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 *
 * https://leetcode-cn.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (44.62%)
 * Likes:    234
 * Dislikes: 0
 * Total Accepted:    24.6K
 * Total Submissions: 51.7K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组，找出最长连续序列的长度。
 *
 * 要求算法的时间复杂度为 O(n)。
 *
 * 示例:
 *
 * 输入: [100, 4, 200, 1, 3, 2]
 * 输出: 4
 * 解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 */

// @lc code=start
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	var length = len(nums)
	var m = make(map[int]bool, length)
	for _, v := range nums {
		m[v] = true
	}

	var max int
	for _, v := range nums {
		if _, ok := m[v-1]; !ok {
			var cnt = 1
			for true {
				if _, ok := m[v+1]; ok {
					cnt++
					v++
				} else {
					break
				}

			}
			max = Max(max, cnt)
		}
	}
	return max
}

// @lc code=end

