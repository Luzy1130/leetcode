/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 *
 * https://leetcode.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (27.41%)
 * Likes:    3025
 * Dislikes: 130
 * Total Accepted:    263.2K
 * Total Submissions: 953.1K
 * Testcase Example:  '"(()"'
 *
 * Given a string containing just the characters '(' and ')', find the length
 * of the longest valid (well-formed) parentheses substring.
 * 
 * Example 1:
 * 
 * 
 * Input: "(()"
 * Output: 2
 * Explanation: The longest valid parentheses substring is "()"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: ")()())"
 * Output: 4
 * Explanation: The longest valid parentheses substring is "()()"
 * 
 * 
 */
package main
import "fmt"
// @lc code=start
func longestValidParentheses(s string) int {
    if len(s) == 0 {
		return 0
	}
	max := 0
	left := 0
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
			dp[i] = 0
		} else {
			if left > 0 {
				if s[i-1] == '(' {
					if i == 1 {
						dp[i] = 2
					} else {
						dp[i] = 2 + dp[i-2]
					}
				} else {
					dp[i] = 2 + dp[i-1]
					if i - dp[i] >= 0 {
						dp[i] += dp[i-dp[i]]
					}
				}
				left--
			} else {
				dp[i] = 0
			}

			if dp[i] > max {
				max = dp[i]
			}
		}
	}

	// fmt.Println(dp)
	return max
}
// @lc code=end

func main() {
	res := longestValidParentheses("()(())")
	fmt.Println(res)
}