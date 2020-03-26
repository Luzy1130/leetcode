/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 *
 * https://leetcode.com/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (19.54%)
 * Likes:    866
 * Dislikes: 2669
 * Total Accepted:    363.7K
 * Total Submissions: 1.9M
 * Testcase Example:  '"the sky is blue"'
 *
 * Given an input string, reverse the string word by word.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "the sky is blue"
 * Output: "blue is sky the"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "  hello world!  "
 * Output: "world! hello"
 * Explanation: Your reversed string should not contain leading or trailing
 * spaces.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "a good   example"
 * Output: "example good a"
 * Explanation: You need to reduce multiple spaces between two words to a
 * single space in the reversed string.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * A word is defined as a sequence of non-space characters.
 * Input string may contain leading or trailing spaces. However, your reversed
 * string should not contain leading or trailing spaces.
 * You need to reduce multiple spaces between two words to a single space in
 * the reversed string.
 * 
 * 
 * 
 * 
 * Follow up:
 * 
 * For C programmers, try to solve it in-place in O(1) extra space.
 */

// @lc code=start
// package main
// import "fmt"

func reverse(s []rune, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func reverseWordsImpl(s []rune) string {
	// 1. 将数组倒序
	reverse(s, 0, len(s)-1)

	// 2. 将每个单词倒序
	start := 0
	end := 0
	for start < len(s) {
		for start < len(s) && s[start] == ' ' {
			start++
		}
		end = start
		for end < len(s) && s[end] != ' ' {
			end++
		}
		if start < len(s) {
			reverse(s, start, end-1)
		}
		start = end
	}

	// 3. 开头和结尾多余的空格
	for ; len(s) > 0 && s[0] == ' '; {
		s = s[1:]
	}
	for ; len(s) > 0 && s[len(s)-1] == ' '; {
		s = s[0:(len(s)-1)]
	}

	// 4. 删除单词之间多余的空格
	var tmp []rune
	spaceCnt := 0
	for _, v := range s {
		if v != ' ' {
			spaceCnt = 0
			tmp = append(tmp, v)
		} else {
			spaceCnt++
			if spaceCnt == 1 {
				tmp = append(tmp, v)
			}
		}
	}

	return string(tmp)

}

func reverseWords(s string) string {
	var tmp []rune
	for _, v := range s {
		tmp = append(tmp, v)
	}
	res := reverseWordsImpl(tmp)
	return res
}

// func main() {
// 	input := "  hello world!  "
// 	res := reverseWords(input)
// 	fmt.Println(res)
// }
// @lc code=end

