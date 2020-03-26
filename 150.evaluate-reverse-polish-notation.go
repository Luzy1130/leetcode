/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 *
 * https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
 *
 * algorithms
 * Medium (34.85%)
 * Likes:    843
 * Dislikes: 431
 * Total Accepted:    207.6K
 * Total Submissions: 594K
 * Testcase Example:  '["2","1","+","3","*"]'
 *
 * Evaluate the value of an arithmetic expression in Reverse Polish Notation.
 * 
 * Valid operators are +, -, *, /. Each operand may be an integer or another
 * expression.
 * 
 * Note:
 * 
 * 
 * Division between two integers should truncate toward zero.
 * The given RPN expression is always valid. That means the expression would
 * always evaluate to a result and there won't be any divide by zero
 * operation.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: ["2", "1", "+", "3", "*"]
 * Output: 9
 * Explanation: ((2 + 1) * 3) = 9
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: ["4", "13", "5", "/", "+"]
 * Output: 6
 * Explanation: (4 + (13 / 5)) = 6
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
 * Output: 22
 * Explanation:tmp* ⁠ ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5tmp = 17 + 5
 * = 22
tmp
 * 
 */

// @lc code=start
func evalRPN(tokens []string) int {
	var res int = 0
	var numStack []int
	
	for _, str := range tokens {
		if str != "+" && str != "-" && str != "*" && str != "/" {
			num, _ := strconv.Atoi(str)
			numStack = append(numStack, num)
		} else {
			var1 := numStack[len(numStack)-2]
			var2 := numStack[len(numStack)-1]
			numStack = numStack[0:len(numStack)-2]
			tmp := 0

			switch str {
			case "+":
				tmp = var1 + var2
			case "-":
				tmp = var1 - var2
			case "*":
				tmp = var1 * var2
			case "/":
				tmp = var1 / var2
			default:
			}
			numStack = append(numStack, tmp)
		}
	}

	res = numStack[0]
	return res
}
// @lc code=end

