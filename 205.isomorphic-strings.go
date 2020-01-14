/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 *
 * https://leetcode.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (38.02%)
 * Likes:    1032
 * Dislikes: 290
 * Total Accepted:    252.6K
 * Total Submissions: 651.2K
 * Testcase Example:  '"egg"\n"add"'
 *
 * Given two strings s and t, determine if they are isomorphic.
 * 
 * Two strings are isomorphic if the characters in s can be replaced to get t.
 * 
 * All occurrences of a character must be replaced with another character while
 * preserving the order of characters. No two characters may map to the same
 * character but a character may map to itself.
 * 
 * Example 1:
 * 
 * 
 * Input: s = "egg", t = "add"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "foo", t = "bar"
 * Output: false
 * 
 * Example 3:
 * 
 * 
 * Input: s = "paper", t = "title"
 * Output: true
 * 
 * Note:
 * You may assume both s and t have the same length.
 * 
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}

	letterMap1 := make(map[byte]byte)
	letterMap2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if v, ok := letterMap1[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else if v, ok := letterMap2[t[i]]; ok {
			if v != s[i] {
				return false
			}
		} else {
			letterMap1[s[i]] = t[i]
			letterMap2[t[i]] = s[i]
		}
	}

	return true
}
// @lc code=end

