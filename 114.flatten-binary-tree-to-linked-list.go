/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 *
 * https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (46.65%)
 * Likes:    2243
 * Dislikes: 277
 * Total Accepted:    305.7K
 * Total Submissions: 655.1K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * Given a binary tree, flatten it to a linked list in-place.
 * 
 * For example, given the following tree:
 * 
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \   \
 * 3   4   6
 * 
 * 
 * The flattened tree should look like:
 * 
 * 
 * 1
 * ⁠\
 * ⁠ 2
 * ⁠  \
 * ⁠   3
 * ⁠    \
 * ⁠     4
 * ⁠      \
 * ⁠       5
 * ⁠        \
 * ⁠         6
 * 
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

// 将左子树全部移到右子树
func flatten(root *TreeNode)  {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	if root.Left != nil {
		pos := root.Left
		
		for pos.Right != nil {
			pos = pos.Right
		}

		pos.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	
	return
}
// @lc code=end

