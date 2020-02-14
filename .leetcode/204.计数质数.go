/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (32.12%)
 * Likes:    268
 * Dislikes: 0
 * Total Accepted:    41.9K
 * Total Submissions: 130.3K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 *
 * 示例:
 *
 * 输入: 10
 * 输出: 4
 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 */

// @lc code=start
//厄拉多塞筛法
func countPrimes(n int) int {
	var isPrim = make([]bool, n)
	var cnt int

	for i := 2; i*i < n; i++ {
		if !isPrim[i] {
			for j := i * i; j < n; j += i {
				isPrim[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !isPrim[i] {
			cnt++
		}
	}
	return cnt
}

//bitmap存储
//func countPrimes(n int) int {
//	var isPrim = make([]int32, n/32+1)
//	var cnt int
//
//	for i := 2; i*i < n; i++ {
//		if isPrim[i/32]&(1<<(uint32(i)&31)) == 0 {
//			for j := i * i; j < n; j += i {
//				isPrim[j/32] |= 1 << (uint32(j) & 31)
//			}
//		}
//	}
//	for i := 2; i < n; i++ {
//		if isPrim[i/32]&(1<<(uint32(i)&31)) == 0 {
//			cnt++
//		}
//	}
//	return cnt
//}
// @lc code=end

