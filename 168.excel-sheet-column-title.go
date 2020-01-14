/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 *
 * https://leetcode.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (29.45%)
 * Likes:    911
 * Dislikes: 192
 * Total Accepted:    194.2K
 * Total Submissions: 648.5K
 * Testcase Example:  '1'
 *
 * Given a positive integer, return its corresponding column title as appear in
 * an Excel sheet.
 * 
 * For example:
 * 
 * 
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB 
 * ⁠   ...
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: 1
 * Output: "A"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 28
 * Output: "AB"
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 701
 * Output: "ZY"
 * 
 */

// @lc code=start

// 相当于是10进制转为26进制
func convertToTitle(n int) string {
	var stack []byte
	base := 26

	for n > 0 {
		tmp := (n-1) % base
		c := byte('A') + byte(tmp)
		stack = append(stack, c)
		n = (n-1) / base
	}
	
	var nums []byte
	for i := len(stack)-1; i >= 0; i-- {
		nums = append(nums, stack[i])
	}
	res := string(nums)
	return res
}
// @lc code=end

