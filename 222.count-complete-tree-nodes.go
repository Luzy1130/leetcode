/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 *
 * https://leetcode.com/problems/count-complete-tree-nodes/description/
 *
 * algorithms
 * Medium (41.12%)
 * Likes:    1677
 * Dislikes: 179
 * Total Accepted:    187.3K
 * Total Submissions: 450.2K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * Given a complete binary tree, count the number of nodes.
 * 
 * Note: 
 * 
 * Definition of a complete binary tree from Wikipedia:
 * In a complete binary tree every level, except possibly the last, is
 * completely filled, and all nodes in the last level are as far left as
 * possible. It can have between 1 and 2^h nodes inclusive at the last level
 * h.
 * 
 * Example:
 * 
 * 
 * Input: 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠/ \  /
 * 4  5 6
 * 
 * Output: 6
 * 
 */

package main
import "math"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func exists(root *TreeNode, index int) bool {
	var path []int
	for index > 1 {
		if index % 2 == 0 {
			path = append(path, 0)
		} else {
			path = append(path, 1)
		}
		index = index / 2
	}

	cur := root
	for i := len(path)-1; i >= 0; i-- {
		var next *TreeNode
		if path[i] == 0 {
			next = cur.Left
		} else {
			next = cur.Right
		}
		if next == nil {
			return false
		}
		cur = next
	}
	return true
}

func countNodes(root *TreeNode) int {
    if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depth := -1
	for cur := root; cur != nil; cur = cur.Left {
		depth++
	}

	// 计算最后一层的节点个数, 6, 7
	l := int(math.Pow(2.0, float64(depth)))
	r := int(math.Pow(2.0, float64(depth+1))) - 1
	for l < r {
		mid := (l + r) / 2
		if (exists(root, mid)) {
			l =  mid+1
		} else {
			r = mid - 1
		}
	}

	if !exists(root, l) {
		l -= 1
	}
	return l
}
// @lc code=end

