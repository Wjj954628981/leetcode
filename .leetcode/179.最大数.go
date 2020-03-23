/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 *
 * https://leetcode-cn.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (32.20%)
 * Likes:    210
 * Dislikes: 0
 * Total Accepted:    20K
 * Total Submissions: 56.9K
 * Testcase Example:  '[10,2]'
 *
 * 给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
 *
 * 示例 1:
 *
 * 输入: [10,2]
 * 输出: 210
 *
 * 示例 2:
 *
 * 输入: [3,30,34,5,9]
 * 输出: 9534330
 *
 * 说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
 *
 */

import (
	"sort"
	"strconv"
)

// @lc code=start
func largestNumber(nums []int) string {
	//设置capacity
	ss := make([]string, 0, len(nums))
	for _, num := range nums {
		ss = append(ss, strconv.Itoa(num))
	}
	sort.Slice(ss, func(i, j int) bool {
		s1, s2 := ss[i], ss[j]
		return (s1 + s2) > (s2 + s1)
	})

	if ss[0] == "0" {
		return "0"
	}

	var res string
	for _, s := range ss {
		res += s
	}
	return res
}

// @lc code=end

