/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 *
 * https://leetcode.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (41.60%)
 * Likes:    1641
 * Dislikes: 142
 * Total Accepted:    382.3K
 * Total Submissions: 905.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, determine if it is height-balanced.
 * 
 * For this problem, a height-balanced binary tree is defined as:
 * 
 * 
 * a binary tree in which the left and right subtrees of every node differ in
 * height by no more than 1.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * Given the following tree [3,9,20,null,null,15,7]:
 * 
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * Return true.
 * 
 * Example 2:
 * 
 * Given the following tree [1,2,2,3,3,null,null,4,4]:
 * 
 * 
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 * 
 * 
 * Return false.
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

func treeHeight(root *TreeNode) (float64, bool) {
	if root == nil {
		return 0, true
	}
	var height float64 = 1
	if root.Left == nil && root.Right == nil {
		return height, true
	}

	heightRight, balancedRight := treeHeight(root.Right)
	if !balancedRight {
		return 0, false
	}
	heightLeft, balancedLeft := treeHeight(root.Left)
	if !balancedLeft {
		return 0, false
	}
	
	diff := math.Abs(heightLeft - heightRight)
	if diff > 1 {
		return 0, false
	}

	height += math.Max(heightLeft, heightRight)

	return height, true
}
func isBalanced(root *TreeNode) bool {
	_, balanced := treeHeight(root)

	return balanced
}
// @lc code=end

