/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 *
 * https://leetcode.com/problems/word-pattern/description/
 *
 * algorithms
 * Easy (35.48%)
 * Likes:    857
 * Dislikes: 120
 * Total Accepted:    166.4K
 * Total Submissions: 461.3K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * Given a pattern and a string str, find if str follows the same pattern.
 * 
 * Here follow means a full match, such that there is a bijection between a
 * letter in pattern and a non-empty word in str.
 * 
 * Example 1:
 * 
 * 
 * Input: pattern = "abba", str = "dog cat cat dog"
 * Output: true
 * 
 * Example 2:
 * 
 * 
 * Input:pattern = "abba", str = "dog cat cat fish"
 * Output: false
 * 
 * Example 3:
 * 
 * 
 * Input: pattern = "aaaa", str = "dog cat cat dog"
 * Output: false
 * 
 * Example 4:
 * 
 * 
 * Input: pattern = "abba", str = "dog dog dog dog"
 * Output: false
 * 
 * Notes:
 * You may assume pattern contains only lowercase letters, and str contains
 * lowercase letters that may be separated by a single space.
 * 
 */

// @lc code=start
func wordPattern(pattern string, str string) bool {
	var strArray []string
	
	var tmp []rune
	for _, v := range str {
		if v != ' ' {
			tmp = append(tmp, v)
		} else {
			if len(tmp) > 0 {
				strArray = append(strArray, string(tmp))
				tmp = tmp[0:0]
			}
		}
	}
	if len(tmp) > 0 {
		strArray = append(strArray, string(tmp))
	}
	fmt.Println(strArray)
	if len(pattern) != len(strArray) {
		return false
	}

	maps1 := make(map[rune]string)
	maps2 := make(map[string]rune)

	for i, v := range pattern {
		if str, ok := maps1[v]; ok {
			if str != strArray[i] {
				return false
			}
		} else {
			if _, ok := maps2[strArray[i]]; ok {
				return false
			}
			maps1[v] = strArray[i]
			maps2[strArray[i]] = v
		}
	}
	return true
}
// @lc code=end

