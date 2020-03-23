/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode-cn.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (41.60%)
 * Likes:    122
 * Dislikes: 0
 * Total Accepted:    12.3K
 * Total Submissions: 29K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 *
 * 返回滑动窗口中的最大值。
 *
 *
 *
 * 示例:
 *
 * 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
 * 输出: [3,3,5,5,6,7]
 * 解释:
 *
 * ⁠ 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 *
 * 提示：
 *
 * 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
 *
 *
 *
 * 进阶：
 *
 * 你能在线性时间复杂度内解决此题吗？
 *
 */
//双端队列
type deque []int

//push时需要将前面小于v的值pop出去
func (d deque) push(v int) deque {
	var tmp = d
	for len(tmp) > 0 && tmp[len(tmp)-1] < v {
		tmp = tmp[:len(tmp)-1]
	}
	return append(tmp, v)
}
func (d deque) max() int {
	return d[0]
}

//pop时需要判断该元素是否存在（有可能在push中被pop出去）
func (d deque) pop(v int) deque {
	if len(d) > 0 && d[0] == v {
		return d[1:]
	}
	return d
}

func maxSlidingWindow(nums []int, k int) []int {
	var dq = make(deque, 0)

	var ret = make([]int, 0)

	for idx, v := range nums {
		dq = dq.push(v)
		//滑动窗口
		if idx >= k-1 {
			ret = append(ret, dq.max())
			dq = dq.pop(nums[idx-k+1])
		}
	}
	return ret
}

