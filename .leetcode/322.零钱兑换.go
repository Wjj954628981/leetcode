/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (33.03%)
 * Likes:    352
 * Dislikes: 0
 * Total Accepted:    39.7K
 * Total Submissions: 106.2K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 *
 * 示例 1:
 *
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3
 * 解释: 11 = 5 + 5 + 1
 *
 * 示例 2:
 *
 * 输入: coins = [2], amount = 3
 * 输出: -1
 *
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 *
 */

// @lc code=start
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func coinChange(coins []int, amount int) int {
//	var visited = make(map[int]int)
//	sort.Ints(coins)
//
//	visited[0] = 0
//	return dp(coins, amount, visited)
//}
//
//func dp(coins []int, amount int, visited map[int]int) int {
//	if v, ok := visited[amount]; ok {
//		return v
//	}
//	var ret = math.MaxInt8
//	for _, v := range coins {
//		if v == amount {
//			visited[v] = 1
//			return 1
//		} else if v < amount {
//			if last := dp(coins, amount-v, visited); last != -1 {
//				ret = Min(ret, last+1)
//			}
//		} else {
//			break
//		}
//	}
//	if ret == math.MaxInt8 {
//		ret = -1
//	}
//	visited[amount] = ret
//	return ret
//}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+2)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = Min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

// @lc code=end

