/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 *
 * https://leetcode.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (51.79%)
 * Likes:    1769
 * Dislikes: 100
 * Total Accepted:    244.3K
 * Total Submissions: 469.7K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * Given a binary tree, imagine yourself standing on the right side of it,
 * return the values of the nodes you can see ordered from top to bottom.
 * 
 * Example:
 * 
 * 
 * Input: [1,2,3,null,5,null,4]
 * Output: [1, 3, 4]
 * Explanation:
 * 
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
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
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) != 0 {
		res = append(res, queue[len(queue)-1].Val)
		var tmp []*TreeNode

		for _, n := range queue {
			if n.Left != nil {
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil {
				tmp = append(tmp, n.Right)
			}
		}

		queue = tmp
	}

	return res
}
// @lc code=end

