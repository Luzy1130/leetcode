/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 *
 * https://leetcode.com/problems/sort-colors/description/
 *
 * algorithms
 * Medium (43.05%)
 * Likes:    2419
 * Dislikes: 189
 * Total Accepted:    403.8K
 * Total Submissions: 908.1K
 * Testcase Example:  '[2,0,2,1,1,0]'
 *
 * Given an array with n objects colored red, white or blue, sort them in-place
 * so that objects of the same color are adjacent, with the colors in the order
 * red, white and blue.
 * 
 * Here, we will use the integers 0, 1, and 2 to represent the color red,
 * white, and blue respectively.
 * 
 * Note: You are not suppose to use the library's sort function for this
 * problem.
 * 
 * Example:
 * 
 * 
 * Input: [2,0,2,1,1,0]
 * Output: [0,0,1,1,2,2]
 * 
 * Follow up:
 * 
 * 
 * A rather straight forward solution is a two-pass algorithm using counting
 * sort.
 * First, iterate the array counting number of 0's, 1's, and 2's, then
 * overwrite array with total number of 0's, then 1's and followed by 2's.
 * Could you come up with a one-pass algorithm using only constant space?
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"

func sortColors(nums []int)  {
	s := 0
	e := len(nums)-1
	for i := 0; i <=e; {
		switch nums[i] {
		case 0:
			nums[i] = 1
			nums[s] = 0
			s++
			i++
		case 2:
			nums[i], nums[e] = nums[e], nums[i]
			e--
		default:
			i++
		}
	}
}

// func main() {
// 	var input []int = []int{2,0,2,1,1,0}
// 	sortColors(input)
// 	fmt.Println(input)
// }
// @lc code=end

