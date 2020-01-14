/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 *
 * https://leetcode.com/problems/excel-sheet-column-number/description/
 *
 * algorithms
 * Easy (52.17%)
 * Likes:    709
 * Dislikes: 131
 * Total Accepted:    252.8K
 * Total Submissions: 477.8K
 * Testcase Example:  '"A"'
 *
 * Given a column title as appear in an Excel sheet, return its corresponding
 * column number.
 * 
 * For example:
 * 
 * 
 * ⁠   A -> 1
 * ⁠   B -> 2
 * ⁠   C -> 3
 * ⁠   ...
 * ⁠   Z -> 26
 * ⁠   AA -> 27
 * ⁠   AB -> 28 
 * ⁠   ...
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "A"
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "AB"
 * Output: 28
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "ZY"
 * Output: 701
 * 
 */

// @lc code=start
// 26进制转10进制
func titleToNumber(s string) int {
	var res int = 0
	base := 26

	carry := 1
    for i := len(s)-1; i >=0; i-- {
		num := int(byte(s[i]) - byte('A')) + 1

		res += (carry * num)
		carry = carry * base
	}
	return res
}
// @lc code=end

