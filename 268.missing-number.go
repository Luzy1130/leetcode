/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 *
 * https://leetcode.com/problems/missing-number/description/
 *
 * algorithms
 * Easy (49.10%)
 * Likes:    1268
 * Dislikes: 1676
 * Total Accepted:    359.7K
 * Total Submissions: 721.1K
 * Testcase Example:  '[3,0,1]'
 *
 * Given an array containing n distinct numbers taken from 0, 1, 2, ..., n,
 * find the one that is missing from the array.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,0,1]
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [9,6,4,2,3,5,7,0,1]
 * Output: 8
 * 
 * 
 * Note:
 * Your algorithm should run in linear runtime complexity. Could you implement
 * it using only constant extra space complexity?
 */

// @lc code=start
func missingNumber(nums []int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	tmpSum := (len(nums)) * (len(nums)+1) / 2
	return tmpSum - sum
}
// @lc code=end

