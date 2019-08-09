import "container/heap"

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
func findKthLargest1(nums []int, k int) int {
	length := len(nums)
	mid, i, head, tail := nums[0], 1, 0, length-1
	for head < tail {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	if length-head == k {
		return nums[head]
	} else if length-head < k {
		return findKthLargest1(nums[0:head], k+head-length)
	} else {
		return findKthLargest1(nums[head+1:length], k)
	}
}

func findKthLargest(nums []int, k int) int {
	h := new(intHeap)
	for i := 0; i < k; i++ {
		h.Push(nums[i])
	}
	heap.Init(h)

	for i := k; i < len(nums); i++ {
		if nums[i] > (*h)[0] {
			(*h)[0] = nums[i]
			heap.Fix(h, 0)
		}
	}
	return (*h)[0]
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return (h)[i] < (h)[j]
}

func (h intHeap) Swap(i, j int) {
	(h)[i], (h)[j] = (h)[j], (h)[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (intHeap) Pop() interface{} { return nil }