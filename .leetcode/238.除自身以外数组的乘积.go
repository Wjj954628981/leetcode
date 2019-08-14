/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */
func productExceptSelf(nums []int) []int {
	ltr := 1
	rtl := 1
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i += 1 {
		ret[i] = 1
	}
	for i := 0; i < len(nums); i += 1 {
		ret[i] *= ltr
		ret[len(nums)-1-i] *= rtl
		ltr *= nums[i]
		rtl *= nums[len(nums)-1-i]
	}
	return ret
}

