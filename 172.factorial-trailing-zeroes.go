/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 *
 * https://leetcode.com/problems/factorial-trailing-zeroes/description/
 *
 * algorithms
 * Easy (37.55%)
 * Likes:    612
 * Dislikes: 866
 * Total Accepted:    181.6K
 * Total Submissions: 481.9K
 * Testcase Example:  '3'
 *
 * Given an integer n, return the number of trailing zeroes in n!.
 * 
 * Example 1:
 * 
 * 
 * Input: 3
 * Output: 0
 * Explanation: 3! = 6, no trailing zero.
 * 
 * Example 2:
 * 
 * 
 * Input: 5
 * Output: 1
 * Explanation: 5! = 120, one trailing zero.
 * 
 * Note: Your solution should be in logarithmic time complexity.
 * 
 */

// @lc code=start
func trailingZeroes(n int) int {
	cnt5 := 0
	for base := 5; base <= n; base=base*5 {
		cnt5 += (n / base)
	}

	return cnt5
}
// @lc code=end

