/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (47.99%)
 * Likes:    940
 * Dislikes: 177
 * Total Accepted:    265.5K
 * Total Submissions: 539.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
 * 
 * 
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 
 * return its bottom-up level order traversal as:
 * 
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
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
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	var queue1 []*TreeNode

	if root == nil {
		return res
	}

	queue1 = append(queue1, root)
	for len(queue1) != 0 {
		var row []int
		var queue2 []*TreeNode
		for len(queue1) != 0 {
			node := queue1[0]
			queue1 = queue1[1:]
			row = append(row, node.Val)
			if node.Left != nil {
				queue2 = append(queue2, node.Left)
			}
			if node.Right != nil {
				queue2 = append(queue2, node.Right)
			}
		}
		res = append(res, row)
		queue1 = queue2
	}

	// 翻转
	i := 0
	j := len(res) -1 
	for ;i != j && i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}
// @lc code=end

