/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 *
 * https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (42.78%)
 * Likes:    3102
 * Dislikes: 164
 * Total Accepted:    411.7K
 * Total Submissions: 952K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n1'
 *
 * Given a binary tree, find the lowest common ancestor (LCA) of two given
 * nodes in the tree.
 * 
 * According to the definition of LCA on Wikipedia: “The lowest common ancestor
 * is defined between two nodes p and q as the lowest node in T that has both p
 * and q as descendants (where we allow a node to be a descendant of itself).”
 * 
 * Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
 * Output: 3
 * Explanation: The LCA of nodes 5 and 1 is 3.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
 * Output: 5
 * Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant
 * of itself according to the LCA definition.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * All of the nodes' values will be unique.
 * p and q are different and both values will exist in the binary tree.
 * 
 * 
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

// 递归遍历
// 有冗余计算，记录每个节点的状态来减少一些重复计算
func findHelper(root, p, q *TreeNode, ans **TreeNode) int {
	if root == nil {
		return 0
	}

	// 左子树有能匹配几个
	left := findHelper(root.Left, p, q, ans)
	// 右子树能匹配几个
	right := findHelper(root.Right, p, q, ans)
	// 当前节点是否可以匹配
	count := left + right
	if root == p || root == q {
		count += 1
	} 

	// count表示节点的树匹配到的p或q的数量
	if count >= 2 {
		if *ans == nil {
			*ans = root
		}
	}

	// fmt.Println(count, ans, root, p, q)
	return count
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode = nil
	findHelper(root, p, q, &ans)
	return ans
}
// @lc code=end

