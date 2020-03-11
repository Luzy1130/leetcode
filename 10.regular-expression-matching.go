/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 *
 * https://leetcode.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (25.60%)
 * Likes:    3619
 * Dislikes: 633
 * Total Accepted:    389.1K
 * Total Submissions: 1.5M
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement regular expression
 * matching with support for '.' and '*'.
 * 
 * 
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 * 
 * 
 * The matching should cover the entire input string (not partial).
 * 
 * Note:
 * 
 * 
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like . or *.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the preceding element, 'a'.
 * Therefore, by repeating 'a' once, it becomes "aa".
 * 
 * 
 * Example 3:
 * 
 * 
 * Input:
 * s = "ab"
 * p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 * 
 * 
 * Example 4:
 * 
 * 
 * Input:
 * s = "aab"
 * p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore,
 * it matches "aab".
 * 
 * 
 * Example 5:
 * 
 * 
 * Input:
 * s = "mississippi"
 * p = "mis*is*p*."
 * Output: false
 * 
 * 
 */

// @lc code=start
// package main

// import "fmt"

// 回溯法
func btMatch(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := (len(s) > 0 && (s[0] == p[0] || p[0] == '.'))
	if len(p) > 1 && p[1] == '*' {
		// 1. 假设*匹配长度为0
		// 2. *匹配了一个字符，再继续尝试匹配下一个字符
		return btMatch(s, p[2:]) || (firstMatch && btMatch(s[1:], p))
	} else {
		return firstMatch && btMatch(s[1:], p[1:])
	}
}

func isMatch(s string, p string) bool {
	return btMatch(s, p)
}

// func main() {
// 	// res := isMatch("ab", ".*")
// 	// res := isMatch("aa", "a")
// 	// res := isMatch("aa", "a*")
// 	// res := isMatch("aab", "c*a*b")
// 	res := isMatch("mississippi", "mis*is*p*.")
// 	fmt.Println(res)
// }


// @lc code=end

