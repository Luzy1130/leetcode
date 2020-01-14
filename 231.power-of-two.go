/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 *
 * https://leetcode.com/problems/power-of-two/description/
 *
 * algorithms
 * Easy (42.27%)
 * Likes:    570
 * Dislikes: 150
 * Total Accepted:    263.5K
 * Total Submissions: 617.5K
 * Testcase Example:  '1'
 *
 * Given an integer, write a function to determine if it is a power of two.
 * 
 * Example 1:
 * 
 * 
 * Input: 1
 * Output: true 
 * Explanation: 2^0 = 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 16
 * Output: true
 * Explanation: 2^4 = 16
 * 
 * Example 3:
 * 
 * 
 * Input: 218
 * Output: false
 * 
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	for i := 1; i <= n; i = i << 1 {
		if i == n {
			return true
		}
	}
	return false

}
// @lc code=end

