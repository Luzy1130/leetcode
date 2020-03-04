/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 *
 * https://leetcode.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (42.09%)
 * Likes:    532
 * Dislikes: 953
 * Total Accepted:    192.3K
 * Total Submissions: 447.5K
 * Testcase Example:  '"hello"'
 *
 * Write a function that takes a string as input and reverse only the vowels of
 * a string.
 * 
 * Example 1:
 * 
 * 
 * Input: "hello"
 * Output: "holle"
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "leetcode"
 * Output: "leotcede"
 * 
 * 
 * Note:
 * The vowels does not include the letter "y".
 * 
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"

var vowelset map[byte]bool = map[byte]bool {
	'a': true,
	'A': true,
	'e': true,
	'E': true,
	'i': true,
	'I': true,
	'o': true,
	'O': true,
	'u': true,
	'U': true,
}

func isVowel(c byte) bool {
	if _, ok := vowelset[c]; ok {
		return true
	}
	return false
}

func reverseVowels(s string) string {
	i := 0
	j := len(s) - 1
	ss := []byte(s)
	
	for i < j {
		if isVowel(ss[i]) && isVowel(ss[j]) {
			ss[i], ss[j] = ss[j], ss[i]
			i++
			j--
		} else if isVowel(ss[i]) && !isVowel(ss[j]) {
			j--
		} else if !isVowel(ss[i]) && isVowel(ss[j]) {
			i++
		} else {
			i++
			j--
		}
	}
	return string(ss)
}

// func main() {
// 	res := reverseVowels("leetcode")
// 	fmt.Println(res)
// }
// @lc code=end

