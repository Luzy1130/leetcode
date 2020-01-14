/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (34.04%)
 * Likes:    1574
 * Dislikes: 1437
 * Total Accepted:    528.9K
 * Total Submissions: 1.6M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 * 
 * If there is no common prefix, return an empty string "".
 * 
 * Example 1:
 * 
 * 
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 * 
 * 
 * Note:
 * 
 * All given inputs are in lowercase letters a-z.
 * 
 */
func longestCommonPrefix(strs []string) string {
	var result []byte
	
	i := 0
	j := 0
	var tmp byte
	for ; i < len(strs); {
		if j >= len(strs[i]) {
			break
		}

		if i == 0 {
			tmp = strs[i][j]
		} else {
			if tmp != strs[i][j] {
				break
			}
		}

		if i == len(strs) - 1 {
			result = append(result, tmp)
			j++
			i = 0
		} else {
			i++
		}
	}

	return string(result)
}

