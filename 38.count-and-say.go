/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 *
 * https://leetcode.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (41.53%)
 * Likes:    899
 * Dislikes: 7042
 * Total Accepted:    315.7K
 * Total Submissions: 755.9K
 * Testcase Example:  '1'
 *
 * The count-and-say sequence is the sequence of integers with the first five
 * terms as following:
 *
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 *
 *
 * 1 is read off as "one 1" or 11.
 * 11 is read off as "two 1s" or 21.
 * 21 is read off as "one 2, then one 1" or 1211.
 *
 * Given an integer n where 1 ≤ n ≤ 30, generate the n^th term of the
 * count-and-say sequence.
 *
 * Note: Each term of the sequence of integers will be represented as a
 * string.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: "1"
 *
 *
 * Example 2:
 *
 *
 * Input: 4
 * Output: "1211"
 *
 */
func countAndSay(n int) string {
	start := "1"
	if n == 1 {
		return start
	}

	tmp := start
	for i := 2; i <= n; i++ {
		c := tmp[0]
		var cnt rune = 1
		var tmpslice []rune
		for j := 1; j < len(tmp); j++ {
			if c != tmp[j] {
				tmpslice = append(tmpslice, cnt-0+'0')
				tmpslice = append(tmpslice, rune(c))
				cnt = 1
				c = tmp[j]
			} else {
				cnt++
			}
		}
		tmpslice = append(tmpslice, cnt-0+'0')
		tmpslice = append(tmpslice, rune(c))
		tmp = string(tmpslice)
	}

	return tmp
}

