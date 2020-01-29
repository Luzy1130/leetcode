/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (48.84%)
 * Likes:    2480
 * Dislikes: 146
 * Total Accepted:    475.1K
 * Total Submissions: 911.6K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings, group anagrams together.
 * 
 * Example:
 * 
 * 
 * Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * Output:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 * 
 * Note:
 * 
 * 
 * All inputs will be in lowercase.
 * The order of your output does not matter.
 * 
 * 
 */

// @lc code=start
// package main
// import (
// 	"sort"
// 	"fmt"
// )

type RuneSlice []rune
func (s RuneSlice) Len() int { return len(s) }
func (s RuneSlice) Swap(i, j int) { s[i],s[j] = s[j], s[i] }
func (s RuneSlice) Less(i, j int) bool {return s[i] < s[j] }

func sortedString(str string) string {
	var runes RuneSlice
	for _, v := range str {
		runes = append(runes, v)
	}

	sort.Sort(runes)
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	strMap := make(map[string][]string)
	
	for _, v := range strs {
		sortedStr := sortedString(v)
		strMap[sortedStr] = append(strMap[sortedStr], v)
	}

	for _, str := range strMap {
		res = append(res, str)
	}
	return res
}

// func main() {
// 	var input []string = []string{
// 		"eat", "tea", "tan", "ate", "nat", "bat",
// 	}

// 	res := groupAnagrams(input)
// 	fmt.Println(res)
// }
// @lc code=end

