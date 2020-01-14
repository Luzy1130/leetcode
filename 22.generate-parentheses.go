/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (56.66%)
 * Likes:    3852
 * Dislikes: 217
 * Total Accepted:    447.9K
 * Total Submissions: 758.9K
 * Testcase Example:  '3'
 *
 * 
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 * 
 * 
 * 
 * For example, given n = 3, a solution set is:
 * 
 * 
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 * 
 */

// @lc code=start
var Symble []rune = []rune{'(', ')'}

// 回溯法： DFS
func generateParentesisImpl(n int, lNum int) []string {
    var res []string
    if n == 0 && lNum == 0 {
        return res
    }

    for i := 0; i < 2; i++ {
        tmpN := n
        tmpLNum := lNum
        c := Symble[i]
        if c == '(' {
            if n == 0 {
                continue
            }
            tmpLNum++
            tmpN--
        } else {
            if tmpLNum == 0 {
                continue
            }
            tmpLNum--
        }

        tmp := generateParentesisImpl(tmpN, tmpLNum)
        if len(tmp) > 0 {
            for _, v := range tmp {
                res = append(res, string(c)+v)
            }
        } else {
            res = append(res, string(c))
        }
    }
    return res
}

func generateParenthesis(n int) []string {
    res := generateParentesisImpl(n, 0)

    return res
}
// @lc code=end

