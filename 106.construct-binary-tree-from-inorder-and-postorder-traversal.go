/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (43.29%)
 * Likes:    1293
 * Dislikes: 28
 * Total Accepted:    194.3K
 * Total Submissions: 448.8K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * Given inorder and postorder traversal of a tree, construct the binary tree.
 * 
 * Note:
 * You may assume that duplicates do not exist in the tree.
 * 
 * For example, given
 * 
 * 
 * inorder = [9,3,15,20,7]
 * postorder = [9,15,7,20,3]
 * 
 * Return the following binary tree:
 * 
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1]
	root := &TreeNode{val, nil, nil}
	leftIndex := 0
	for ;inorder[leftIndex] != val; {
		leftIndex++
	}

	root.Left = buildTree(inorder[0:leftIndex], postorder[0:leftIndex])
	root.Right = buildTree(inorder[leftIndex+1:], postorder[leftIndex: len(postorder)-1])

	return root
}
// @lc code=end

