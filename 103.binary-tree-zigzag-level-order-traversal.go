/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (42.96%)
 * Likes:    1597
 * Dislikes: 85
 * Total Accepted:    309.4K
 * Total Submissions: 681.6K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the zigzag level order traversal of its nodes'
 * values. (ie, from left to right, then right to left for the next level and
 * alternate between).
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
 * return its zigzag level order traversal as:
 * 
 * [
 * ⁠ [3],
 * ⁠ [20,9],
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
func zigzagLevelOrder(root *TreeNode) [][]int {
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

		for j := len(queue[i])-1; j >= 0; j-- {
			node := queue[i][j]
			tmp = append(tmp, node.Val)
			if i == 0 {
				if node.Left != nil {
					queue[nextIndex] = append(queue[nextIndex], node.Left)
				}
				if node.Right != nil {
					queue[nextIndex] = append(queue[nextIndex], node.Right)
				}
			} else {
				if node.Right != nil {
					queue[nextIndex] = append(queue[nextIndex], node.Right)
				}
				if node.Left != nil {
					queue[nextIndex] = append(queue[nextIndex], node.Left)
				}
			}
		}

		queue[i] = queue[i][0:0]
		i = nextIndex
		res = append(res, tmp)
	}
	
	return res
}
// @lc code=end

