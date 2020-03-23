/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 *
 * https://leetcode-cn.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (59.61%)
 * Likes:    402
 * Dislikes: 0
 * Total Accepted:    33.8K
 * Total Submissions: 53.9K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和
 * n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
 *
 * 示例 1:
 *
 * 输入: [1,3,4,2,2]
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: [3,1,3,4,2]
 * 输出: 3
 *
 *
 * 说明：
 *
 *
 * 不能更改原数组（假设数组是只读的）。
 * 只能使用额外的 O(1) 的空间。
 * 时间复杂度小于 O(n^2) 。
 * 数组中只有一个重复的数字，但它可能不止重复出现一次。
 *
 *
 */

// @lc code=start
func findDuplicate(nums []int) int {
	var m = make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = true
		}
	}
	return 0
}

//快慢指针
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	slow, fast = nums[slow], nums[nums[fast]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	before, after := 0, slow
	for before != after {
		before, after = nums[before], nums[after]
	}
	return before
}

// @lc code=end

