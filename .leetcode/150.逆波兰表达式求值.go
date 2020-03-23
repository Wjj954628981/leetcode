/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 *
 * https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/description/
 *
 * algorithms
 * Medium (44.66%)
 * Likes:    95
 * Dislikes: 0
 * Total Accepted:    24.9K
 * Total Submissions: 51.1K
 * Testcase Example:  '["2","1","+","3","*"]'
 *
 * 根据逆波兰表示法，求表达式的值。
 *
 * 有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
 *
 * 说明：
 *
 *
 * 整数除法只保留整数部分。
 * 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
 *
 *
 * 示例 1：
 *
 * 输入: ["2", "1", "+", "3", "*"]
 * 输出: 9
 * 解释: ((2 + 1) * 3) = 9
 *
 *
 * 示例 2：
 *
 * 输入: ["4", "13", "5", "/", "+"]
 * 输出: 6
 * 解释: (4 + (13 / 5)) = 6
 *
 *
 * 示例 3：
 *
 * 输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
 * 输出: 22
 * 解释:
 * ⁠ ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5
 * = 17 + 5
 * = 22
 *
 */import "strconv"

// @lc code=start
type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}
func (s *Stack) Pop() int {
	var arr = *s
	var ret = arr[len(arr)-1]
	*s = arr[:len(arr)-1]
	return ret
}

func evalRPN(tokens []string) int {
	var s = new(Stack)
	for _, v := range tokens {
		val, err := strconv.Atoi(v)
		if err != nil {
			b, a := s.Pop(), s.Pop()
			switch v {
			case "+":
				s.Push(a + b)
			case "-":
				s.Push(a - b)
			case "*":
				s.Push(a * b)
			case "/":
				s.Push(a / b)
			}
		} else {
			s.Push(val)
		}
	}
	return s.Pop()
}

// @lc code=end

