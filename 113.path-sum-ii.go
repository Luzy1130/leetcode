/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 *
 * https://leetcode.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (44.50%)
 * Likes:    1383
 * Dislikes: 47
 * Total Accepted:    297.1K
 * Total Submissions: 667.5K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * Given a binary tree and a sum, find all root-to-leaf paths where each path's
 * sum equals the given sum.
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
 * ⁠/  \    / \
 * 7    2  5   1
 * 
 * 
 * Return:
 * 
 * 
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
 * ]
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

func helper(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return [][]int{[]int{root.Val}}
		} else {
			return [][]int{}
		}
	}

	res := [][]int{}
	if root.Left != nil {
		leftChilds := helper(root.Left, sum - root.Val)
		if len(leftChilds) != 0 {
			for _, child :=range leftChilds {
				res = append(res, append([]int{root.Val}, child...))
			}
		}
	}

	if root.Right != nil {
		rightChilds := helper(root.Right, sum - root.Val)
		if len(rightChilds) != 0 {
			for _, child :=range rightChilds {
				res = append(res, append([]int{root.Val}, child...))
			}
		}
	}

	return res
}

func pathSum(root *TreeNode, sum int) [][]int {
	return helper(root, sum)
}
// @lc code=end

