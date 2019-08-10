/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */
func singleNumber(nums []int) int {
	var result int
	for k, v := range nums {
		if k == 0 {
			result = v
		} else {
			result ^= v
		}
	}
	return result
}

