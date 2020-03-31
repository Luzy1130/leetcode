/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 *
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (44.94%)
 * Likes:    4418
 * Dislikes: 161
 * Total Accepted:    581.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * Given a 2d grid map of '1's (land) and '0's (water), count the number of
 * islands. An island is surrounded by water and is formed by connecting
 * adjacent lands horizontally or vertically. You may assume all four edges of
 * the grid are all surrounded by water.
 * 
 * Example 1:
 * 
 * 
 * Input:
 * 11110
 * 11010
 * 11000
 * 00000
 * 
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * 11000
 * 11000
 * 00100
 * 00011
 * 
 * Output: 3
 * 
 */

// @lc code=start
func numIslandsImpl(grid [][]byte, i, j int) {
	if grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'
	// 上
	if i != 0 {
		numIslandsImpl(grid, i-1, j)
	}
	// 下
	if i != len(grid)-1 {
		numIslandsImpl(grid, i+1, j)
	}
	// 左
	if j != 0 {
		numIslandsImpl(grid, i, j-1)
	}
	// 右
	if j != len(grid[i])-1 {
		numIslandsImpl(grid, i, j+1)
	}
}

func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var res int = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			num := grid[i][j]
			if num == '1' {
				numIslandsImpl(grid, i, j)
				res++
			}
		}
	}
	return res
}
// @lc code=end

