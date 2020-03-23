/*
 * @lc app=leetcode.cn id=630 lang=golang
 *
 * [630] 课程表 III
 *
 * https://leetcode-cn.com/problems/course-schedule-iii/description/
 *
 * algorithms
 * Hard (28.89%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    1.4K
 * Total Submissions: 4.8K
 * Testcase Example:  '[[100,200],[200,1300],[1000,1250],[2000,3200]]'
 *
 * 这里有 n 门不同的在线课程，他们按从 1 到 n 编号。每一门课程有一定的持续上课时间（课程时间）t 以及关闭时间第 d 天。一门课要持续学习 t
 * 天直到第 d 天时要完成，你将会从第 1 天开始。
 *
 * 给出 n 个在线课程用 (t, d) 对表示。你的任务是找出最多可以修几门课。
 *
 *
 *
 * 示例：
 *
 *
 * 输入: [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
 * 输出: 3
 * 解释:
 * 这里一共有 4 门课程, 但是你最多可以修 3 门:
 * 首先, 修第一门课时, 它要耗费 100 天，你会在第 100 天完成, 在第 101 天准备下门课。
 * 第二, 修第三门课时, 它会耗费 1000 天，所以你将在第 1100 天的时候完成它, 以及在第 1101 天开始准备下门课程。
 * 第三, 修第二门课时, 它会耗时 200 天，所以你将会在第 1300 天时完成它。
 * 第四门课现在不能修，因为你将会在第 3300 天完成它，这已经超出了关闭日期。
 *
 *
 *
 * 提示:
 *
 *
 * 整数 1 <= d, t, n <= 10,000 。
 * 你不能同时修两门课程。
 *
 *
 *
 *
 */
import (
	"container/heap"
	"sort"
)

// @lc code=start
type sortBy2nd [][]int

func (s sortBy2nd) Len() int {
	return len(s)
}
func (s sortBy2nd) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}
func (s sortBy2nd) Less(a, b int) bool {
	return s[a][1] < s[b][1]
}

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}
func (m maxHeap) Swap(a, b int) {
	m[a], m[b] = m[b], m[a]
}
func (m maxHeap) Less(a, b int) bool {
	return m[a] > m[b]
}
func (m maxHeap) Top() int {
	if m.Len() == 0 {
		return 0
	}
	return m[0]
}
func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *maxHeap) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

func scheduleCourse(courses [][]int) int {
	var s sortBy2nd = courses
	sort.Sort(s)
	nowLen, m := 0, new(maxHeap)
	for _, val := range s {
		if nowLen+val[0] <= val[1] {
			nowLen += val[0]
			heap.Push(m, val[0])
		} else if m.Len() > 0 && m.Top() > val[0] {
			nowLen += val[0] - heap.Pop(m).(int)
			heap.Push(m, val[0])
		}
	}
	return m.Len()
}

// @lc code=end

