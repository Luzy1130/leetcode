/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 *
 * https://leetcode.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (26.18%)
 * Likes:    3175
 * Dislikes: 450
 * Total Accepted:    593.4K
 * Total Submissions: 2.2M
 * Testcase Example:  '[2,1,3]'
 *
 * Given a binary tree, determine if it is a valid binary search tree (BST).
 * 
 * Assume a BST is defined as follows:
 * 
 * 
 * The left subtree of a node contains only nodes with keys less than the
 * node's key.
 * The right subtree of a node contains only nodes with keys greater than the
 * node's key.
 * Both the left and right subtrees must also be binary search trees.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 
 * Input: [2,1,3]
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * ⁠    5
 * ⁠   / \
 * ⁠  1   4
 *  / \
 * 3   6
 * 
 * Input: [5,1,4,null,null,3,6]
 * Output: false
 * Explanation: The root node's value is 5 but its right child's value is 4.
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
// // 递归
// func isValidBSTImpl1(root *TreeNode, small, large int) bool {
// 	if root == nil {
// 		return true
// 	}
// 	val := root.Val
// 	if small != -999999999 && val <= small {
// 		return false
// 	} 
// 	if large != -999999999 && val >= large {
// 		return false
// 	}

// 	//  相对于右子树，root节点应该是最小的
// 	if !isValidBSTImpl1(root.Right, val, large) {
// 		return false
// 	}
// 	// 相对于左子树，root节点应该是最大的
// 	if !isValidBSTImpl1(root.Left, small, val) {
// 		return false
// 	}

// 	return true
// }

// 中序遍历应该有序
func isValidBSTImpl2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var stack []*TreeNode
	pos := root
	curVal := 0
	begin := false
	for pos != nil || len(stack) != 0 {
		for pos != nil {
			stack = append(stack, pos)
			pos = pos.Left
		}

		if len(stack) != 0 {
			pos = stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			if begin == false {
				begin = true
				curVal = pos.Val
			} else if pos.Val <= curVal {
				return false
			}
			curVal = pos.Val
			pos = pos.Right
		}
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	// return isValidBSTImpl1(root, -999999999, -999999999)
	return isValidBSTImpl2(root)
}
// @lc code=end

