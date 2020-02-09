/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (32.17%)
 * Likes:    2744
 * Dislikes: 137
 * Total Accepted:    387.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given a 2D board and a word, find if the word exists in the grid.
 * 
 * The word can be constructed from letters of sequentially adjacent cell,
 * where "adjacent" cells are those horizontally or vertically neighboring. The
 * same letter cell may not be used more than once.
 * 
 * Example:
 * 
 * 
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 * 
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"

func existImpl(board [][]byte, path[][]bool, word []byte, posr, posc int) bool {
	i := posr
	j := posc
	if len(word) == 0 {
		return true
	}

	c := word[0]
	for k := 0; k < 4; k++ {
		res := false
		switch k {
		case 0:
			// 向上
			if i != 0 && path[i-1][j] == false && board[i-1][j] == c {
				path[i-1][j] = true
				res = existImpl(board, path, word[1:], i-1, j)
				path[i-1][j] = false
			}
		case 1:
			// 向下
			if i < len(board)-1 && path[i+1][j] == false && board[i+1][j] == c {
				path[i+1][j] = true
				res = existImpl(board, path, word[1:], i+1, j)
				path[i+1][j] = false
			}
		case 2:
			// 向左
			if j != 0 && path[i][j-1] == false && board[i][j-1] == c {
				path[i][j-1] = true
				res = existImpl(board, path, word[1:], i, j-1)
				path[i][j-1] = false
			}
		case 3:
			// 向右
			if j < len(board[0])-1 && path[i][j+1] == false && board[i][j+1] == c {
				path[i][j+1] = true
				res = existImpl(board, path, word[1:], i, j+1)
				path[i][j+1] = false
			}
		}
		if res == true {
			return true
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	r := len(board)
	c := len(board[0])

	path := make([][]bool, r)
	for i := 0; i < r; i++ {
		path[i] = make([]bool, c)
	}

	wordTmp := []byte(word)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == wordTmp[0] {
				path[i][j] = true
				res := existImpl(board, path, wordTmp[1:], i, j)
				if res == true {
					return true
				} else {
					path[i][j] = false
				}
			}
		}
	}
	
	return false
}

// func main() {
// 	var board [][]byte = [][]byte{
// 		[]byte{ 'A','B','C','E' },
// 		[]byte{ 'S','F','C','S' },
// 		[]byte{ 'A','D','E','E' },
// 	}

// 	res := exist(board, "ABFDF")
// 	fmt.Println(res)
// }
// @lc code=end

