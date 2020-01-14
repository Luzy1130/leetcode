/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (33.02%)
 * Likes:    1316
 * Dislikes: 3925
 * Total Accepted:    393.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 * 
 * 
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 * 
 * 
 * And then read line by line: "PAHNAPLSIIGYIR"
 * 
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 * 
 * 
 * string convert(string s, int numRows);
 * 
 * Example 1:
 * 
 * 
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 * 
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 * 
 */

// @lc code=start
func convert(s string, numRows int) string {
    var tmpRes []byte
	if numRows == 1 {
		return s
	}

	for i := 0; i < numRows; i++ {
		intvl1 := 2 * (numRows - (i + 1))
		intvl2 := 2 * i
		intervals := []int{intvl1, intvl2}

		k := 0
		for j := i; j < len(s); {
			tmpRes = append(tmpRes, s[j])
			for intervals[k] == 0 {
				k ^= 1
			}
			j += intervals[k]
			k ^= 1
		}
	}

	res := string(tmpRes)
	return res
}
// @lc code=end

