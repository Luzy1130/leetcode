/*
 * @lc app=leetcode id=166 lang=golang
 *
 * [166] Fraction to Recurring Decimal
 *
 * https://leetcode.com/problems/fraction-to-recurring-decimal/description/
 *
 * algorithms
 * Medium (20.83%)
 * Likes:    714
 * Dislikes: 1580
 * Total Accepted:    112.7K
 * Total Submissions: 540.1K
 * Testcase Example:  '1\n2'
 *
 * Given two integers representing the numerator and denominator of a fraction,
 * return the fraction in string format.
 * 
 * If the fractional part is repeating, enclose the repeating part in
 * parentheses.
 * 
 * Example 1:
 * 
 * 
 * Input: numerator = 1, denominator = 2
 * Output: "0.5"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: numerator = 2, denominator = 1
 * Output: "2"
 * 
 * Example 3:
 * 
 * 
 * Input: numerator = 2, denominator = 3
 * Output: "0.(6)"
 * 
 * 
 */
package main
import (
	"fmt"
	"strconv"
)
// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	num := numerator
	den := denominator

	if num == 0 {
		return "0"
	}

	var res []rune

	// 判断结果是不是负数
	flag1 := (num < 0)
	flag2 := (den < 0)

	if flag1 {
		num = -num
	}
	if flag2 {
		den = -den
	}

	if flag1 != flag2 {
		res = append(res, '-')
	}

	a := num / den
	b := num % den

	res = append(res, []rune(strconv.Itoa(a))...)

	if b == 0 {
		return string(res)
	}
	res = append(res, '.')
	b *= 10
	modMap := make(map[int]int)
	modMap[b] = len(res)
	for b != 0 {
		a = b / den
		res = append(res, rune(a + int('0')))
		b = b % den
		b *= 10
		if v, ok := modMap[b]; ok {
			tmp := string(res[0:v])+ "(" + string(res[v:]) +")"
			return tmp
		}
		modMap[b] = len(res)
	}

	return string(res)
}
// @lc code=end

func main() {
	res := fractionToDecimal(2, 3)
	fmt.Println(res)
}