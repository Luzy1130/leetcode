/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 *
 * https://leetcode.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (44.20%)
 * Likes:    2662
 * Dislikes: 59
 * Total Accepted:    466.3K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * Given a binary tree, check whether it is a mirror of itself (ie, symmetric
 * around its center).
 * 
 * For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
 * 
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 * 
 * 
 * 
 * 
 * But the following [1,2,2,null,3,null,3] is not:
 * 
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 * 
 * 
 * 
 * 
 * Note:
 * Bonus points if you could solve it both recursively and iteratively.
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


 // 递归算法
 func isSymmetricRecursionImpl(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil  && right == nil {
		return false
	}
	if left == nil && right != nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return isSymmetricRecursionImpl(left.Right, right.Left) &&
		isSymmetricRecursionImpl(left.Left, right.Right)
 }

func isSymmetricRecursion(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricRecursionImpl(root.Left, root.Right)
}

// 非递归算法：
func isSymmetricNonRecursion(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stackLeft := make([]*TreeNode, 0)
	stackRight := make([]*TreeNode, 0)
	posLeft := root.Left
	posRight := root.Right

	for (len(stackLeft) != 0 || posLeft != nil) || (len(stackRight) != 0 || posRight != nil) {
		for posLeft != nil || posRight != nil {
			if posLeft != nil && posRight == nil {
				return false
			}
			if posLeft == nil && posRight != nil {
				return false
			}
			if posLeft.Val != posRight.Val {
				return false
			}

			stackLeft = append(stackLeft, posLeft)
			stackRight = append(stackRight, posRight)
			posLeft = posLeft.Left
			posRight = posRight.Right
		}

		if len(stackLeft) != 0 {
			posLeft = stackLeft[len(stackLeft)-1]
			stackLeft = stackLeft[:len(stackLeft)-1]
			posLeft = posLeft.Right
		}

		if len(stackRight) != 0 {
			posRight = stackRight[len(stackRight)-1]
			stackRight = stackRight[:len(stackRight)-1]
			posRight = posRight.Left
		}
	}

	return true
}

func isSymmetric(root *TreeNode) bool {
    // return isSymmetricRecursion(root)

	return isSymmetricNonRecursion(root)
}

