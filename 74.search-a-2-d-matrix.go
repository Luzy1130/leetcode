/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (35.15%)
 * Likes:    1274
 * Dislikes: 138
 * Total Accepted:    281.6K
 * Total Submissions: 786.5K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 * 
 * 
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * Output: false
 * 
 */

// @lc code=start
// package main
// import "fmt"

// 整体有序，使用二分查找
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	r := len(matrix)
	c := len(matrix[0])

	if (target < matrix[0][0]) || (target > matrix[r-1][c-1]) {
		return false
	}

	// 先找出可能在哪一行
	rowS := 0
	rowE := r
	row := -1
	for rowS < rowE {
		mid := (rowS + rowE) / 2
		if target >= matrix[mid][0] && target <= matrix[mid][c-1] {
			row = mid
			break
		}
		if target < matrix[mid][0] {
			rowE = mid
		} else {
			rowS = mid+1
		}
	}
	if row == -1 {
		return false
	}
	// fmt.Println(row)
	// 从某一行中查找
	colS := 0
	colE := c
	for colS < colE {
		mid := (colS + colE) / 2
		// fmt.Println(colS, colE, mid)
		if target == matrix[row][mid] {
			return true
		}
		if target < matrix[row][mid] {
			colE = mid
		} else {
			colS = mid+1
		}
	}
	return false
}

// func main() {
// 	var input [][]int = [][]int{
// 		[]int{ 1, 3, 5, 7 },
// 		[]int{ 10, 11, 16, 20 },
// 		[]int{ 23, 30, 34, 50 },
// 	}

// 	res := searchMatrix(input, 1)
// 	fmt.Println(res)
// }
// @lc code=end

