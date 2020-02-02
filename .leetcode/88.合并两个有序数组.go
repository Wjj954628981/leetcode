/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid, i := nums[0], 1
	head, tail := 0, len(nums)-1
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
	nums[head] = mid
	QuickSort(nums[:head])
	QuickSort(nums[head+1:])
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}
	var p1 = m - 1
	var p2 = n - 1
	var p = m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}
	copy(nums1[:p2+1], nums2)
}

