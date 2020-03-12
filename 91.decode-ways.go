/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 *
 * https://leetcode.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (22.77%)
 * Likes:    2155
 * Dislikes: 2408
 * Total Accepted:    348.3K
 * Total Submissions: 1.5M
 * Testcase Example:  '"12"'
 *
 * A message containing letters from A-Z is being encoded to numbers using the
 * following mapping:
 * 
 * 
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 * 
 * 
 * Given a non-empty string containing only digits, determine the total number
 * of ways to decode it.
 * 
 * Example 1:
 * 
 * 
 * Input: "12"
 * Output: 2
 * Explanation: It could be decoded as "AB" (1 2) or "L" (12).
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "226"
 * Output: 3
 * Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2
 * 6).
 * 
 */

// @lc code=start
// package main
import (
	"fmt"
	// "strconv"
)

func numDecodings(s string) int {
    if len(s) == 0 {
		return 0
	}

	l := len(s)
	var dp []int = make([]int, l)
	if s[l-1] != '0' {
		dp[0] = 1
	}
	
	for i := 1; i < len(s); i++ {
		if s[l-i-1] == '0' {
			continue
		}
		num1 := dp[i-1]
		num2 := 0
		n, _ := strconv.Atoi(string(s[l-i-1:l-i+1]))
		// fmt.Println(n)
		if n > 0 && n <= 26 {
			if i > 1 {
				num2 = dp[i-2]
			} else {
				num2 = 1
			}
		}
		dp[i] = num1 +num2
	}
	return dp[len(s)-1]
}

// func main() {
// 	res := numDecodings("1010")
// 	fmt.Println(res)
// }
// @lc code=end

