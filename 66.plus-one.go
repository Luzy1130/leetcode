/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (41.61%)
 * Likes:    988
 * Dislikes: 1753
 * Total Accepted:    444.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of digits representing a non-negative integer, plus
 * one to the integer.
 *
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contain a single digit.
 *
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 *
 *
 * Example 2:
 *
 *
 * Input: [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 *
 */
func plusOne(digits []int) []int {
	carry := 1
	length := len(digits)
	res := make([]int, length+1)

	i := length - 1
	j := length
	for i >= 0 {
		temp := digits[i] + carry
		res[j] = temp % 10
		carry = temp / 10
		i--
		j--
	}

	if carry > 0 {
		res[j] = carry
	} else {
		res = res[1:]
	}
	return res
}

