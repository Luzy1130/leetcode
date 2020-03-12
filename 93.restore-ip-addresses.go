/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 *
 * https://leetcode.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (32.27%)
 * Likes:    982
 * Dislikes: 443
 * Total Accepted:    172.2K
 * Total Submissions: 509K
 * Testcase Example:  '"25525511135"'
 *
 * Given a string containing only digits, restore it by returning all possible
 * valid IP address combinations.
 * 
 * Example:
 * 
 * 
 * Input: "25525511135"
 * Output: ["255.255.11.135", "255.255.111.35"]
 * 
 * 
 */

// @lc code=start
// package main
import (
	// "fmt"
	"strconv"
)

func restoreIpAddressesImpl(s string, level int) [][]string {
	var res [][]string
	// fmt.Println(s)
	if len(s) == 0 {
		return res
	}
	if level == 1 {
		n, _ := strconv.Atoi(s)
		if n <= 255 && s[0] != '0' {
			res = append(res, []string{s})
		} else  if s[0] == '0' && len(s) == 1 {
			res = append(res, []string{s})
		}
		return res
	}

	for i := 0; i < len(s) && i < 3; i++ {
		n, _ := strconv.Atoi(s[0:i+1])
		// fmt.Println(n)
		if n > 255 {
			break
		}
		suffix := restoreIpAddressesImpl(s[i+1:], level-1)
		for _, v :=  range suffix {
			if len(v) != level - 1 {
				continue
			}
			tmp := []string{string(s[0:i+1])}
			tmp = append(tmp, v...)
			res = append(res, tmp)
		}
		if s[0] == '0' {
			break
		}
	}
	return res
}

func restoreIpAddresses(s string) []string {
	var res []string
	hash := make(map[string]bool)
	arrays := restoreIpAddressesImpl(s, 4)
	// fmt.Println(arrays)
	for _, array := range arrays {
		var tmp string
		for j, v := range array {
			if j == 0 {
				tmp = v
			} else {
				tmp = tmp + "." + v
			}
		}
		if _, ok := hash[tmp]; ok {
			continue
		}
		hash[tmp] = true
		res = append(res, tmp)
	} 
	return res
}

// func main() {
// 	res := restoreIpAddresses("25525511135")
// 	fmt.Println(res)
// }
// @lc code=end

