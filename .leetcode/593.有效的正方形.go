import "sort"

/*
 * @lc app=leetcode.cn id=593 lang=golang
 *
 * [593] 有效的正方形
 */
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	var disp = func(a []int, b []int) int {
		return (b[0]-a[0])*(b[0]-a[0]) + (b[1]-a[1])*(b[1]-a[1])
	}
	var dispList = make([]int, 0)
	dispList = append(dispList, disp(p1, p2), disp(p1, p3), disp(p1, p4), disp(p2, p3), disp(p2, p4), disp(p3, p4))
	sort.Ints(dispList)
	return dispList[0] == dispList[3] && dispList[0] != 0
}

