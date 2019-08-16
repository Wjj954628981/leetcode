/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (54.83%)
 * Likes:    374
 * Dislikes: 0
 * Total Accepted:    60.5K
 * Total Submissions: 108K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 示例:
 *
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 * 说明:
 *
 *
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 *
 *
 */
func moveZeroes(nums []int) {
	for cur, tmp := 0, 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			if tmp != cur {
				nums[tmp], nums[cur] = nums[cur], nums[tmp]
			}
			tmp++
		}
	}
}

