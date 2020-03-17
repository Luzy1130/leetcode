/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 *
 * https://leetcode.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (27.77%)
 * Likes:    2511
 * Dislikes: 1019
 * Total Accepted:    368.9K
 * Total Submissions: 1.3M
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * Given two words (beginWord and endWord), and a dictionary's word list, find
 * the length of shortest transformation sequence from beginWord to endWord,
 * such that:
 * 
 * 
 * Only one letter can be changed at a time.
 * Each transformed word must exist in the word list. Note that beginWord is
 * not a transformed word.
 * 
 * 
 * Note:
 * 
 * 
 * Return 0 if there is no such transformation sequence.
 * All words have the same length.
 * All words contain only lowercase alphabetic characters.
 * You may assume no duplicates in the word list.
 * You may assume beginWord and endWord are non-empty and are not the same.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 * 
 * Output: 5
 * 
 * Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" ->
 * "dog" -> "cog",
 * return its length 5.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 * 
 * Output: 0
 * 
 * Explanation: The endWord "cog" is not in wordList, therefore no possible
 * transformation.
 * 
 * 
 * 
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}
	hash := make(map[string]bool)
	visited := make(map[string]bool)

	for _, word := range wordList {
		hash[word] = true
	}

	res := 2
	length := len(beginWord)
	visited[beginWord] = true
	var curWords []string = []string{beginWord}
	for len(curWords) != 0 {
		// fmt.Println(curWords)
		var nextWords []string
		for _, word := range curWords {
			// 变更word中的第i个字符
			for i := 0; i < length; i++ {
				var tmp []rune
				for _, v := range word {
					tmp = append(tmp, v)
				}
				for c := 'a'; c <= 'z'; {
					tmp[i] = c
					tmpStr := string(tmp)
					// fmt.Println(tmpStr)
					c = rune(int(c)+1)
					// 已经转变为endStr
					if _, ok := hash[tmpStr]; !ok {
						continue
					}
					if _, ok := visited[tmpStr]; ok {
						continue
					}
					if tmpStr == endWord {
						return res
					}
					visited[tmpStr] = true
					nextWords = append(nextWords, tmpStr)
				}
			}
		}
		curWords = nextWords
		res++
	}

	return 0
}

// func main() {
// 	beginWord := "hit"
// 	endWord := "cog"
// 	wordList := []string{"hot","dot","dog","lot","log", "cog"}

// 	res := ladderLength(beginWord, endWord, wordList)
// 	fmt.Println(res)
// }
// @lc code=end

