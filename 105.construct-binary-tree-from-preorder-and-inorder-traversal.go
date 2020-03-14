/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (45.98%)
 * Likes:    2667
 * Dislikes: 77
 * Total Accepted:    308.8K
 * Total Submissions: 671.6K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * Given preorder and inorder traversal of a tree, construct the binary tree.
 * 
 * Note:
 * You may assume that duplicates do not exist in the tree.
 * 
 * For example, given
 * 
 * 
 * preorder = [3,9,20,15,7]
 * inorder = [9,3,15,20,7]
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
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	indexInorder := 0
	for ; inorder[indexInorder] != preorder[0]; {
		indexInorder++
	}

	root.Left = buildTree(preorder[1:1+indexInorder], inorder[0:indexInorder])
	root.Right = buildTree(preorder[indexInorder+1:], inorder[indexInorder+1:])
	
	return root
}
// @lc code=end

