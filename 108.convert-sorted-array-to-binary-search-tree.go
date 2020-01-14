/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (52.32%)
 * Likes:    1641
 * Dislikes: 163
 * Total Accepted:    325.7K
 * Total Submissions: 602.6K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given an array where elements are sorted in ascending ordemodels/product_merchant_price.gor, convert it to a
 * height balanced BST.
 * 
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 * 
 * Example:
 * 
 * 
 * Given the sorted array: [-10,-3,0,5,9],
 * 
 * One possible answer is: [0,-3,9,-10,null,5], which represents the following
 * height balanced BST:
 * 
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
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

func insertTree(node *TreeNode, root **Tree, heigt) int {
	
}

func convertSortedArrayToBST(nums []int) *TreeNode {
	root := new(TreeNode)
	root.Val = nums[0]
	leftHeight := 0		// 记录左子树高度
	rightHeight := 0	// 记录右子树高度
	for i := 1; i < len(nums); i++ {
		node := new(TreeNode)
		node.Val = nums[1]

		if node.Val < root.Val {
			insertTree(node, root, leftHeight)
		} else {
			insertTree(node, root, rightHeight)
		}
	}

	return root
}

 // 平衡二叉树构造
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
		return nil
	}

	result := convertSortedArrayToBST(nums)

	return result
}
// @lc code=end

