/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有K个重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/description/
 *
 * algorithms
 * Medium (39.52%)
 * Likes:    104
 * Dislikes: 0
 * Total Accepted:    6.6K
 * Total Submissions: 16K
 * Testcase Example:  '"aaabb"\n3'
 *
 * 找到给定字符串（由小写字符组成）中的最长子串 T ， 要求 T 中的每一字符出现次数都不少于 k 。输出 T 的长度。
 *
 * 示例 1:
 *
 *
 * 输入:
 * s = "aaabb", k = 3
 *
 * 输出:
 * 3
 *
 * 最长子串为 "aaa" ，其中 'a' 重复了 3 次。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * s = "ababbc", k = 2
 *
 * 输出:
 * 5
 *
 * 最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
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

func longestSubstring(s string, k int) int {
	var length = len(s)
	if length == 0 || k > length {
		return 0
	}
	if k < 2 {
		return length
	}
	//统计频率
	var freq = make(map[uint8]int, 26)
	for i := 0; i < length; i++ {
		freq[s[i]]++
	}

	start, end := 0, length-1
	//去除首末<k的元素
	for freq[s[start]] < k && end-1 > start {
		start++
	}

	for freq[s[end]] < k && end-1 > start {
		end--
	}

	//分治
	for i := start; i <= end; i++ {
		if freq[s[i]] < k {
			return Max(longestSubstring(s[start:i], k), longestSubstring(s[i+1:end+1], k))
		}
	}
	return end - start + 1
}

// @lc code=end

