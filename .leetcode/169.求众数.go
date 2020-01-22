/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 求众数
 * 投票法
 */
func majorityElement(nums []int) int {
	var count = 0
	var candidate = 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count += 1
		} else {
			count -= 1
		}
	}
	return candidate
}

