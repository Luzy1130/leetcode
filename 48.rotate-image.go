/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 *
 * https://leetcode.com/problems/rotate-image/description/
 *
 * algorithms
 * Medium (50.23%)
 * Likes:    2295
 * Dislikes: 188
 * Total Accepted:    333.8K
 * Total Submissions: 631.2K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * You are given an n x n 2D matrix representing an image.
 * 
 * Rotate the image by 90 degrees (clockwise).
 * 
 * Note:
 * 
 * You have to rotate the image in-place, which means you have to modify the
 * input 2D matrix directly. DO NOT allocate another 2D matrix and do the
 * rotation.
 * 
 * Example 1:
 * 
 * 
 * Given input matrix = 
 * [
 * ⁠ [1,2,3],
 * ⁠ [4,5,6],
 * ⁠ [7,8,9]
 * ],
 * 
 * rotate the input matrix in-place such that it becomes:
 * [
 * ⁠ [7,4,1],
 * ⁠ [8,5,2],
 * ⁠ [9,6,3]
 * ]
 * 
 * 
 * Example 2:
 * 
 * 
 * Given input matrix =
 * [
 * ⁠ [ 5, 1, 9,11],
 * ⁠ [ 2, 4, 8,10],
 * ⁠ [13, 3, 6, 7],
 * ⁠ [15,14,12,16]
 * ], 
 * 
 * rotate the input matrix in-place such that it becomes:
 * [
 * ⁠ [15,13, 2, 5],
 * ⁠ [14, 3, 4, 1],
 * ⁠ [12, 6, 8, 9],
 * ⁠ [16, 7,10,11]
 * ]
 * 
 * 
 */

// @lc code=start
// package main
// import (
// 	"fmt"
// )
func rotate(matrix [][]int)  {
	// 从外向内逐圈旋转，n表示从外到内的层次
    for level := 0; level < len(matrix)/2; level++ {
		n := len(matrix)
		for col := level; col < len(matrix)-level-1; col++ {
			// fmt.Println(level, n)
			i := level
			j := col

			replaced := matrix[i][j]
			for cnt := 0; cnt < 4; cnt++ {
				newi := j
				newj := n - 1 - i

				tmp := matrix[newi][newj]
				matrix[newi][newj] = replaced
				replaced = tmp

				i = newi
				j = newj
			}
		}
		// fmt.Println(matrix)
	}
}

// func main() {
// 	var input [][]int = [][]int {
// 		[]int{5,1,9,11},
// 		[]int{2,4,8,10},
// 		[]int{13,3,6,7},
// 		[]int{15,14,12,16},
// 	}

// 	rotate(input)
// 	fmt.Println(input)
// }
// @lc code=end

