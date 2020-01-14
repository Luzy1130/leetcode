/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (32.12%)
 * Likes:    808
 * Dislikes: 2309
 * Total Accepted:    457.6K
 * Total Submissions: 1.4M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * Given a string, determine if it is a palindrome, considering only
 * alphanumeric characters and ignoring cases.
 * 
 * Note: For the purpose of this problem, we define empty string as valid
 * palindrome.
 * 
 * Example 1:
 * 
 * 
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "race a car"
 * Output: false
 * 
 * 
 */

// @lc code=start
func isPalindrome(s string) bool {
	var tmp []rune
    for _, c := range s {
		if unicode.IsLetter(c) {
			tmp = append(tmp, unicode.ToLower(c))
		}
		if unicode.IsDigit(c) {
			tmp = append(tmp, c)
		}
	}
	ss := string(tmp)

	i := 0
	j := len(ss) - 1
	for i < j {
		if ss[i] != ss[j] {
			return false
		}
		i++
		j--
	}

	return true
}
// @lc code=end

