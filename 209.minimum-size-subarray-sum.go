/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 *
 * https://leetcode.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (36.81%)
 * Likes:    1847
 * Dislikes: 96
 * Total Accepted:    237K
 * Total Submissions: 641.5K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * Given an array of n positive integers and a positive integer s, find the
 * minimal length of a contiguous subarray of which the sum ≥ s. If there isn't
 * one, return 0 instead.
 * 
 * Example: 
 * 
 * 
 * Input: s = 7, nums = [2,3,1,2,4,3]
 * Output: 2
 * Explanation: the subarray [4,3] has the minimal length under the problem
 * constraint.
 * 
 * Follow up:
 * 
 * If you have figured out the O(n) solution, try coding another solution of
 * which the time complexity is O(n log n). 
 * 
 */

 package main
 import (
	"fmt"
	"math"
 )

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sum := 0
	left := 0
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= s {
			if res > i - left +1 {
				res = i - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt32 {
		res = 0
	}
	return res
}
// @lc code=end

func main() {
	res := minSubArrayLen(7, []int{2,3,1,2,4,3})
	fmt.Println(res)
}
