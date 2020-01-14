/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.29%)
 * Likes:    442
 * Dislikes: 1816
 * Total Accepted:    297.9K
 * Total Submissions: 922.3K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 *
 * Example:
 *
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 *
 *
 */
func lengthOfLastWord(s string) int {
	length := 0

	start := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if start == true {
				break
			} else {
				continue
			}
		}
		if start == false {
			start = true
		}
		length++
	}
	return length
}

