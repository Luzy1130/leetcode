/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 *
 * https://leetcode.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (47.59%)
 * Likes:    971
 * Dislikes: 85
 * Total Accepted:    315.1K
 * Total Submissions: 640.1K
 * Testcase Example:  '5'
 *
 * Given a non-negative integer numRows, generate the first numRows of Pascal's
 * triangle.
 * 
 * 
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 * 
 * Example:
 * 
 * 
 * Input: 5
 * Output:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 * 
 * 
 */

// @lc code=start
func generate(numRows int) [][]int {
	var res [][]int
	if numRows == 0 {
		return res
	}
	
	res = append(res, []int{1})

	for i := 1; i < numRows; i++ {
		preRow := res[i-1]
		var row []int
		row = append(row, 1) 	// first
		for j := 1; j < i; j++ {
			val := preRow[j-1] + preRow[j]
			row = append(row, val)
		}
		row = append(row, 1) 	// last
		res = append(res, row)
	}

	return res
}
// @lc code=end

