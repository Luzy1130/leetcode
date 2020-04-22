/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 *
 * https://leetcode.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (35.95%)
 * Likes:    1423
 * Dislikes: 44
 * Total Accepted:    155.4K
 * Total Submissions: 431.2K
 * Testcase Example:  '[2,3,2]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed. All houses at this place are
 * arranged in a circle. That means the first house is the neighbor of the last
 * one. Meanwhile, adjacent houses have security system connected and it will
 * automatically contact the police if two adjacent houses were broken into on
 * the same night.
 * 
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 * 
 * Example 1:
 * 
 * 
 * Input: [2,3,2]
 * Output: 3
 * Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money
 * = 2),
 * because they are adjacent houses.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 * 
 */

package main
import "fmt"

// 这里的难点是有一个”环”
// 分成两个问题：
// 1. 第一家不偷的情况下，最多收益P1；
// 2. 最后一家不偷的情况下，最多收益P2
// 取max(P1, P2)
// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums)-1)
	p1 := 0
	p2 := 0
	// 第一家不偷
	dp[0] = nums[1]
	for i := 2; i < len(nums); i++ {
		if i == 2 {
			if nums[i] > nums[i-1] {
				dp[i-1] = nums[i]
			} else {
				dp[i-1] = dp[i-2]
			}
 		} else {
			 dp[i-1] = dp[i-3] + nums[i]
			 if dp[i-1] < dp[i-2] {
				dp[i-1] = dp[i-2]
			 }
		 }
	}
	p1 = dp[len(dp)-1]


	// 最后一家不偷
	dp[0] = nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if i == 1 {
			if nums[i] > nums[i-1] {
				dp[i] = nums[i]
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			dp[i] = dp[i-2] + nums[i]
			if dp[i] < dp[i-1] {
				dp[i] = dp[i-1]
			}
		}
	}
	p2 = dp[len(dp)-1]

	if p1 < p2 {
		p1 = p2
	}
	return p1
}
// @lc code=end

func main() {
	input := []int{1,2,3,1}
	res := rob(input)
	fmt.Println(res)
}