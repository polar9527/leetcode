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
	board := make([][]string, n)

	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	var res [][]string
	var backtracing func(int)
	backtracing = func(curRow int) {
		if curRow == n {
			tmp := []string{}
			for _, v := range board {
				tmp = append(tmp, strings.Join(v, ""))
			}
			res = append(res, tmp)
			return
		}

		for col := 0; col < n; col++ {
			if isValidPlacement(curRow, col, board) {
				// place
				board[curRow][col] = "Q"
				backtracing(curRow + 1)
				board[curRow][col] = "."
			}
		}
	}
	backtracing(0)
	return res
}

func isValidPlacement(row, col int, board [][]string) bool {
	for r := 0; r < len(board); r++ {
		// 同列
		if board[r][col] == "Q" {
			return false
		}
	}
	// 同一对角线
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == "Q" {
			return false
		}
	}
	for r, c := row-1, col+1; r >= 0 && c < len(board); r, c = r-1, c+1 {
		if board[r][c] == "Q" {
			return false
		}
	}
	return true
}

// @lc code=end
