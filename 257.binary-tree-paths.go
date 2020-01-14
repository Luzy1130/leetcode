/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 *
 * https://leetcode.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (47.02%)
 * Likes:    1195
 * Dislikes: 83
 * Total Accepted:    267.4K
 * Total Submissions: 552.5K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * Given a binary tree, return all root-to-leaf paths.
 * 
 * Note: A leaf is a node with no children.
 * 
 * Example:
 * 
 * 
 * Input:
 * 
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 * 
 * Output: ["1->2->5", "1->3"]
 * 
 * Explanation: All root-to-leaf paths are: 1->2->5, 1->3
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
 
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
		return res
	}

	var childPaths []string
	if root.Left != nil {
		lPaths := binaryTreePaths(root.Left)
		childPaths = append(childPaths, lPaths...)
	}
	if root.Right != nil {
		rPaths := binaryTreePaths(root.Right)
		childPaths = append(childPaths, rPaths...)
	}

	for _, v := range childPaths {
		tmp := strconv.Itoa(root.Val) + "->" + v
		res = append(res, tmp)
	}
	return res
}
// @lc code=end

