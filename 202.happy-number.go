/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 *
 * https://leetcode.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (46.30%)
 * Likes:    1267
 * Dislikes: 311
 * Total Accepted:    304.4K
 * Total Submissions: 636.1K
 * Testcase Example:  '19'
 *
 * Write an algorithm to determine if a number is "happy".
 * 
 * A happy number is a number defined by the following process: Starting with
 * any positive integer, replace the number by the sum of the squares of its
 * digits, and repeat the process until the number equals 1 (where it will
 * stay), or it loops endlessly in a cycle which does not include 1. Those
 * numbers for which this process ends in 1 are happy numbers.
 * 
 * Example: 
 * 
 * 
 * Input: 19
 * Output: true
 * Explanation: 
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 * 
 */

// @lc code=start
func getNext(n int) int {
	res := 0

	for n > 0 {
		num := n % 10;
		n = n / 10
		if num > 0 {
			res += num * num
		}
	}
	return res
}

func isHappy(n int) bool {
	seen := make(map[int]bool)
	
	for n != 1 {
		if _, ok := seen[n]; ok {
			return false
		} else {
			seen[n] = true
			n = getNext(n)
		}
	}
	return true
}
// @lc code=end

