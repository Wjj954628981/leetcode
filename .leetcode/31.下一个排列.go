/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
func nextPermutation(nums []int) {
	len := len(nums)
	i := len - 2
	j := len - 1
	for {
		if i >= 0 && nums[i] >= nums[i+1] {
			i--
		} else {
			break
		}
	}
	if i >= 0 {
		for {
			if j >= i && nums[j] <= nums[i] {
				j--
			} else {
				break
			}
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, len-1)
}

func swap(nums []int, idx1 int, idx2 int) {
	tmp := nums[idx1]
	nums[idx1] = nums[idx2]
	nums[idx2] = tmp
}

func reverse(nums []int, start int, end int) {
	for {
		if start < end {
			swap(nums, start, end)
			start++
			end--
		} else {
			break
		}
	}
}

