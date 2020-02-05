/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (37.39%)
 * Likes:    386
 * Dislikes: 0
 * Total Accepted:    59.3K
 * Total Submissions: 148K
 * Testcase Example:  '[1,2]'
 *
 * 请判断一个链表是否为回文链表。
 *
 * 示例 1:
 *
 * 输入: 1->2
 * 输出: false
 *
 * 示例 2:
 *
 * 输入: 1->2->2->1
 * 输出: true
 *
 *
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//反转
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		var tmp = cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

//后半部分链表
func endOfFirstHalf(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//比较前后半部分链表是否一致
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	var middle = endOfFirstHalf(head)
	var middleReverse = reverseList(middle)
	for middleReverse.Next != nil {
		if middleReverse.Val == head.Val {
			middleReverse = middleReverse.Next
			head = head.Next
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

