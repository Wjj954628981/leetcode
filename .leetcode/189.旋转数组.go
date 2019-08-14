/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */
func rotate(nums []int, k int) {
	var reverse = func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	length := len(nums)
	k = k % length
	tail := length - 1
	if k == 0 {
		return
	}
	reverse(0, tail)
	reverse(0, k-1)
	reverse(k, tail)
}

