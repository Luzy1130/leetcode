/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 *
 * https://leetcode.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (44.97%)
 * Likes:    1455
 * Dislikes: 56
 * Total Accepted:    205.6K
 * Total Submissions: 456.6K
 * Testcase Example:  '"aab"'
 *
 * Given a string s, partition s such that every substring of the partition is
 * a palindrome.
 * 
 * Return all possible palindrome partitioning of s.
 * 
 * Example:
 * 
 * 
 * Input: "aab"
 * Output:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"

func isPalindrome(s string) bool {
	i := 0
	j := len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func partition(s string) [][]string {
	var res [][]string
	if len(s) == 0 {
		return res
	}

	for i := 0; i < len(s); i++ {
		tmp := s[0:i+1]
		tmpStr := string(tmp)
		if isPalindrome(tmpStr) != true {
			continue
		}
		if i == len(s) -1 {
			res = append(res, []string{tmpStr})
			continue
		}

		suffixes := partition(string(s[i+1:]))
		for _, suffix := range suffixes {
			res = append(res, append([]string{tmpStr}, suffix...))
		}
	}
	return res
}

// func main() {
// 	res := partition("abb")
// 	fmt.Println(res)
// }
// @lc code=end

