/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 *
 * https://leetcode-cn.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (60.86%)
 * Likes:    390
 * Dislikes: 0
 * Total Accepted:    38.3K
 * Total Submissions: 60.4K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 *
 * 示例 1:
 *
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 *
 *
 * 示例 2:
 *
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge(left *ListNode, right *ListNode) *ListNode {
	var root = &ListNode{}
	var tmp = root
	for left != nil && right != nil {
		if left.Val < right.Val {
			root.Next = left
			left = left.Next
		} else {
			root.Next = right
			right = right.Next
		}
		root = root.Next
	}
	if left != nil {
		root.Next = left
	} else {
		root.Next = right
	}
	return tmp.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp := slow.Next
	slow.Next = nil

	return merge(sortList(head), sortList(tmp))
}

// @lc code=end

