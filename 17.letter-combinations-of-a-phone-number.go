/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (42.74%)
 * Likes:    2961
 * Dislikes: 356
 * Total Accepted:    501.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent.
 * 
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Input: "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 * 
 * 
 * Note:
 * 
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 * 
 */

// @lc code=start
var digitMap map[rune][]rune = map[rune][]rune {
	'2': []rune{'a', 'b', 'c'},
	'3': []rune{'d', 'e', 'f'},
	'4': []rune{'g', 'h', 'i'},
	'5': []rune{'j', 'k', 'l'},
	'6': []rune{'m', 'n', 'o'},
	'7': []rune{'p', 'q', 'r', 's'},
	'8': []rune{'t', 'u', 'v'},
	'9': []rune{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	
	d := rune(digits[0])
	if _, ok := digitMap[d]; ok {
		letters := digitMap[d]
		for _, c := range letters {
			var tmp []string
			pre := letterCombinations(string(digits[1:]))
			if len(pre) > 0 {
				for j := 0; j < len(pre); j++ {
					tmp = append(tmp, string(c) + pre[j])
				}
			} else {
				tmp = append(tmp, string(c))
			}
			res = append(res, tmp...)
		}
	}

	return res
}
// @lc code=end

