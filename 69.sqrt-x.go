/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (31.97%)
 * Likes:    892
 * Dislikes: 1518
 * Total Accepted:    418.2K
 * Total Submissions: 1.3M
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 * the decimal part is truncated, 2 is returned.
 *
 *
 */
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 || x == 2 {
		return 1
	}

	low := 1
	high := x
	i := (low + high) / 2
	for i >= 1 {
		if i == x/i {
			break
		}
		if i < x/i {
			if (i + 1) > x/(i+1) {
				break
			}
			low = i
		} else {
			high = i
		}
		i = (low + high) / 2
	}

	return i
}

