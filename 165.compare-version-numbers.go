/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 *
 * https://leetcode.com/problems/compare-version-numbers/description/
 *
 * algorithms
 * Medium (25.99%)
 * Likes:    408
 * Dislikes: 1321
 * Total Accepted:    168.9K
 * Total Submissions: 647.3K
 * Testcase Example:  '"0.1"\n"1.1"'
 *
 * Compare two version numbers version1 and version2.
 * If version1 > version2 return 1; if version1 < version2 return -1;otherwise
 * return 0.
 * 
 * You may assume that the version strings are non-empty and contain only
 * digits and the . character.
 * The . character does not represent a decimal point and is used to separate
 * number sequences.
 * For instance, 2.5 is not "two and a half" or "half way to version three", it
 * is the fifth second-level revision of the second first-level revision.
 * You may assume the default revision number for each level of a version
 * number to be 0. For example, version number 3.4 has a revision number of 3
 * and 4 for its first and second level revision number. Its third and fourth
 * level revision number are both 0.
 * 
 * 
 * 
 * Example 1:
 * 
 * Input: version1 = "0.1", version2 = "1.1"
 * Output: -1
 * 
 * Example 2:
 * 
 * Input: version1 = "1.0.1", version2 = "1"
 * Output: 1
 * 
 * Example 3:
 * 
 * Input: version1 = "7.5.2.4", version2 = "7.5.3"
 * Output: -1
 * 
 * Example 4:
 * 
 * Input: version1 = "1.01", version2 = "1.001"
 * Output: 0
 * Explanation: Ignoring leading zeroes, both “01” and “001" represent the same
 * number “1”
 * 
 * Example 5:
 * 
 * Input: version1 = "1.0", version2 = "1.0.0"
 * Output: 0
 * Explanation: The first version number does not have a third level revision
 * number, which means its third level revision number is default to "0"
 * 
 * 
 * 
 * Note:
 * 
 * Version strings are composed of numeric strings separated by dots . and this
 * numeric strings may have leading zeroes. 
 * Version strings do not start or end with dots, and they will not be two
 * consecutive dots.
 * 
 */
 package main
 import (
	"fmt"
	"strconv"
 )

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	i := 0
	j := 0
	for i < len(version1) || j < len(version2) {
		var tmp1 []rune
		var v1 int = 0
		for i < len(version1) && version1[i] != '.' {
			tmp1 = append(tmp1, rune(version1[i]))
			i++
		}
		if len(tmp1) != 0 {
			v1, _ = strconv.Atoi(string(tmp1))
			i++
		}

		var v2 int = 0
		var tmp2 []rune
		for j < len(version2) && version2[j] != '.' {
			tmp2 = append(tmp2, rune(version2[j]))
			j++
		}
		if len(tmp2) != 0 {
			v2, _ = strconv.Atoi(string(tmp2))
			j++
		}

		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}

	return 0
}
// @lc code=end

func main() {
	res := compareVersion("1.01", "1.001")
	fmt.Println(res)
}

