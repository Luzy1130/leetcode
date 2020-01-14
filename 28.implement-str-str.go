/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (32.76%)
 * Likes:    1020
 * Dislikes: 1469
 * Total Accepted:    483.1K
 * Total Submissions: 1.5M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 *
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 *
 * Example 1:
 *
 *
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 *
 *
 * Clarification:
 *
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 *
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 *
 */

func getNext(needle string) []int {
	i := 0
	j := -1
	next := make([]int, len(needle)+1)
	for k := 0; k < len(next); k++ {
		next[k] = 0
	}
	next[0] = -1

	for i < len(needle) {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}

	return next
}

func getPmt(needle string) []int {
	pmt := make([]int, len(needle))

	pmt[0] = 0
	i := 1
	j := 0
	for i < len(needle) {
		if needle[i] == needle[j] {
			j++
			pmt[i] = j
			i++
		} else {
			pmt[i] = 0
			if j > 0 {
				j = pmt[j-1]
			} else {
				j = 0
				i++
			}
		}
	}
	return pmt
}

func strStr(haystack string, needle string) int {
	// 特殊情况特殊处理
	if len(needle) == 0 {
		return 0
	}
	if len(needle) == 1 {
		for i := 0; i < len(haystack); i++ {
			if needle[0] == haystack[i] {
				return i
			}
			return -1
		}
	}
	if len(needle) == len(haystack) && needle == haystack {
		return 0
	}

	// 字符串匹配，建议使用KMP算法
	// next := getNext(needle)

	// i := 0
	// j := 0
	// for i < len(haystack) && j < len(needle) {
	// 	if j == -1 || haystack[i] == needle[j] {
	// 		i++
	// 		j++
	// 	} else {
	// 		j = next[j]
	// 	}
	// }

	pmt := getPmt(needle)
	i := 0
	j := 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j > 0 {
				j = pmt[j-1]
			} else {
				i++
				j = 0
			}
		}
	}
	if j == len(needle) {
		return i - j
	} else {
		return -1
	}
}
