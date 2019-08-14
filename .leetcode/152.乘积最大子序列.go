import "math"

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 */
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func maxProduct(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	max := math.MinInt8
	iMax, iMin := 1, 1
	for i := 0; i < length; i++ {
		num := nums[i]
		if nums[i] < 0 {
			iMax, iMin = iMin, iMax
		}
		iMax = Max(iMax*num, num)
		iMin = Min(iMin*num, num)

		max = Max(max, iMax)
	}
	return max
}

