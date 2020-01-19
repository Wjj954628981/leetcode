package main

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
func makeList(input []int, pos int) *ListNode {
	if len(input) == 0 {
		return nil
	}
	var head = &ListNode{
		Val:  input[0],
		Next: nil,
	}
	var tmp = head
	var posNode *ListNode
	for k, v := range input {
		if k > 0 {
			var node = &ListNode{
				Val:  v,
				Next: nil,
			}
			tmp.Next = node
			tmp = node
		}
		if pos == k {
			posNode = tmp
		}
	}
	if posNode != nil {
		tmp.Next = posNode
	}
	return head
}

func (head *ListNode) toString() {
	var meet = detectCycle(head)
	var slice []int
	var flag bool
	var pos, count int
	for head != nil {
		if meet != nil && head == meet {
			if !flag {
				pos = count
				flag = true
			} else {
				break
			}
		}
		slice = append(slice, head.Val)
		head = head.Next
		count++
	}
	fmt.Println(slice, pos)
}*/

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	var meet *ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			meet = slow
			break
		}
	}
	if meet != nil {
		slow = head
		for slow != meet {
			slow = slow.Next
			meet = meet.Next
		}
	}
	return meet
}

// @lc code=end
