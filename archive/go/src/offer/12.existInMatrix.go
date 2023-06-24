/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (33.67%)
 * Likes:    2949
 * Dislikes: 150
 * Total Accepted:    408.2K
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
 *
 * Constraints:
 *
 *
 * board and word consists only of lowercase and uppercase English letters.
 * 1 <= board.length <= 200
 * 1 <= board[i].length <= 200
 * 1 <= word.length <= 10^3
 *
 *
 */
package main

import (
	"fmt"
)

// @lc code=start
func exist(board [][]byte, word string) bool {
	pathLength := 0

	rows := len(board)
	if rows == 0 {
		return false
	}
	cols := len(board[0])

	visited := make([]bool, rows*cols)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if hasPath(&board, r, c, rows, cols, &visited, &word, &pathLength) {
				return true
			}
		}
	}
	return false

}

func hasPath(board *[][]byte, row, col, rows, cols int, visited *[]bool, word *string, pathLength *int) bool {
	if *pathLength == len(*word) {
		return true
	}

	var hasPathFlag = false
	if row >= 0 && row < rows && col >= 0 && col < cols && !(*visited)[row*cols+col] && (*board)[row][col] == (*word)[*pathLength] {
		*pathLength++
		(*visited)[row*cols+col] = true
		hasPathFlag = hasPath(board, row+1, col, rows, cols, visited, word, pathLength) ||
			hasPath(board, row, col+1, rows, cols, visited, word, pathLength) ||
			hasPath(board, row-1, col, rows, cols, visited, word, pathLength) ||
			hasPath(board, row, col-1, rows, cols, visited, word, pathLength)
		if !hasPathFlag {
			*pathLength--
			(*visited)[row*cols+col] = false
		}

	}

	return hasPathFlag
}

// @lc code=end

func main() {
	var maxrix = [][]byte{
		{byte('A'), byte('B'), byte('C'), byte('E')},
		{byte('S'), byte('F'), byte('C'), byte('S')},
		{byte('A'), byte('D'), byte('E'), byte('E')},
	}
	ret := exist(maxrix, "ABCCED")
	fmt.Println(ret)
}
