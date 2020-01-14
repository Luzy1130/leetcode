/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (53.28%)
 * Likes:    1041
 * Dislikes: 126
 * Total Accepted:    449.2K
 * Total Submissions: 822.3K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t , write a function to determine if t is an anagram
 * of s.
 * 
 * Example 1:
 * 
 * 
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "rat", t = "car"
 * Output: false
 * 
 * 
 * Note:
 * You may assume the string contains only lowercase alphabets.
 * 
 * Follow up:
 * What if the inputs contain unicode characters? How would you adapt your
 * solution to such case?
 * 
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	letterMap := make(map[byte]int)
	
	for i := 0; i < len(s); i++ {
		if _, ok := letterMap[s[i]]; ok {
			letterMap[s[i]]++
		} else {
			letterMap[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := letterMap[t[i]]; ok {
			letterMap[t[i]]--
			if letterMap[t[i]] == 0 {
				delete(letterMap, t[i])
			}
		} else {
			return false
		}
	}
	return true
}
// @lc code=end

