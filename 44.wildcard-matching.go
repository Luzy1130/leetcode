/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 *
 * https://leetcode.com/problems/wildcard-matching/description/
 *
 * algorithms
 * Hard (24.08%)
 * Likes:    1720
 * Dislikes: 96
 * Total Accepted:    229.7K
 * Total Submissions: 947.9K
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement wildcard pattern
 * matching with support for '?' and '*'.
 * 
 * 
 * '?' Matches any single character.
 * '*' Matches any sequence of characters (including the empty sequence).
 * 
 * 
 * The matching should cover the entire input string (not partial).
 * 
 * Note:
 * 
 * 
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like ? or *.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "*"
 * Output: true
 * Explanation: '*' matches any sequence.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input:
 * s = "cb"
 * p = "?a"
 * Output: false
 * Explanation: '?' matches 'c', but the second letter is 'a', which does not
 * match 'b'.
 * 
 * 
 * Example 4:
 * 
 * 
 * Input:
 * s = "adceb"
 * p = "*a*b"
 * Output: true
 * Explanation: The first '*' matches the empty sequence, while the second '*'
 * matches the substring "dce".
 * 
 * 
 * Example 5:
 * 
 * 
 * Input:
 * s = "acdcb"
 * p = "a*c?b"
 * Output: false
 * 
 * 
 */
package main
import "fmt"

// @lc code=start

// 使用dp, matched[i]记录当p[i]为’*‘时，匹配s的位置j
// 最坏的情况复杂度为O(nm)
// 也可以使用二维数组来记录状态
func dbImpl(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(p) == 0 {
		return false
	}

	if len(s) == 0 {
		for i := 0 ; i < len(p); i++ {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}

	n := len(p)
	m := len(s)
	matched := make([]int, n)
	for i := 0; i < n; i++ {
		matched[0] = 0
	}
	i := 0
	j := 0
	for i < n && j < m {
		fmt.Println(i, j)
		if p[i] == '?' {
			i++
			j++
		} else if p[i] == '*' {
			matched[i] = j
			i++
			// 如果*是最后一个p，则已经匹配完成
			if i == n {
				j = m
			}
		} else {
			if p[i] == s[j] {
				i++
				j++
			} else {
				// 如果不匹配则去找之前*匹配的位置
				// 找最近的’*‘,再之前的其实没有必要再找出来了
				// 因为最近的’*'反正可以匹配后面任何s的子串
				for ; i >= 0; i-- {
					if p[i] == '*' {
						break
					}
				}
				if i < 0 {
					return false
				} else {
					matched[i] += 1
					j = matched[i]
					i++
				}
			}
		}

		if i == n && j < m {
			i--
			for ; i >= 0; i-- {
				if p[i] == '*' {
					break
				}
			}
			if i < 0 {
				return false
			} else {
				matched[i] += 1
				j = matched[i]
				i++
			}
		}

		for j == m && i >= 0 && i < n && p[i] == '*' {
			i++
		}
	}

	if i == n && j == m {
		return true
	}

	return false
}

// 递归求解，超时！
func bt(s string, p string) bool {
	// fmt.Println(s, p)

	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) == 1 {
		if p[0] == '*' {
			return true
		}
		if p[0] == '?' {
			return len(s) == 1
		} 
		return len(s) == 1 && s[0] == p[0]
	}

	if len(s) == 0 {
		for i := 0; i < len(p); i++ {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}

	i := 0
	j := 0
    for i < len(p) && j < len(s) {
		// fmt.Println(i, j)
		switch p[i] {
		case '?':
			i++
			j++
		case '*':
			matchNum := 0
			for matchNum <= len(s)-j {
				res := bt(string(s[j+matchNum:]), string(p[i+1:]))
				if res {
					return true
				} else {
					matchNum++
				}
			}
			return false
		default:
			if p[i] == s[j] {
				i++
				j++
			} else {
				return false
			}
		}
	}

	// fmt.Println(string(s[j:]), string(p[i:]))
	return bt(string(s[j:]), string(p[i:]))
}

func isMatch(s string, p string) bool {
	return dbImpl(s, p)
}
// @lc code=end

func main() {
	res := isMatch("", "*")
	fmt.Println(res)
}