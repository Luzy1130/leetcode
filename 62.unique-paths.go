/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 *
 * https://leetcode.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (48.88%)
 * Likes:    2326
 * Dislikes: 168
 * Total Accepted:    384.3K
 * Total Submissions: 754.1K
 * Testcase Example:  '3\n2'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 * 
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 * 
 * How many possible unique paths are there?
 * 
 * 
 * Above is a 7 x 3 grid. How many possible unique paths are there?
 * 
 * Note: m and n will be at most 100.
 * 
 * Example 1:
 * 
 * 
 * Input: m = 3, n = 2
 * Output: 3
 * Explanation:
 * From the top-left corner, there are a total of 3 ways to reach the
 * bottom-right corner:
 * 1. Right -> Right -> Down
 * 2. Right -> Down -> Right
 * 3. Down -> Right -> Right
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: m = 7, n = 3
 * Output: 28
 * 
 */

// @lc code=start
// package main
// import "fmt"

// 
// PATH[i][j] = PATH[i-1][j]+ PATH[i][j-1]
func uniquePaths(m int, n int) int {
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
			if i == 0 {
				path[i][j] = path[i][j-1]
			} else if j == 0 {
				path[i][j] = path[i-1][j]
			} else {
				path[i][j] = path[i-1][j] + path[i][j-1]
			}
		}
	}

	return path[n-1][m-1]
}

// func main() {
// 	res := uniquePaths(7, 3)
// 	fmt.Println(res)
// }
// @lc code=end

