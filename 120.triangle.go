/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 *
 * https://leetcode.com/problems/triangle/description/
 *
 * algorithms
 * Medium (40.49%)
 * Likes:    1598
 * Dislikes: 189
 * Total Accepted:    223.4K
 * Total Submissions: 527.2K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * Given a triangle, find the minimum path sum from top to bottom. Each step
 * you may move to adjacent numbers on the row below.
 * 
 * For example, given the following triangle
 * 
 * 
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 * 
 * 
 * The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
 * 
 * Note:
 * 
 * Bonus point if you are able to do this using only O(n) extra space, where n
 * is the total number of rows in the triangle.
 * 
 */

// @lc code=start
// package main

import (
	// "fmt"
	"math"
)

func db1(triangle [][]int) int {
    if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	sum := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		sum[i] = make([]int, i+1)
	}
	sum[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				sum[i][j] = sum[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				sum[i][j] = sum[i-1][j-1] + triangle[i][j]
			} else {
				// fmt.Println(sum[i-1][j-1], sum[i-1][j])
				if sum[i-1][j-1] < sum[i-1][j] {
					sum[i][j] = sum[i-1][j-1] + triangle[i][j]
				} else {
					sum[i][j] = sum[i-1][j] + triangle[i][j]
				}
			}
		}
	}

	var res int = math.MaxInt32
	for i := 0; i < len(sum); i++ {
		if res > sum[len(sum)-1][i] {
			res = sum[len(sum)-1][i]
		}
	} 
	return res
}

// 反过来，从下向上找
// dp[j] = min(t[i][j]+dp[j], t[i][j]+dp[j+1])
func dp2(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	n := len(triangle)

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = triangle[n-1][i]
	}

	for i := n-2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			sum1 := triangle[i][j] + dp[j]
			sum2 := triangle[i][j] + dp[j+1]
			if sum1 < sum2 {
				dp[j] = sum1
			} else {
				dp[j] = sum2
			}
		}
		// fmt.Println(dp)
	}

	return dp[0]
}

func minimumTotal(triangle [][]int) int {
	// res := dp1(triangle)
	res := dp2(triangle)
	return res
}

// func main() {
// 	var input [][]int = [][]int {
// 		[]int{2},
// 		[]int{3, 4},
// 		[]int{6,5,7},
// 		[]int{4,1,8,3},
// 	}

// 	res := minimumTotal(input)
// 	fmt.Println(res)
// }
// @lc code=end

