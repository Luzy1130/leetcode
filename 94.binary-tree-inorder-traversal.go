/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (58.13%)
 * Likes:    2532
 * Dislikes: 106
 * Total Accepted:    638.6K
 * Total Submissions: 1M
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the inorder traversal of its nodes' values.
 * 
 * Example:
 * 
 * 
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * Output: [1,3,2]
 * 
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	pos := root
	if root == nil {
		return res
	}
	
	for pos != nil || len(stack) != 0 {
		for pos != nil {
			stack = append(stack, pos)
			pos = pos.Left
		}

		if len(stack) != 0 {
			node := stack[len(stack)-1]
			stack = stack[0:(len(stack)-1)]
			res = append(res, node.Val)
			pos = node.Right
		}
	}
	return res
}
// @lc code=end

