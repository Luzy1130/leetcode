/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 *
 * https://leetcode.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (25.50%)
 * Likes:    1207
 * Dislikes: 552
 * Total Accepted:    189.4K
 * Total Submissions: 741.5K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * Given a 2D board containing 'X' and 'O' (the letter O), capture all regions
 * surrounded by 'X'.
 * 
 * A region is captured by flipping all 'O's into 'X's in that surrounded
 * region.
 * 
 * Example:
 * 
 * 
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 * 
 * 
 * After running your function, the board should be:
 * 
 * 
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 * 
 * 
 * Explanation:
 * 
 * Surrounded regions shouldn’t be on the border, which means that any 'O' on
 * the border of the board are not flipped to 'X'. Any 'O' that is not on the
 * border and it is not connected to an 'O' on the border will be flipped to
 * 'X'. Two cells are connected if they are adjacent cells connected
 * horizontally or vertically.
 * 
 */

// @lc code=start
func dfs(board [][]byte, visited [][]bool, i, j int) {
	m := len(board)
	n := len(board[0])

	if i >= 0 && i < m && j >= 0 && j < n && visited[i][j] == false && board[i][j] == 'O' {
		visited[i][j] = true
		board[i][j] = '#'
		// 上下左右遍历
		dfs(board, visited, i-1, j)
		dfs(board, visited, i+1, j)
		dfs(board, visited, i, j-1)
		dfs(board, visited, i, j+1)
	}
}

func solve(board [][]byte)  {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	// 从四边向里面传染进去,没有被传染的就是被X包围的
	for i := 0; i < m; i++ {
		dfs(board, visited, i, 0)
		dfs(board, visited, i, n-1)
	}
	for j := 1; j < n-1; j++ {
		dfs(board, visited, 0, j)
		dfs(board, visited, m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
	
}
// @lc code=end

