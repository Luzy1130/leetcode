/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (49.84%)
 * Likes:    2360
 * Dislikes: 62
 * Total Accepted:    527.7K
 * Total Submissions: 1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
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
 * return its level order traversal as:
 * 
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
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
func levelOrder(root *TreeNode) [][]int {
	var queue [][]*TreeNode
	if root == nil {
		return [][]int{}
	}
	queue = append(queue, []*TreeNode{})
	queue = append(queue, []*TreeNode{})

	var res [][]int
	i := 0
	queue[i] = append(queue[i], root)
	for len(queue[i]) != 0 {
		var tmp []int
		nextIndex := 1^i
		for j := 0; j < len(queue[i]); j++ {
			node := queue[i][j]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue[nextIndex] = append(queue[nextIndex], node.Left)
			}
			if node.Right != nil {
				queue[nextIndex] = append(queue[nextIndex], node.Right)
			}
		}
		queue[i] = queue[i][0:0]
		i = nextIndex
		res = append(res, tmp)
	}
	
	return res
}
// @lc code=end

