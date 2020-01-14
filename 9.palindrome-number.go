/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (44.62%)
 * Likes:    1594
 * Dislikes: 1362
 * Total Accepted:    658.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 * 
 * Example 1:
 * 
 * 
 * Input: 121
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 * 
 * 
 * Follow up:
 * 
 * Coud you solve it without converting the integer to a string?
 * 
 */
func isPalindrome(x int) bool {
	// solve it without converting the integer to a string
	// 基本思路是将源数字倒序后再比较，但是只需要将一般长度的数字倒序就
	// 可以进行比较了，不需要完全倒序。
	if x < 0 {
		return false
	}
	if x >= 0 && x < 10 {
		return true
	}
	if x % 10 == 0 {
		return false
	}

	num := 0
	// 倒一半
	for ; x != 0 && x > num; x = x / 10 {
		num = num * 10 + x % 10
	}

	if x == num || num / 10 == x {
		return true
	}
	return false
}

