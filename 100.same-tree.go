/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 *
 * https://leetcode.com/problems/same-tree/description/
 *
 * algorithms
 * Easy (50.61%)
 * Likes:    1316
 * Dislikes: 43
 * Total Accepted:    419.8K
 * Total Submissions: 826.6K
 * Testcase Example:  '[1,2,3]\n[1,2,3]'
 *
 * Given two binary trees, write a function to check if they are the same or
 * not.
 * 
 * Two binary trees are considered the same if they are structurally identical
 * and the nodes have the same value.
 * 
 * Example 1:
 * 
 * 
 * Input:     1         1
 * ⁠         / \       / \
 * ⁠        2   3     2   3
 * 
 * ⁠       [1,2,3],   [1,2,3]
 * 
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:     1         1
 * ⁠         /           \
 * ⁠        2             2
 * 
 * ⁠       [1,2],     [1,null,2]
 * 
 * Output: false
 * 
 * 
 * Example 3:
 * 
 * 
 * Input:     1         1
 * ⁠         / \       / \
 * ⁠        2   1     1   2
 * 
 * ⁠       [1,2,1],   [1,1,2]
 * 
 * Output: false
 * 
 * 
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTreeRecursion(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q == nil && p != nil {
		return false
	} 
	if q != nil && p == nil {
		return false
	}
	if q.Val != p.Val {
		return false
	}

	return (isSameTreeRecursion(q.Left, p.Left) && isSameTreeRecursion(q.Right, p.Right))
}

func isSameTreeNonRecursion(p *TreeNode, q *TreeNode) bool {
	stack1 := make([]*TreeNode, 0)
	stack2 := make([]*TreeNode, 0)

	pos1 := p
	pos2 := q
	
	for (len(stack1) != 0 || pos1 != nil) || (len(stack2) != 0 || pos2 != nil) {
		for pos1 != nil || pos2 != nil {
			if pos1 == nil && pos2 != nil {
				return false
			}
			if pos1 != nil && pos2 == nil {
				return false
			}
			if (pos1.Val != pos2.Val) {
				return false
			}
			stack1 = append(stack1, pos1)
			stack2 = append(stack2, pos2)
			pos1 = pos1.Left
			pos2 = pos2.Left
		}

		if len(stack1) > 0 {
			pos1 = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			if pos1 != nil {
				pos1 = pos1.Right
			}
		}
		if len(stack2) > 0 {
			pos2 = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
			if pos2 != nil {
				pos2 = pos2.Right
			}
		}
	}

	return true
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 递归方式
	// return isSameTreeRecursion(p, q)
	// 非递归方式：两颗树使用相同的遍历方式，每一步都要求相同
	return isSameTreeNonRecursion(p, q)
}

