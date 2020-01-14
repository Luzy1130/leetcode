/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 *
 * https://leetcode.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (41.25%)
 * Likes:    3482
 * Dislikes: 115
 * Total Accepted:    411.9K
 * Total Submissions: 992.4K
 * Testcase Example:  '[1,2,3,1]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed, the only constraint stopping
 * you from robbing each of them is that adjacent houses have security system
 * connected and it will automatically contact the police if two adjacent
 * houses were broken into on the same night.
 * 
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 * 
 * Example 2:
 * 
 * 
 * Input: [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house
 * 5 (money = 1).
 * Total amount you can rob = 2 + 9 + 1 = 12.
 * 
 * 
 */

// @lc code=start
// 即不能选择相邻的元素
// dp[i]记为经过第i个房子后能偷到的最大资金
// dp[i] = max(dp[i-2]+v[i], dp[i-1])
func rob(nums []int) int {
	res := 0
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[0]
	if nums[1] > nums[0] {
		dp[1] = nums[1]
	}

	for i := 2; i < len(nums); i++ {
		val1 := dp[i-2]
		val2 := dp[i-1]

		dp[i] = int(math.Max(float64(val1 + nums[i]),
			float64(val2)))

	}

	res = dp[len(dp)-1]
	return res
}
// @lc code=end

