/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
//func candy(ratings []int) int {
//	var max = func(a, b int) int {
//		if a > b {
//			return a
//		}
//		return b
//	}
//	var getCandy = func(val int) int {
//		return val * (val + 1) / 2
//	}
//
//	var length = len(ratings)
//	if length <= 0 {
//		return length
//	}
//	var candies, up, down, oldSlope int
//	for i := 1; i < length; i++ {
//		var newSlope int
//		if ratings[i] > ratings[i-1] {
//			newSlope = 1
//		} else if ratings[i] < ratings[i-1] {
//			newSlope = -1
//		} else {
//			newSlope = 0
//		}
//
//		if (oldSlope > 0 && newSlope == 0) || (oldSlope < 0 && newSlope >= 0) {
//			candies += getCandy(up) + getCandy(down) + max(up, down)
//			up, down = 0, 0
//		}
//
//		if newSlope > 0 {
//			up++
//		} else if newSlope < 0 {
//			down++
//		} else {
//			candies++
//		}
//		oldSlope = newSlope
//	}
//	candies += getCandy(up) + getCandy(down) + max(up, down) + 1
//	return candies
//}

func candy(ratings []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var length = len(ratings)
	left, right := make([]int, length), make([]int, length)
	for i := 0; i < length; i++ {
		left[i], right[i] = 1, 1
	}
	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	var cnt = left[length-1]
	for i := length - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
		cnt += max(left[i], right[i])
	}
	return cnt
}

// @lc code=end

