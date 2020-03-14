/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (45.16%)
 * Likes:    1586
 * Dislikes: 81
 * Total Accepted:    220.4K
 * Total Submissions: 488K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given a singly linked list where elements are sorted in ascending order,
 * convert it to a height balanced BST.
 * 
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 * 
 * Example:
 * 
 * 
 * Given the sorted linked list: [-10,-3,0,5,9],
 * 
 * One possible answer is: [0,-3,9,-10,null,5], which represents the following
 * height balanced BST:
 * 
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 * 
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// package main
// import "fmt"

var newHead *ListNode

// 利用了中序遍历的思想
// 还有很多方法，重点都是需要找到中间的节点，然后分治
func helper(low, high int) *TreeNode {
	if low > high {
		return nil
	}

	mid := (low + high) / 2
	root := &TreeNode{}
	root.Left = helper(low, mid-1)
	
	fmt.Println(newHead.Val)
	root.Val = newHead.Val
	newHead = newHead.Next

	root.Right = helper(mid+1, high)

	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	l := 0
	pos := head
	for ; pos != nil; pos = pos.Next {
		l++
	}
	newHead = head
	return helper(0, l-1)
}
// @lc code=end

