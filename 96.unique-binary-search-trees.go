/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 *
 * https://leetcode.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (47.46%)
 * Likes:    2602
 * Dislikes: 98
 * Total Accepted:    253.4K
 * Total Submissions: 508.6K
 * Testcase Example:  '3'
 *
 * Given n, how many structurally unique BST's (binary search trees) that store
 * values 1 ... n?
 * 
 * Example:
 * 
 * 
 * Input: 3
 * Output: 5
 * Explanation:
 * Given n = 3, there are a total of 5 unique BST's:
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

// package main
// import "fmt"

// func generateTreesImpl(nums []int) []*TreeNode {
// 	var res []*TreeNode
// 	if len(nums) == 0 {
// 		res = append(res, nil)
// 		return res
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		leftChilds := generateTreesImpl(nums[0:i])
// 		rightChilds := generateTreesImpl(nums[i+1:])
// 		for j := 0; j < len(leftChilds); j++ {
// 			for k := 0; k < len(rightChilds); k++ {
// 				root := &TreeNode{nums[i], nil, nil}
// 				root.Left = leftChilds[j]
// 				root.Right = rightChilds[k]
// 				res = append(res, root)
// 			}
// 		}
// 	}
// 	return res
// }

// // 分治
// func divid(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	var nums []int
// 	for i := 1; i <= n; i++ {
// 		nums = append(nums, i)
// 	}

// 	trees := generateTreesImpl(nums)
// 	return len(trees)
// }


// dp, 只是计算数量的话，分治算法有很多重复的计算
func dpImpl(n int, dpNums []int) int {
	if n == 0 {
		return 1
	}
	if dpNums[n] != 0 {
		return dpNums[n]
	}

	num := 0
	for i := 1; i <= n; i++ {
		numLeft := dpImpl(i-1, dpNums)
		numRight := dpImpl(n-i, dpNums)
		num += numLeft * numRight 
	}
	dpNums[n] = num
	return num
}

func dp(n int) int {
	dpNums := make([]int, n+1)
	
	res := dpImpl(n, dpNums)

	return res
}

func numTrees(n int) int {
	// return divide(n)
	return dp(n)
}

// func main() {
// 	res := numTrees(19)
// 	fmt.Println(res)
// }
// @lc code=end

