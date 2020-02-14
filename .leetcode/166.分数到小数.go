/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 *
 * https://leetcode-cn.com/problems/fraction-to-recurring-decimal/description/
 *
 * algorithms
 * Medium (25.61%)
 * Likes:    100
 * Dislikes: 0
 * Total Accepted:    8.1K
 * Total Submissions: 31.5K
 * Testcase Example:  '1\n2'
 *
 * 给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。
 *
 * 如果小数部分为循环小数，则将循环的部分括在括号内。
 *
 * 示例 1:
 *
 * 输入: numerator = 1, denominator = 2
 * 输出: "0.5"
 *
 *
 * 示例 2:
 *
 * 输入: numerator = 2, denominator = 1
 * 输出: "2"
 *
 * 示例 3:
 *
 * 输入: numerator = 2, denominator = 3
 * 输出: "0.(6)"
 *
 *
 */
import (
	"math"
	"strconv"
	"strings"
)

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var byteSlice []byte
	if numerator*denominator < 0 {
		byteSlice = append(byteSlice, '-')
		numerator = int(math.Abs(float64(numerator)))
		denominator = int(math.Abs(float64(denominator)))
	}
	byteSlice = append(byteSlice, []byte(strconv.Itoa(numerator/denominator))...)
	var remainder = numerator % denominator
	if remainder == 0 {
		var sb strings.Builder
		sb.Write(byteSlice)
		return sb.String()
	}
	byteSlice = append(byteSlice, '.')
	var m = make(map[int]int)

	for remainder != 0 {
		if idx, OK := m[remainder]; OK {
			rear := append([]byte{}, byteSlice[idx:]...)
			rear = append(rear, ')')
			byteSlice = append(append(byteSlice[:idx], '('), rear...)
			break
		}
		m[remainder] = len(byteSlice)
		remainder *= 10
		byteSlice = append(byteSlice, []byte(strconv.Itoa(remainder/denominator))...)
		remainder %= denominator
	}
	var sb strings.Builder
	sb.Write(byteSlice)
	return sb.String()
}

// @lc code=end

