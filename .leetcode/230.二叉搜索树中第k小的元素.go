/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (64.59%)
 * Likes:    142
 * Dislikes: 0
 * Total Accepted:    33.8K
 * Total Submissions: 49.6K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
 *
 * 说明：
 * 你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
 *
 * 示例 1:
 *
 * 输入: root = [3,1,4,null,2], k = 1
 * ⁠  3
 * ⁠ / \
 * ⁠1   4
 * ⁠ \
 * 2
 * 输出: 1
 *
 * 示例 2:
 *
 * 输入: root = [5,3,6,2,4,null,null,1], k = 3
 * ⁠      5
 * ⁠     / \
 * ⁠    3   6
 * ⁠   / \
 * ⁠  2   4
 * ⁠ /
 * ⁠1
 * 输出: 3
 *
 * 进阶：
 * 如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
 *
 */
import (
	"container/list"
	"reflect"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreeBySlice(s []interface{}) *TreeNode {
	var tSlice = make([]*TreeNode, 0)
	for _, v := range s {
		var tmp *TreeNode
		if v == nil {
			tmp = nil
		} else {
			tmp = &TreeNode{Val: v.(int)}
		}
		tSlice = append(tSlice, tmp)
	}
	for i := 0; i <= len(s)/2-1; i++ {
		tSlice[i].Left = tSlice[i*2+1]
		if i*2+2 < len(s) {
			tSlice[i].Right = tSlice[i*2+2]
		}
	}
	return tSlice[0]
}

func Tree2Slice(root *TreeNode) []interface{} {
	q := list.New()
	q.PushBack(root)

	var ret []interface{}
	for q.Len() > 0 {
		tmp := q.Front()
		q.Remove(tmp)
		//注意类型转换
		tmpVal := tmp.Value.(*TreeNode)

		if tmpVal == nil {
			ret = append(ret, tmpVal)
		} else {
			ret = append(ret, tmpVal.Val)
			//if tmpVal.Left != nil || tmpVal.Right != nil {
			q.PushBack(tmpVal.Left)
			q.PushBack(tmpVal.Right)
			//}
		}
	}
	var maxLen int
	for idx, v := range ret {
		if reflect.TypeOf(v).Kind() == reflect.Int {
			maxLen = idx + 1
		}
	}
	return ret[:maxLen]
}

func kthSmallest(root *TreeNode, k int) int {
	stack := list.New()
	for {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		var val = stack.Back()
		stack.Remove(val)

		k--
		root = val.Value.(*TreeNode)
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

// @lc code=end

