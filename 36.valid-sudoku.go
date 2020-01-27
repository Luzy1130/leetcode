/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 *
 * https://leetcode.com/problems/valid-sudoku/description/
 *
 * algorithms
 * Medium (44.45%)
 * Likes:    1218
 * Dislikes: 399
 * Total Accepted:    301.2K
 * Total Submissions: 652.1K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be
 * validated according to the following rules:
 * 
 * 
 * Each row must contain the digits 1-9 without repetition.
 * Each column must contain the digits 1-9 without repetition.
 * Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without
 * repetition.
 * 
 * 
 * 
 * A partially filled sudoku which is valid.
 * 
 * The Sudoku board could be partially filled, where empty cells are filled
 * with the character '.'.
 * 
 * Example 1:
 * 
 * 
 * Input:
 * [
 * ⁠ ["5","3",".",".","7",".",".",".","."],
 * ⁠ ["6",".",".","1","9","5",".",".","."],
 * ⁠ [".","9","8",".",".",".",".","6","."],
 * ⁠ ["8",".",".",".","6",".",".",".","3"],
 * ⁠ ["4",".",".","8",".","3",".",".","1"],
 * ⁠ ["7",".",".",".","2",".",".",".","6"],
 * ⁠ [".","6",".",".",".",".","2","8","."],
 * ⁠ [".",".",".","4","1","9",".",".","5"],
 * ⁠ [".",".",".",".","8",".",".","7","9"]
 * ]
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * [
 * ["8","3",".",".","7",".",".",".","."],
 * ["6",".",".","1","9","5",".",".","."],
 * [".","9","8",".",".",".",".","6","."],
 * ["8",".",".",".","6",".",".",".","3"],
 * ["4",".",".","8",".","3",".",".","1"],
 * ["7",".",".",".","2",".",".",".","6"],
 * [".","6",".",".",".",".","2","8","."],
 * [".",".",".","4","1","9",".",".","5"],
 * [".",".",".",".","8",".",".","7","9"]
 * ]
 * Output: false
 * Explanation: Same as Example 1, except with the 5 in the top left corner
 * being 
 * ⁠   modified to 8. Since there are two 8's in the top left 3x3 sub-box, it
 * is invalid.
 * 
 * 
 * Note:
 * 
 * 
 * A Sudoku board (partially filled) could be valid but is not necessarily
 * solvable.
 * Only the filled cells need to be validated according to the mentioned
 * rules.
 * The given board contain only digits 1-9 and the character '.'.
 * The given board size is always 9x9.
 * 
 * 
 */

// @lc code=start
func isValidSudokuImpl(board [][]byte, row, col int, rowMaps, colMaps, regionMaps []map[int]bool) bool {
	for i := row; i < 9; i++ {
		for j := col; j < 9; j++ {
			if board[i][j] == byte('.') {
				var num int
				for num := 1; num <=9; num++ {
					if _, ok := rowMaps[i][num]; ok {
						continue
					}
					if _, ok := colMaps[j][num]; ok {
						continue
					}
					if _, ok := regionMaps[i/3 * 3 + j/3][num]; ok {
						continue
					}
					rowMaps[i][num] = true
					colMaps[j][num] = true
					regionMaps[i/3 * 3 + j/3][num] = true

					nextRow := i
					nextCol := j
					if nextRow == 8 {
						nextRow++
						nextCol = 0
					} else {
						nextCol++
					}
					res := isValidSudokuImpl(board, nextRow, nextCol, rowMaps, colMaps, regionMaps)
					if res == true {
						return true
					} else {
						delete(rowMaps[i], num)
						delete(colMaps[j], num)
						delete(regionMaps[i/3 * 3 + j/3], num)
					}
				}
				if num == 10 {
					return false
				}
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	var rowMaps []map[int]bool
	var colMaps []map[int]bool
	var regionMaps []map[int]bool

	for i := 0; i < 9; i++ {
		rowMaps = append(rowMaps, make(map[int]bool))
		colMaps = append(colMaps, make(map[int]bool))
		regionMaps = append(regionMaps, make(map[int]bool))
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != byte('.') {
				num, _ := strconv.Atoi(string(board[i][j]))
				if _, ok := rowMaps[i][num]; ok {
					return false
				}
				if _, ok := colMaps[j][num]; ok {
					return false
				}
				if _, ok := regionMaps[i/3 * 3 + j/3][num]; ok {
					return false
				}
				rowMaps[i][num] = true
				colMaps[j][num] = true
				regionMaps[i/3 * 3 + j/3][num] = true
			}
		}
	}

	res := isValidSudokuImpl(board, 0,0, rowMaps, colMaps, regionMaps)
	return res
}
// @lc code=end

