package go_case

import "strings"

/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 *
 * https://leetcode.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (67.84%)
 * Likes:    12034
 * Dislikes: 271
 * Total Accepted:    693.7K
 * Total Submissions: 1M
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n x n
 * chessboard such that no two queens attack each other.
 *
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 * You may return the answer in any order.
 *
 * Each solution contains a distinct board configuration of the n-queens'
 * placement, where 'Q' and '.' both indicate a queen and an empty space,
 * respectively.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 4
 * Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as
 * shown above
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: [["Q"]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 9
 *
 *
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	isValid := func(row, col int, b [][]string) bool {
		// col
		for r := row; r >= 0; r-- {
			if b[r][col] == "Q" {
				return false
			}
		}
		// diagonal
		for r, l := row-1, col-1; r >= 0 && l >= 0; r, l = r-1, l-1 {
			if b[r][l] == "Q" {
				return false
			}
		}
		for r, l := row-1, col+1; r >= 0 && l < n; r, l = r-1, l+1 {
			if b[r][l] == "Q" {
				return false
			}
		}
		return true
	}

	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	var res [][]string
	var bt func(int)
	bt = func(curRow int) {
		if curRow == n {
			var solution []string
			for _, row := range board {
				solution = append(solution, strings.Join(row, ""))
			}
			res = append(res, solution)
			return
		}

		for col := 0; col < n; col++ {
			if isValid(curRow, col, board) {
				board[curRow][col] = "Q"
				bt(curRow + 1)
				board[curRow][col] = "."
			}
		}
	}

	bt(0)
	return res
}

// @lc code=end
