/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 *
 * https://leetcode.com/problems/maximal-square/description/
 *
 * algorithms
 * Medium (35.44%)
 * Likes:    2311
 * Dislikes: 59
 * Total Accepted:    196.1K
 * Total Submissions: 550.7K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * Given a 2D binary matrix filled with 0's and 1's, find the largest square
 * containing only 1's and return its area.
 * 
 * Example:
 * 
 * 
 * Input: 
 * 
 * 1 0 1 0 0
 * 1 0 1 1 1
 * 1 1 1 1 1
 * 1 0 0 1 0
 * 
 * Output: 4
 * 
 */
package main
import (
	"fmt"
	"math"
)
// @lc code=start
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var res int = 0
	tmp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		tmp[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] == '1' {
					tmp[i][j] = 1
				}
			} else {
				if matrix[i][j] == '1' {
					tmp[i][j] = 1
					m := tmp[i-1][j-1]
					if m >= tmp[i-1][j] {
						m = tmp[i-1][j]
					}
					if m >=  tmp[i][j-1] {
						m =  tmp[i][j-1]
					}
					tmp[i][j] += (m + int(math.Sqrt(float64(m))) << 1)
				}
			}
			if res < tmp[i][j] {
				res = tmp[i][j]
			}
		}
	}

	// fmt.Println(tmp)
	return res
}
// @lc code=end

func main() {
	input := [][]byte{
		[]byte{'0','0','0','1'},
		[]byte{'1','1','0','1'},
		[]byte{'1','1','1','1'},
		[]byte{'0','1','1','1'},
		[]byte{'0','1','1','1'},
	}

	res := maximalSquare(input)
	fmt.Println(res)
}