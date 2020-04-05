/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 *
 * https://leetcode.com/problems/bitwise-and-of-numbers-range/description/
 *
 * algorithms
 * Medium (37.15%)
 * Likes:    598
 * Dislikes: 82
 * Total Accepted:    97.1K
 * Total Submissions: 261K
 * Testcase Example:  '5\n7'
 *
 * Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND
 * of all numbers in this range, inclusive.
 * 
 * Example 1:
 * 
 * 
 * Input: [5,7]
 * Output: 4
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [0,1]
 * Output: 0
 */
package main
import "fmt"

// @lc code=start
func rangeBitwiseAnd(m int, n int) int {
	i := 0
	// 最小数m和最大数高位相同，相同部分后面的位肯定产品了变化，那么肯定存在0
	for m != n {
		n >>= 1
		m >>= 1
		i++
	}
	res := m << i
	return res
}
// @lc code=end

func main() {
	res := rangeBitwiseAnd(5, 7)
	fmt.Println(res)
}