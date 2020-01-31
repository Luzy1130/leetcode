/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 *
 * https://leetcode.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (47.97%)
 * Likes:    743
 * Dislikes: 98
 * Total Accepted:    169.8K
 * Total Submissions: 333.9K
 * Testcase Example:  '3'
 *
 * Given a positive integer n, generate a square matrix filled with elements
 * from 1 to n^2 in spiral order.
 * 
 * Example:
 * 
 * 
 * Input: 3
 * Output:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"

func generateMatrix(n int) [][]int {
    if n == 0 {
		return [][]int{}
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	val := 1

	for level := 0;level < (n+1)/2; level++ {
		rowS := level
		rowE := n - level
		colS := level
		colE := n - level

		i := rowS
		j := colS
		
		for ; j < colE; j++ {
			res[i][j] = val
			val++
		}
		j -= 1
		for i = i+1; i < rowE; i++ {
			res[i][j] = val
			val++
		}
		i = i-1

		if n > 1 {
			for j = j-1; j >=colS; j-- {
				res[i][j] = val
				val++
			}
			j = j+1
		}
		if n > 2 {
			for i = i-1; i > rowS; i-- {
				res[i][j] = val
				val++
			}
		}
	}
	return res
}

// func main() {
// 	res := generateMatrix(4)
// 	fmt.Println(res)
// }
// @lc code=end

