/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */
func permuteUnique(nums []int) [][]int {
	var ret [][]int
	len := len(nums)
	var backtrace func(nums []int, first int)
	var isDupl func(nums []int, start int, end int) bool

	isDupl = func(nums []int, start int, end int) bool {
		for i := start; i < end; i++ {
			if nums[i] == nums[end] {
				return false
			}
		}
		return true
	}

	backtrace = func(nums []int, first int) {
		if first == len {
			var tmp = make([]int, len)
			copy(tmp, nums)
			ret = append(ret, tmp)
		}
		for i := first; i < len; i++ {
			if !isDupl(nums, first, i) {
				continue
			}
			nums[first], nums[i] = nums[i], nums[first]
			backtrace(nums, first+1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrace(nums, 0)
	return ret
}

