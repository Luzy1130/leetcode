/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (52.32%)
 * Likes:    1641
 * Dislikes: 163
 * Total Accepted:    325.7K
 * Total Submissions: 602.6K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given an array where elements are sorted in ascending order, convert it to a
 * height balanced BST.
 * 
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 * 
 * Example:
 * 
 * 
 * Given the sorted array: [-10,-3,0,5,9],
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
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// package main
// import "fmt"

func sortedArrayToBSTImpl(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}

	mid := (low+high) / 2
	left := sortedArrayToBSTImpl(nums, low, mid-1)
	right := sortedArrayToBSTImpl(nums, mid+1, high)

	root := &TreeNode{nums[mid], left, right}
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
    return sortedArrayToBSTImpl(nums, 0, len(nums)-1);
}

// func main() {
// 	input := []int{-10, -3, 0, 5, 9}
// 	res := sortedArrayToBST(input)
// 	fmt.Println(res)
// }
// @lc code=end

