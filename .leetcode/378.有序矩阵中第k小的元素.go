/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第K小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/description/
 *
 * algorithms
 * Medium (55.61%)
 * Likes:    140
 * Dislikes: 0
 * Total Accepted:    13.1K
 * Total Submissions: 22.6K
 * Testcase Example:  '[[1,5,9],[10,11,13],[12,13,15]]\n8'
 *
 * 给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第k小的元素。
 * 请注意，它是排序后的第k小元素，而不是第k个元素。
 *
 * 示例:
 *
 *
 * matrix = [
 * ⁠  [ 1,  5,  9],
 * ⁠  [10, 11, 13],
 * ⁠  [12, 13, 15]
 * ],
 * k = 8,
 *
 * 返回 13。
 *
 *
 * 说明:
 * 你可以假设 k 的值永远是有效的, 1 ≤ k ≤ n^2 。
 *
 */

// @lc code=start
//找到矩阵中<=target的数量
func search_less_equal(matrix [][]int, target int) int {
	var length = len(matrix)
	i, j, cnt := length-1, 0, 0
	for i >= 0 && j < length {
		if matrix[i][j] <= target {
			j++
			cnt += i + 1
		} else {
			i--
		}
	}
	return cnt
}

func kthSmallest(matrix [][]int, k int) int {
	var length = len(matrix)
	left, rigth := matrix[0][0], matrix[length-1][length-1]
	for left < rigth {
		var mid = left + (rigth-left)/2
		cnt := search_less_equal(matrix, mid)
		if cnt < k {
			left = mid + 1
		} else {
			rigth = mid
		}
	}
	return left
}

// @lc code=end

