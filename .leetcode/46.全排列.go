/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
func permute(nums []int) [][]int {
	var ret [][]int
	len := len(nums)
	var backtrace func(nums []int, first int)
	backtrace = func(nums []int, first int) {
		if first == len {
			var tmp = make([]int, len)
			copy(tmp, nums)
			ret = append(ret, tmp)
		}
		for i := first; i < len; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrace(nums, first+1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrace(nums, 0)
	return ret
}

