/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 *
 * https://leetcode.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (38.53%)
 * Likes:    1306
 * Dislikes: 402
 * Total Accepted:    385.7K
 * Total Submissions: 978.6K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * Given a binary tree and a sum, determine if the tree has a root-to-leaf path
 * such that adding up all the values along the path equals the given sum.
 * 
 * Note: A leaf is a node with no children.
 * 
 * Example:
 * 
 * Given the below binary tree and sum = 22,
 * 
 * 
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \      \
 * 7    2      1
 * 
 * 
 * return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
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
func hasPathSumImpl(root *TreeNode, sum int) bool {
	val := root.Val
	if root.Left == nil && root.Right == nil {
		if val == sum {
			return true
		} else {
			return false
		}
	}

	var res bool = false
	if root.Left != nil {
		res = hasPathSumImpl(root.Left, sum - val)
	}
	if res == false && root.Right != nil {
		res = hasPathSumImpl(root.Right, sum - val)
	}
	return res
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
    return hasPathSumImpl(root, sum)
}
// @lc code=end

