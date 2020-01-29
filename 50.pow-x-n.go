/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (28.46%)
 * Likes:    1140
 * Dislikes: 2592
 * Total Accepted:    399.5K
 * Total Submissions: 1.4M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 * 
 * Example 1:
 * 
 * 
 * Input: 2.00000, 10
 * Output: 1024.00000
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 2.10000, 3
 * Output: 9.26100
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 * 
 * 
 * Note:
 * 
 * 
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"

func myPow(x float64, n int) float64 {
    if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	var res float64 = x
	i := 2
	for ; i <= n; i=i*2 {
		res *= res
	}
	i = i / 2
	if i < n {
		res *= myPow(x, n-i)
	}

	return res
}

// func main() {
// 	res := myPow(2, 12)
// 	fmt.Println(res)
// }
// @lc code=end

