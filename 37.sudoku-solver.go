/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 *
 * https://leetcode.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (41.12%)
 * Likes:    1511
 * Dislikes: 87
 * Total Accepted:    172.4K
 * Total Submissions: 413.5K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Write a program to solve a Sudoku puzzle by filling the empty cells.
 * 
 * A sudoku solution must satisfy all of the following rules:
 * 
 * 
 * Each of the digits 1-9 must occur exactly once in each row.
 * Each of the digits 1-9 must occur exactly once in each column.
 * Each of the the digits 1-9 must occur exactly once in each of the 9 3x3
 * sub-boxes of the grid.
 * 
 * 
 * Empty cells are indicated by the character '.'.
 * 
 * 
 * A sudoku puzzle...
 * 
 * 
 * ...and its solution numbers marked in red.
 * 
 * Note:
 * 
 * 
 * The given board contain only digits 1-9 and the character '.'.
 * You may assume that the given Sudoku puzzle will have a single unique
 * solution.
 * The given board size is always 9x9.
 * 
 * 
 */

// @lc code=start
func solveSudokuImpl(board[][]byte, row, col, reg []map[int]bool, begini, beginj int) bool {
	// fmt.Println(begini, beginj, board)
	for i := begini; i < 9; i++ {
		var j int
		if i != begini {
			j = 0
		} else {
			j = beginj
		}
		for ; j < 9; j++ {
			if board[i][j] == '.' {
				for k := 1; k <= 9; k++ {
					if _, ok := row[i][k]; ok {
						continue
					}
					if _, ok := col[j][k]; ok {
						continue
					}
					if _, ok := reg[(i/3) * 3 + (j/3)][k]; ok {
						continue
					}
					row[i][k] = true
					col[j][k] = true
					reg[(i/3)*3 + (j/3)][k] = true
					board[i][j] = byte(int('0')+k)
					res := solveSudokuImpl(board, row, col, reg, i, j+1)
					if res == false {
						delete(row[i], k)
						delete(col[j], k)
						delete(reg[(i/3)*3 + (j/3)], k)
						// board[i][j] = '.'
					} else {
						return true
					}
				}
				board[i][j] = '.'
				return false
			}
		}
	}
	return true
}

func solveSudoku(board [][]byte)  {
	row := make([]map[int]bool, 9)
	col := make([]map[int]bool, 9)
	reg := make([]map[int]bool, 9)

	for i := 0; i < 9; i++ {
		row[i] = make(map[int]bool)
		col[i] = make(map[int]bool)
		reg[i] = make(map[int]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				row[i][num] = true
				col[j][num] = true
				regIndex := (i/3) * 3 + (j/3)
				reg[regIndex][num] = true
			}
		}
	}
	solveSudokuImpl(board, row, col, reg, 0, 0)

}
// @lc code=end

