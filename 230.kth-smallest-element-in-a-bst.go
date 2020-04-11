/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 *
 * https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (56.16%)
 * Likes:    1939
 * Dislikes: 53
 * Total Accepted:    319.5K
 * Total Submissions: 564.6K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * Given a binary search tree, write a function kthSmallest to find the kth
 * smallest element in it.
 * 
 * Note: 
 * You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
 * 
 * Example 1:
 * 
 * 
 * Input: root = [3,1,4,null,2], k = 1
 * ⁠  3
 * ⁠ / \
 * ⁠1   4
 * ⁠ \
 * 2
 * Output: 1
 * 
 * Example 2:
 * 
 * 
 * Input: root = [5,3,6,2,4,null,null,1], k = 3
 * ⁠      5
 * ⁠     / \
 * ⁠    3   6
 * ⁠   / \
 * ⁠  2   4
 * ⁠ /
 * ⁠1
 * Output: 3
 * 
 * 
 * Follow up:
 * What if the BST is modified (insert/delete operations) often and you need to
 * find the kth smallest frequently? How would you optimize the kthSmallest
 * routine?
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
// 利用中序遍历的思想
// 如果按照题中所提示的，假设tree的操作很频繁，则可以参考数据库索引B+tree的做法
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[0: len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return 0
}
// @lc code=end

