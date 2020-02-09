/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (54.66%)
 * Likes:    2852
 * Dislikes: 67
 * Total Accepted:    477.1K
 * Total Submissions: 828.3K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct integers, nums, return all possible subsets (the
 * power set).
 * 
 * Note: The solution set must not contain duplicate subsets.
 * 
 * Example:
 * 
 * 
 * Input: nums = [1,2,3]
 * Output:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 * 
 */

// @lc code=start
// package main
// import "fmt"

func subSetsImpl(nums []int, prefix []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	
	for i, _ := range nums {
		tmpPrefix := make([]int, 0)
		for _, v := range prefix {
			tmpPrefix = append(tmpPrefix, v)
		}
		mem := append(tmpPrefix, nums[i])
		res = append(res, mem)
		tmpRes := subSetsImpl(nums[(i+1):], mem)
		for _, v := range tmpRes {
			res = append(res, v)
		}
	}
	return res
}

func subsets(nums []int) [][]int {
	res := subSetsImpl(nums, []int{})
	res = append(res, []int{})
	return res
}

// func main() {
// 	var input []int = []int{ 1, 2, 3 }
// 	res := subsets(input)
// 	fmt.Println(res)
// }
// @lc code=end

