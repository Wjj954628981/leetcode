/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 *
 * https://leetcode-cn.com/problems/shuffle-an-array/description/
 *
 * algorithms
 * Medium (45.83%)
 * Likes:    29
 * Dislikes: 0
 * Total Accepted:    9.3K
 * Total Submissions: 19.9K
 * Testcase Example:  '["Solution","shuffle","reset","shuffle"]\n[[[1,2,3]],[],[],[]]'
 *
 * 打乱一个没有重复元素的数组。
 *
 * 示例:
 *
 *
 * // 以数字集合 1, 2 和 3 初始化数组。
 * int[] nums = {1,2,3};
 * Solution solution = new Solution(nums);
 *
 * // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
 * solution.shuffle();
 *
 * // 重设数组到它的初始状态[1,2,3]。
 * solution.reset();
 *
 * // 随机返回数组[1,2,3]打乱后的结果。
 * solution.shuffle();
 *
 *
 */
import "math/rand"

type Solution struct {
	original []int
	arr      []int
	length   int
}

func Constructor(nums []int) Solution {
	length := len(nums)
	original := make([]int, length)
	copy(original, nums)
	return Solution{
		original: original,
		arr:      nums,
		length:   length,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	arr := make([]int, this.length)
	copy(arr, this.original)
	this.arr = arr
	return this.original
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := 0; i < this.length; i++ {
		tmp := rand.Intn(this.length)
		this.arr[i], this.arr[tmp] = this.arr[tmp], this.arr[i]
	}
	return this.arr
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

