/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (27.81%)
 * Likes:    5028
 * Dislikes: 443
 * Total Accepted:    749.3K
 * Total Submissions: 2.6M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 * 
 * Example 1:
 * 
 * 
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "cbbd"
 * Output: "bb"
 * 
 * 
 */

// @lc code=start

// 
// dp(i, j)=true ,id s[i:j] is a palindromic
// 

// DP
func longestPalindromeDp(s string) string {
	if len(s) == 0 {
		return ""
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	start := 0
	max := 1

	for i := len(s)-1; i >= 0; i-- {
		for j := i+1; j < len(s); j++ {
			if i + 1 > j - 1 {
				if s[i] == s[j] {
					dp[i][j] = true
				}
			} else {
				// fmt.Println("111",i,j, dp[i+1][j-1])
				if dp[i+1][j-1] && s[i] == s[j] {
					dp[i][j] = true
				}
			}

			if dp[i][j] {
				if max < j-i+1 {
					// fmt.Println(i, j)
					max = j-i+1
					start = i
				}
			}
		}
	}

	res := string(s[start:(start+max)])

	return res
}

// center expand
func expandCenter(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	len := r - l - 1
	return len
}
func longestPalindromeExpandCenter(s string) string {
	if len(s) == 0 {
		return ""
	}

	max := 0
	start := -1

	for i := 0; i < len(s); i++ {
		len1 := expandCenter(s, i, i)
		len2 := expandCenter(s, i, i+1)

		len := len1
		if len < len2 {
			len = len2
		}
		if len > max {
			max = len
			start = i - (len - 1) / 2
		}
	}

	res := string(s[start:(start+max)])
	return res
}


// TODO: manacher's algorithm


func longestPalindrome(s string) string {
	// return longestPalindromeDp(s)
	return longestPalindromeExpandCenter(s)
}
// @lc code=end

