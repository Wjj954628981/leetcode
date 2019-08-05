/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	length := len(inorder)
	inMap := make(map[int]int, length)
	for k, v := range inorder {
		inMap[v] = k
	}
	preIndex := 0

	var helper func(s int, e int) *TreeNode
	helper = func(s int, e int) *TreeNode {
		if s == e || preIndex >= length {
			return nil
		}
		index := inMap[preorder[preIndex]]
		var root TreeNode
		root.Val = preorder[preIndex]
		preIndex++

		root.Left = helper(s, index)
		root.Right = helper(index+1, e)
		return &root
	}
	return helper(0, length)
}

