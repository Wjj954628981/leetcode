/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 *
 * https://leetcode-cn.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (49.10%)
 * Likes:    222
 * Dislikes: 0
 * Total Accepted:    22.3K
 * Total Submissions: 45.4K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 现在你总共有 n 门课需要选，记为 0 到 n-1。
 *
 * 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
 *
 * 给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？
 *
 * 示例 1:
 *
 * 输入: 2, [[1,0]]
 * 输出: true
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
 *
 * 示例 2:
 *
 * 输入: 2, [[1,0],[0,1]]
 * 输出: false
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
 *
 * 说明:
 *
 *
 * 输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
 * 你可以假定输入的先决条件中没有重复的边。
 *
 *
 * 提示:
 *
 *
 * 这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
 * 通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
 *
 * 拓扑排序也可以通过 BFS 完成。
 *
 *
 *
 */
import "container/list"

// @lc code=start
//递归的辣鸡效率可惜可惜
//func dfs(start int, visited []int, m [][]int) bool {
//	if visited[start] == 1 {
//		return false
//	}
//	if visited[start] == -1 {
//		return true
//	}
//	visited[start] = 1
//	for _, item := range m[start] {
//		if !dfs(item, visited, m) {
//			return false
//		}
//	}
//	visited[start] = -1
//	return true
//}
//
//func canFinish(numCourses int, prerequisites [][]int) bool {
//	//学完k后才能学v的课程
//	var path = make([][]int, numCourses)
//	for i := 0; i < len(prerequisites); i++ {
//		path[prerequisites[i][1]] = append(path[prerequisites[i][1]], prerequisites[i][0])
//	}
//	for i := 0; i < numCourses; i++ {
//		//visited的val为1代表访问过，-1代表访问过且return true
//		var visited = make([]int, numCourses)
//		if !dfs(i, visited, path) {
//			return false
//		}
//	}
//	return true
//}

func canFinish(numCourses int, prerequisites [][]int) bool {
	//学完k后才能学v的课程,学k需要先学习的课程数量
	path, inDegree := make([][]int, numCourses), make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		inDegree[prerequisites[i][0]]++
		path[prerequisites[i][1]] = append(path[prerequisites[i][1]], prerequisites[i][0])
	}

	//学习课程列表，l.PushBack为学习了一门课程
	var l = list.New()
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			l.PushBack(i)
		}
	}

	var visited = make(map[int]bool, 0)
	//循环学习过程
	for l.Len() != 0 {
		v := l.Remove(l.Front()).(int)
		visited[v] = true
		for _, item := range path[v] {
			inDegree[item]--
			if inDegree[item] == 0 {
				l.PushBack(item)
			}
		}
	}
	return len(visited) == numCourses
}

// @lc code=end

