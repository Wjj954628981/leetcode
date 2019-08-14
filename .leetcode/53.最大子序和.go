/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	max, sum := nums[0], nums[0]
	for i := 1; i < length; i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		max = Max(max, sum)
	}
	return max
}

