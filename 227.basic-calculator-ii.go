/*
 * @lc app=leetcode id=227 lang=golang
 *
 * [227] Basic Calculator II
 *
 * https://leetcode.com/problems/basic-calculator-ii/description/
 *
 * algorithms
 * Medium (35.81%)
 * Likes:    1163
 * Dislikes: 207
 * Total Accepted:    160.8K
 * Total Submissions: 447K
 * Testcase Example:  '"3+2*2"'
 *
 * Implement a basic calculator to evaluate a simple expression string.
 * 
 * The expression string contains only non-negative integers, +, -, *, /
 * operators and empty spaces  . The integer division should truncate toward
 * zero.
 * 
 * Example 1:
 * 
 * 
 * Input: "3+2*2"
 * Output: 7
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: " 3/2 "
 * Output: 1
 * 
 * Example 3:
 * 
 * 
 * Input: " 3+5 / 2 "
 * Output: 5
 * 
 * 
 * Note:
 * 
 * 
 * You may assume that the given expression is always valid.
 * Do not use the eval built-in library function.
 * 
 * 
 */
package main
import (
	"fmt"
)

// 数字：则入”数字栈“
// 操作：比”操作栈“栈顶优先级高，则入”操作栈“； 比”操作栈“栈顶优先级高或相等，则将
// 栈中的操作全部做完，在入”操作栈“

// @lc code=start
func calculateImpl(v1, v2 int, op rune) int {
	switch op {
	case '+':
		return v1 + v2
	case '-':
		return v1 - v2
	case '*':
		return v1 * v2
	case '/':
		return v1 / v2
	default:
		return 0
	}
	return 0
}

func calculate(s string) int {
	var numStack []int
	var opStack []rune

	num := 0
	for _, v := range s {
		if v == ' ' {
			continue
			// if num != 0 {
			// 	numStack = append(numStack, num)
			// 	num = 0
			// } 
		} else if int(v) >= int('0') && int(v) <= int('9') {
				num = num * 10 + int(v) - int('0')
		} else {
			numStack = append(numStack, num)
			num = 0
			if len(opStack) == 0 {
				opStack = append(opStack, v)
				continue
			}

			// op := opStack[len(opStack)-1]
			// v1 := numStack[len(numStack)-2]
			// v2 := numStack[len(numStack)-1]

			switch v {
			case '+':
				fallthrough
			case '-':
				for len(opStack) != 0 {
					op := opStack[len(opStack)-1]
					v1 := numStack[len(numStack)-2]
					v2 := numStack[len(numStack)-1]
					res := calculateImpl(v1, v2, op)
					numStack = numStack[0:len(numStack)-2]
					numStack = append(numStack, res)
					opStack = opStack[0:len(opStack)-1]
				}
				opStack = append(opStack, v)
			case '*':
				fallthrough
			case '/':
				op := opStack[len(opStack)-1]
				v1 := numStack[len(numStack)-2]
				v2 := numStack[len(numStack)-1]
				if op == '+' || op == '-' {
					opStack = append(opStack, v)
				} else {
					res := calculateImpl(v1, v2, op)
					numStack = numStack[0:len(numStack)-2]
					numStack = append(numStack, res)
					opStack = opStack[0:len(opStack)-1]
					opStack = append(opStack, v)
				}
			default:

			}
		}
	}
	numStack = append(numStack, num)

	for len(opStack) != 0 {
		// fmt.Println(string(opStack), numStack)
		op := opStack[len(opStack)-1]
		opStack = opStack[0:len(opStack)-1]
		v1 := numStack[len(numStack)-2]
		v2 := numStack[len(numStack)-1]
		res := calculateImpl(v1, v2, op)
		numStack = numStack[0:len(numStack)-2]
		numStack = append(numStack, res)
	}

	return numStack[0]
}
// @lc code=end

func main() {
	res := calculate("1*2-3/4+5*6-7*8+9/10")
	fmt.Println(res)
}