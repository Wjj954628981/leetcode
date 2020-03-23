/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 *
 * https://leetcode-cn.com/problems/power-of-three/description/
 *
 * algorithms
 * Easy (46.44%)
 * Likes:    90
 * Dislikes: 0
 * Total Accepted:    35.7K
 * Total Submissions: 76.7K
 * Testcase Example:  '27'
 *
 * 给定一个整数，写一个函数来判断它是否是 3 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 27
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: 0
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: 9
 * 输出: true
 *
 * 示例 4:
 *
 * 输入: 45
 * 输出: false
 *
 * 进阶：
 * 你能不使用循环或者递归来完成本题吗？
 *
 */

import (
	"strconv"
	"strings"
)

// @lc code=start
func isPowerOfThree(n int) bool {
	//最后一定以1结尾，效率提升max
	if n&1 != 1 {
		return false
	}
	//base 10 : 10, 100, 1000 ...
	//base 2 : 10(2), 100(4), 1000(8) ...
	//base 3 : 10(3), 100(9), 1000(27) ...
	s := strconv.FormatInt(int64(n), 3)
	//正则匹配效率辣鸡
	//matched, _ := regexp.MatchString(`^10*$`, s)
	return s[0] == '1' && strings.Count(s, "0") == len(s)-1
}

// @lc code=end

