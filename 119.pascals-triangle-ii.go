/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 *
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (44.92%)
 * Likes:    596
 * Dislikes: 183
 * Total Accepted:    240.8K
 * Total Submissions: 520.4K
 * Testcase Example:  '3'
 *
 * Given a non-negative index k where k ≤ 33, return the k^th index row of the
 * Pascal's triangle.
 * 
 * Note that the row index starts from 0.
 * 
 * 
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 * 
 * Example:
 * 
 * 
 * Input: 3
 * Output: [1,3,3,1]
 * 
 * 
 * Follow up:
 * 
 * Could you optimize your algorithm to use only O(k) extra space?
 * 
 */

// @lc code=start
func getRow(rowIndex int) []int {
	var res []int
	if rowIndex < 0 {
		return res
	}
	
	res = append(res, 1)

	for i := 1; i < rowIndex+1; i++ {
		var row []int
		row = append(row, 1) 	// first
		for j := 1; j < i; j++ {
			val := res[j-1] + res[j]
			row = append(row, val)
		}
		row = append(row, 1) 	// last
		res = row
	}

	return res
}
// @lc code=end

