/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 *
 * https://leetcode.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (36.77%)
 * Likes:    1779
 * Dislikes: 146
 * Total Accepted:    172.5K
 * Total Submissions: 446.1K
 * Testcase Example:  '3'
 *
 * Given an integer n, generate all structurally unique BST's (binary search
 * trees) that store values 1 ... n.
 * 
 * Example:
 * 
 * 
 * Input: 3
 * Output:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * Explanation:
 * The above output corresponds to the 5 unique BST's shown below:
 * 
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
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


func generateTreesImpl(nums []int) []*TreeNode {
	var res []*TreeNode
	if len(nums) == 0 {
		res = append(res, nil)
		return res
	}

	for i := 0; i < len(nums); i++ {
		leftChilds := generateTreesImpl(nums[0:i])
		rightChilds := generateTreesImpl(nums[i+1:])
		for j := 0; j < len(leftChilds); j++ {
			for k := 0; k < len(rightChilds); k++ {
				root := &TreeNode{nums[i], nil, nil}
				root.Left = leftChilds[j]
				root.Right = rightChilds[k]
				res = append(res, root)
			}
		}
	}
	return res
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	var nums []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	res := generateTreesImpl(nums)
	return res
}
// @lc code=end

