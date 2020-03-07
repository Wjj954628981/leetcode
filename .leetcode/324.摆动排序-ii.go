/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 *
 * https://leetcode-cn.com/problems/wiggle-sort-ii/description/
 *
 * algorithms
 * Medium (31.94%)
 * Likes:    90
 * Dislikes: 0
 * Total Accepted:    7.6K
 * Total Submissions: 21.8K
 * Testcase Example:  '[1,5,1,1,6,4]'
 *
 * 给定一个无序的数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
 *
 * 示例 1:
 *
 * 输入: nums = [1, 5, 1, 1, 6, 4]
 * 输出: 一个可能的答案是 [1, 4, 1, 5, 1, 6]
 *
 * 示例 2:
 *
 * 输入: nums = [1, 3, 2, 2, 3, 1]
 * 输出: 一个可能的答案是 [2, 3, 1, 3, 1, 2]
 *
 * 说明:
 * 你可以假设所有输入都会得到有效的结果。
 *
 * 进阶:
 * 你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？
 *
 */
import "sort"

// @lc code=start
func quickSelect(nums []int, begin, end, find int) {
	head, tail, i, keyVal := begin, end, begin+1, nums[begin]
	for i <= tail {
		if nums[i] > keyVal {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else if nums[i] < keyVal {
			nums[i], nums[head] = nums[head], nums[i]
			i++
			head++
		} else {
			i++
		}
	}
	if head > find {
		quickSelect(nums, begin, head-1, find)
	} else if head < find {
		quickSelect(nums, head+1, end, find)
	}
}

func wiggleSort(nums []int) {
	var length = len(nums)

	var numsCp = make([]int, length)
	copy(numsCp, nums)

	//quickSelect(numsCp, 0, length-1, length/2)
	sort.Ints(numsCp)

	middle, last := (length+1)>>1, length
	for i := 0; i < length; i++ {
		if i&1 == 0 {
			middle--
			nums[i] = numsCp[middle]
		} else {
			last--
			nums[i] = numsCp[last]
		}
	}
}

// @lc code=end

