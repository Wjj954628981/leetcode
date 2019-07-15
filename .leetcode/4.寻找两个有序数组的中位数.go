/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */
 func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	min := func(x int, y int) int {
		if x < y {
			return x
		}
		return y
	}

	var findKth func(nums1 []int, nums2 []int, k int) float64
	findKth = func(nums1 []int, nums2 []int, k int) float64 {
		if len(nums1) == 0 {
			return float64(nums2[k-1])
		}
		if len(nums2) == 0 {
			return float64(nums1[k-1])
		}
		if k == 1 {
			return math.Min(float64(nums1[0]), float64(nums2[0]))
		}
		i := min(k>>1, len(nums1))
		j := min(k>>1, len(nums2))
		if nums1[i-1] > nums2[j-1] {
			return findKth(nums1, nums2[j:], k-j)
		}
		return findKth(nums1[i:], nums2, k-i)
	}

	len1 := len(nums1)
	len2 := len(nums2)
	return (findKth(nums1, nums2, (len1+len2+1)>>1) + findKth(nums1, nums2, (len1+len2+2)>>1)) / 2
}

