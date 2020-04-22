/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 *
 * https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (24.81%)
 * Likes:    763
 * Dislikes: 1154
 * Total Accepted:    167.3K
 * Total Submissions: 670.8K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * You are given a string, s, and a list of words, words, that are all of the
 * same length. Find all starting indices of substring(s) in s that is a
 * concatenation of each word in words exactly once and without any intervening
 * characters.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * ⁠ s = "barfoothefoobarman",
 * ⁠ words = ["foo","bar"]
 * Output: [0,9]
 * Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar"
 * respectively.
 * The output order does not matter, returning [9,0] is fine too.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * ⁠ s = "wordgoodgoodgoodbestword",
 * ⁠ words = ["word","good","best","word"]
 * Output: []
 * 
 * 
 */
 package main
 import "fmt"


// 在暴力求解的基础上进一步优化，我使用两个 map，windows 和 needs，windows存储s的子串，needs存储words，并且使用两个指针 l 和 r,表示一个窗口的左右两边的索引
// 1. 外层循环索引 i 从 0 开始，并且 i < oneWordLen 只用循环一个单词的长度
// 2. 内层循环从 i 开始，l = r = i，依次进行截取 let w = s.slice(i, i + oneWordLen)，判断截取的单词 w 是否存在于存储 needs 中，如果不存在，则直接continue跳过当前循环
// 3. 如果存在则将其添加到 windows 中，如果 windows[w] === needs[w]则将count++，表示一个单词已经准备就绪
// 4. 如果 count === Object.key(needs).length，则表示所有的单词都已经准备就绪，由于有可能会出现多个重复的符合要求的单词，所以此时还需要判断 r - l === words所有单词的长度和
// 6. 如果条件成立则表示当前子字符串符合要求，将子字符串的起始索引l，放入结果集中 ans.push(l)
// 7. 从窗口的左边截取子字符串 let w2 = s.slice(l, l + oneWordLen)，如果 w2 存在于 needs 中，则windows[w2]--，并且如果 windows[w2] < needs[w2]，则count--
// 8. 最后返回结果集ans


// @lc code=start
func force(s string, words []string) []int {
	var res []int
	if len(s) == 0 || len(words) == 0 {
		return res
	}

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	wordLen := len(words[0])
	l := len(words) * wordLen

	for i := 0; i < len(s)-l+1; i++ {
		subStr := string(s[i:(i+l)])
		tmpWordMap := make(map[string]int)
		for j := 0; j < len(subStr)-wordLen+1; j += wordLen {
			tmpWord := string(subStr[j:(j+wordLen)])
			if _, ok := wordMap[tmpWord]; ok {
				tmpWordMap[tmpWord]++
			}
		}

		matched := true
		for v, num := range wordMap {
			if _, ok := tmpWordMap[v]; !ok || num != tmpWordMap[v] {
				matched = false
				break
			}
		}
		if matched {
			res = append(res, i)
		}
	}
	
	return res
}

func findSubstring(s string, words []string) []int {
    return force(s, words)
}
// @lc code=end

func main() {
	res := findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","word"})
	fmt.Println(res)
}