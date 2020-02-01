/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (48.30%)
 * Likes:    2038
 * Dislikes: 48
 * Total Accepted:    302.2K
 * Total Submissions: 597.9K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right which minimizes the sum of all numbers along its path.
 * 
 * Note: You can only move either down or right at any point in time.
 * 
 * Example:
 * 
 * 
 * Input:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * Output: 7
 * Explanation: Because the path 1→3→1→1→1 minimizes the sum.
 * 
 * 
 */

// @lc code=start
// package main
// import (
// 	"fmt"
// 	"math"
// )

func minPathSum(grid [][]int) int {
    if len(grid) == 0 {
		return 0
	}
	n := len(grid)
	m := len(grid[0])

	pathSum := make([][]int, n)
	for i := 0; i < n; i++ {
		pathSum[i] = make([]int, m)
	}

	pathSum[0][0] = grid[0][0]

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				pathSum[i][j] = pathSum[i][j-1] + grid[i][j]
			} else if j == 0 {
				pathSum[i][j] = pathSum[i-1][j] + grid[i][j]
			} else {
				pathSum[i][j] = int(math.Min(float64(pathSum[i][j-1]), float64(pathSum[i-1][j]))) + grid[i][j]
			}
		}
	}
	return pathSum[n-1][m-1]
}

// func main() {
// 	var input [][]int = [][]int{
// 		[]int{1, 3, 1},
// 		[]int{1, 5, 1},
// 		[]int{4, 2, 1},
// 	}
// 	res := minPathSum(input)
// 	fmt.Println(res)
// }
// @lc code=end

