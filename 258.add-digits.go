/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 *
 * https://leetcode.com/problems/add-digits/description/
 *
 * algorithms
 * Easy (54.65%)
 * Likes:    573
 * Dislikes: 909
 * Total Accepted:    263.3K
 * Total Submissions: 476.2K
 * Testcase Example:  '38'
 *
 * Given a non-negative integer num, repeatedly add all its digits until the
 * result has only one digit.
 * 
 * Example:
 * 
 * 
 * Input: 38
 * Output: 2 
 * Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2. 
 * Since 2 has only one digit, return it.
 * 
 * 
 * Follow up:
 * Could you do it without any loop/recursion in O(1) runtime?
 */

// @lc code=start
// 迭代计算
func method1(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num = num /10
	}
	
	if res >= 10 {
		res = method1(res)
	}
	return res
}

// O(1) runtime
// https://en.wikipedia.org/wiki/Digital_root
func method2(num int) int {
	if num == 0 {
		return 0
	}

	res := num % 9
	if res == 0 {
		res = 9
	}
	return res
}

func addDigits(num int) int {
	// return method1(num)
	return method2(num)
}
// @lc code=end

