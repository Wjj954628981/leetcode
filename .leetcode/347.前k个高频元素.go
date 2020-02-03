/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前K个高频元素
 *
 * https://leetcode-cn.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (56.90%)
 * Likes:    234
 * Dislikes: 0
 * Total Accepted:    33.9K
 * Total Submissions: 56.6K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
 *
 * 示例 1:
 *
 * 输入: nums = [1,1,1,2,2,3], k = 2
 * 输出: [1,2]
 *
 *
 * 示例 2:
 *
 * 输入: nums = [1], k = 1
 * 输出: [1]
 *
 * 说明：
 *
 *
 * 你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
 * 你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
 *
 *
 */import "sort"

// @lc code=start
type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

//将map转为struct后继承sort接口实现排序
func sortMapByValue(m map[int]int) PairList {
	p := make(PairList, 0)
	for k, v := range m {
		p = append(p, Pair{k, v})
	}
	sort.Sort(p)
	return p
}

func topKFrequent(nums []int, k int) []int {
	var freMap = make(map[int]int)
	for _, v := range nums {
		if _, ok := freMap[v]; ok {
			freMap[v]++
		} else {
			freMap[v] = 1
		}
	}
	sortedMap := sortMapByValue(freMap)
	var ret = make([]int, 0)
	for i := 0; i < k; i++ {
		ret = append(ret, sortedMap[i].Key)
	}
	return ret
}

// @lc code=end

