import "math"

/*
 * @lc app=leetcode.cn id=1131 lang=golang
 *
 * [1131] 绝对值表达式的最大值
 */

// @lc code=start
func maxAbsValExpr(arr1 []int, arr2 []int) int {
	var max1 = math.MinInt16
	var max2 = math.MinInt16
	var max3 = math.MinInt16
	var max4 = math.MinInt16
	var min1 = math.MaxInt16
	var min2 = math.MaxInt16
	var min3 = math.MaxInt16
	var min4 = math.MaxInt16

	var maxMinFun = func(x int, y int, flag bool) int {
		if (flag && x > y) || (!flag && x < y) {
			return x
		}
		return y
	}

	for i := 0; i < len(arr1); i++ {
		var tmp1 = arr1[i] + arr2[i] + i
		var tmp2 = arr1[i] + arr2[i] - i
		var tmp3 = arr1[i] - arr2[i] + i
		var tmp4 = arr1[i] - arr2[i] - i
		max1 = maxMinFun(max1, tmp1, true)
		min1 = maxMinFun(min1, tmp1, false)

		max2 = maxMinFun(max2, tmp2, true)
		min2 = maxMinFun(min2, tmp2, false)

		max3 = maxMinFun(max3, tmp3, true)
		min3 = maxMinFun(min3, tmp3, false)

		max4 = maxMinFun(max4, tmp4, true)
		min4 = maxMinFun(min4, tmp4, false)
	}

	return maxMinFun(maxMinFun(max1-min1, max2-min2, true), maxMinFun(max3-min3, max4-min4, true), true)
}

// @lc code=end

