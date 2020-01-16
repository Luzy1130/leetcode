/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 *
 * https://leetcode.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (16.15%)
 * Likes:    896
 * Dislikes: 4314
 * Total Accepted:    237.7K
 * Total Submissions: 1.5M
 * Testcase Example:  '10\n3'
 *
 * Given two integers dividend and divisor, divide two integers without using
 * multiplication, division and mod operator.
 * 
 * Return the quotient after dividing dividend by divisor.
 * 
 * The integer division should truncate toward zero.
 * 
 * Example 1:
 * 
 * 
 * Input: dividend = 10, divisor = 3
 * Output: 3
 * 
 * Example 2:
 * 
 * 
 * Input: dividend = 7, divisor = -3
 * Output: -2
 * 
 * Note:
 * 
 * 
 * Both dividend and divisor will be 32-bit signed integers.
 * The divisor will never be 0.
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 2^31 − 1 when the
 * division result overflows.
 * 
 * 
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 {
		if divisor == -1 {	// 32位有符号，无法存比math.MaxInt大的数
			return math.MaxInt32
		}
	} 
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	} 
	if divisor == dividend {
		return 1
	}

	neg := false
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		neg = true
	}
	var newDividend uint32
	var newDivisor uint32
	if dividend < 0 {
		newDividend = uint32(-dividend)
	} else {
		newDividend = uint32(dividend)
	}
	if divisor < 0 {
		newDivisor  = uint32(-divisor)
	} else {
		newDivisor = uint32(divisor)
	}

	if newDividend < newDivisor {
		return 0
	}

	var res int = 0
	for newDividend >= newDivisor {
		p := 1
		tmp := newDivisor
		// 以2的倍数减去divisor,比挨个减节省了大量时间
		for newDividend > (tmp << 1) {
			p = p << 1
			tmp = tmp << 1
		}
		res = res + p
		newDividend -= tmp
	}

	if neg {
		res = -res
	}

	return res
}
// @lc code=end

