/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (31.22%)
 * Likes:    1702
 * Dislikes: 487
 * Total Accepted:    306.8K
 * Total Submissions: 944K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given a matrix of m x n elements (m rows, n columns), return all elements of
 * the matrix in spiral order.
 * 
 * Example 1:
 * 
 * 
 * Input:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * Output: [1,2,3,6,9,8,7,4,5]
 * 
 * 
 * Example 2:
 * 
 * Input:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 * 
 */

// @lc code=start
// package main
// import "fmt"

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	maxRow := len(matrix)
	maxCol := len(matrix[0])
	// fmt.Println(maxRow, maxCol)
	// rowLevel和colLevel记录每一圈的第一个元素下标
	rowLevel := 0
	colLevel := 0
    for ; rowLevel < (maxRow+1)/2 && colLevel < (maxCol+1)/2; {
		rowS := rowLevel
		rowE := maxRow - rowLevel
		colS := colLevel
		colE := maxCol - colLevel
		// fmt.Println(rowS, rowE, colS, colE)
		
		i := rowS
		j := colS
		for ;j < colE; j++ {
			res = append(res, matrix[i][j])
		}
		j -= 1 // i=1;j=2
		for i = i+1; i < rowE; i++ {
			res = append(res, matrix[i][j])
		}
		i -= 1	// i=1;j=2
		if rowE-rowS > 1 {
			for j = j-1; j >= colS; j-- {
				res = append(res, matrix[i][j])
			}
		}
		j += 1
		if colE-colS > 1 {
			for i = i-1; i >rowS; i -- {
				res = append(res, matrix[i][j])
			}
		}

		rowLevel++
		colLevel++
	}

	return res
}

// func main() {
// 	// var input [][]int = [][]int {
// 	// 	[]int{ 1, 2, 3, 4 },
// 	// 	[]int{ 5, 6, 7, 8 },
// 	// 	[]int{ 9, 10, 11, 12 },
// 	// }

// 	var input [][]int = [][]int {
// 		[]int{ 1, 2, 3,},
// 		[]int{ 4, 5, 6,},
// 		[]int{ 7, 8, 9,},
// 	}

// 	res := spiralOrder(input)
// 	fmt.Println(res)
// }
// @lc code=end

