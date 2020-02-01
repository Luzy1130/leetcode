/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 *
 * https://leetcode.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (33.64%)
 * Likes:    1243
 * Dislikes: 206
 * Total Accepted:    252.9K
 * Total Submissions: 746K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 * 
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 * 
 * Now consider if some obstacles are added to the grids. How many unique paths
 * would there be?
 * 
 * 
 * 
 * An obstacle and empty space is marked as 1 and 0 respectively in the grid.
 * 
 * Note: m and n will be at most 100.
 * 
 * Example 1:
 * 
 * 
 * Input:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * Output: 2
 * Explanation:
 * There is one obstacle in the middle of the 3x3 grid above.
 * There are two ways to reach the bottom-right corner:
 * 1. Right -> Right -> Down -> Down
 * 2. Down -> Down -> Right -> Right
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if len(obstacleGrid) == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	path := make([][]int, n)
	for i := 0; i < n; i++ {
		path[i] = make([]int, m)
	}

	path[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i == 0 {
				if obstacleGrid[i][j-1] == 0 {
					path[i][j] = path[i][j-1]
				}
			} else if j == 0 {
				if obstacleGrid[i-1][j] == 0 {
					path[i][j] = path[i-1][j]
				}
			} else {
				if obstacleGrid[i][j-1] == 0 && obstacleGrid[i-1][j] == 0 {
					path[i][j] = path[i][j-1] + path[i-1][j]
				} else if obstacleGrid[i][j-1] == 0 {
					path[i][j] = path[i][j-1]
				} else if obstacleGrid[i-1][j] == 0 {
					path[i][j] = path[i-1][j]
				}
			}
		}
	}

	// fmt.Println(path)
	return path[n-1][m-1]
}

// func main() {
// 	var input [][]int = [][]int {
// 		[]int{0,0,0},
// 		[]int{0,1,0},
// 		[]int{0,0,0},
// 	}
// 	res := uniquePathsWithObstacles(input)
// 	fmt.Println(res)
// }
// @lc code=end

