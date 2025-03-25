/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N 皇后 II
 *
 * https://leetcode.cn/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (83.13%)
 * Likes:    565
 * Dislikes: 0
 * Total Accepted:    186.7K
 * Total Submissions: 224.7K
 * Testcase Example:  '4'
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4
 * 输出：2
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 9
 *
 *
 *
 *
 */

// @lc code=start
func totalNQueens(n int) int {
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

	var res int
	var bt func(int)
	bt = func(curRow int) {
		if curRow == n {
			// var solution []string
			// for _, row := range board {
			// 	solution = append(solution, strings.Join(row, ""))
			// }
			// res = append(res, solution)
			res++
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

