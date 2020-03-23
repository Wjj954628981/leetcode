/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 *
 * https://leetcode-cn.com/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (35.61%)
 * Likes:    109
 * Dislikes: 0
 * Total Accepted:    9.4K
 * Total Submissions: 22.5K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n[[],[1],[2],[],[3],[]]'
 *
 * 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
 *
 * 例如，
 *
 * [2,3,4] 的中位数是 3
 *
 * [2,3] 的中位数是 (2 + 3) / 2 = 2.5
 *
 * 设计一个支持以下两种操作的数据结构：
 *
 *
 * void addNum(int num) - 从数据流中添加一个整数到数据结构中。
 * double findMedian() - 返回目前所有元素的中位数。
 *
 *
 * 示例：
 *
 * addNum(1)
 * addNum(2)
 * findMedian() -> 1.5
 * addNum(3)
 * findMedian() -> 2
 *
 * 进阶:
 *
 *
 * 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
 * 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
 *
 *
 */

// @lc code=start
import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	lo *maxHeap
	hi *minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := MedianFinder{
		lo: new(maxHeap),
		hi: new(minHeap),
	}
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.lo, num)

	if loMax := heap.Pop(this.lo); loMax != nil {
		heap.Push(this.hi, loMax)
	}

	if this.lo.Len() < this.hi.Len() {
		if hiMin := heap.Pop(this.hi); hiMin != nil {
			heap.Push(this.lo, hiMin)
		}
	}
	fmt.Println(this.lo, this.hi)
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return float64(this.lo.Top().(int))
	} else {
		return (float64(this.lo.Top().(int)) + float64(this.hi.Top().(int))) * 0.5
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return (h)[i] < (h)[j]
}

func (h minHeap) Swap(i, j int) {
	(h)[i], (h)[j] = (h)[j], (h)[i]
}

func (h minHeap) Top() interface{} {
	if h.Len() == 0 {
		return 0
	}
	return h[0]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return (h)[i] > (h)[j]
}

func (h maxHeap) Swap(i, j int) {
	(h)[i], (h)[j] = (h)[j], (h)[i]
}

func (h maxHeap) Top() interface{} {
	if h.Len() == 0 {
		return 0
	}
	return h[0]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

