/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = new(ListNode)
	var headPoint = head
	for l1 != nil && l2 != nil {
		var tmp *ListNode
		if l1.Val < l2.Val {
			tmp = l1
			l1 = l1.Next
			tmp.Next = nil
		} else {
			tmp = l2
			l2 = l2.Next
			tmp.Next = nil
		}
		head.Next = tmp
		head = tmp
	}
	if l1 != nil {
		head.Next = l1
	} else {
		head.Next = l2
	}
	return headPoint.Next
}

// @lc code=end

