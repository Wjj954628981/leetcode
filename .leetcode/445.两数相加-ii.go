/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 *
 * https://leetcode-cn.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (49.21%)
 * Likes:    75
 * Dislikes: 0
 * Total Accepted:    6.1K
 * Total Submissions: 12.1K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * 给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
 *
 *
 *
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 * 进阶:
 *
 * 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
 *
 * 示例:
 *
 *
 * 输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出: 7 -> 8 -> 0 -> 7
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Stack struct {
	nodes []int
	count int
}

func (s *Stack) Push(n int) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() int {
	if s.count == 0 {
		return -1
	}
	s.count--
	return s.nodes[s.count]
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	s1, s2 := Stack{nodes: []int{}, count: 0}, Stack{nodes: []int{}, count: 0}
	for l1 != nil {
		s1.Push(l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2.Push(l2.Val)
		l2 = l2.Next
	}
	var isCarry = 0
	var result *ListNode
	for s1.count > 0 || s2.count > 0 {
		var sum = isCarry
		if s1.count > 0 {
			sum += s1.Pop()
		}
		if s2.count > 0 {
			sum += s2.Pop()
		}
		if sum >= 10 {
			isCarry = 1
			sum -= 10
		} else {
			isCarry = 0
		}
		result = &ListNode{Val: sum, Next: result}
	}
	if isCarry == 1 {
		result = &ListNode{Val: 1, Next: result}
	}
	return result
}

