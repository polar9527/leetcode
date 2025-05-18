package go_case

import "strings"

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 *
 * https://leetcode.cn/problems/n-queens/description/
 *
 * algorithms
 * Hard (75.15%)
 * Likes:    2267
 * Dislikes: 0
 * Total Accepted:    508.6K
 * Total Submissions: 676.6K
 * Testcase Example:  '4'
 *
 * 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
 *
 *
 *
 * 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4
 * 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[["Q"]]
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
// func solveNQueens(n int) [][]string {
// 	isValid := func(row, col int, b [][]string) bool {
// 		// col
// 		for r := 0; r < row; r++ {
// 			if b[r][col] == "Q" {
// 				return false
// 			}
// 		}
// 		// diagonal
// 		for r, l := row-1, col-1; r >= 0 && l >= 0; r, l = r-1, l-1 {
// 			if b[r][l] == "Q" {
// 				return false
// 			}
// 		}
// 		for r, l := row-1, col+1; r >= 0 && l < n; r, l = r-1, l+1 {
// 			if b[r][l] == "Q" {
// 				return false
// 			}
// 		}
// 		return true
// 	}

// 	board := make([][]string, n)
// 	for i := 0; i < n; i++ {
// 		board[i] = make([]string, n)
// 		for j := 0; j < n; j++ {
// 			board[i][j] = "."
// 		}
// 	}

// 	var res [][]string
// 	var bt func(int)
// 	bt = func(curRow int) {
// 		if curRow == n {
// 			var solution []string
// 			for _, row := range board {
// 				solution = append(solution, strings.Join(row, ""))
// 			}
// 			res = append(res, solution)
// 		}

// 		for col := 0; col < n; col++ {
// 			if isValid(curRow, col, board) {
// 				board[curRow][col] = "Q"
// 				bt(curRow + 1)
// 				board[curRow][col] = "."
// 			}
// 		}
// 	}

// 	bt(0)
// 	return res
// }

// func solveNQueens(n int) [][]string {

// 	var res [][]string
// 	// 第i行的皇后 放在第  queens[i] 列
// 	queens := make([]int, n)
// 	// 第 i 列, 是否有 皇后
// 	col := make([]bool, n)

// 	// 判断斜线上是否有效
// 	valid := func(r int, c int) bool {
// 		for ri := 0; ri < r; ri++ {
// 			// r 之前的 行中，放置的皇后所在的列
// 			cq := queens[ri]
// 			// 两个方向 上 时候存在皇后
// 			if ri+cq == r+c || r-c == ri-cq {
// 				return false
// 			}
// 		}
// 		return true
// 	}

// 	var bt func(int)
// 	bt = func(r int) {
// 		if r == n {
// 			board := make([]string, n)
// 			for i, c := range queens {
// 				// 第i行的皇后 在 第c列, c 在范围[0, n-1] 内
// 				// 在第 c 列 意味着前面 有 c 个 “.”
// 				// 后面 有  n-c-1 个 “.”
// 				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-c-1)
// 			}
// 			res = append(res, board)
// 			return
// 		}
// 		// 遍历当前行 r 的 每一列，

// 		// 用 valid 确认 两个斜线时候上时候有其他 皇后
// 		for c, ok := range col {

// 			if !ok && valid(r, c) {
// 				// 从 col 中选出 剩余 没有被皇后占据的 列， c, ok := range col
// 				queens[r] = c
// 				col[c] = true
// 				bt(r + 1)
// 				col[c] = false
// 			}
// 		}
// 	}
// 	bt(0)
// 	return res
// }

func solveNQueens(n int) [][]string {

	var res [][]string
	// 第i行的皇后 放在第  queens[i] 列
	queens := make([]int, n)
	// 第 i 列, 是否有 皇后
	col := make([]bool, n)

	// 判断斜线上是否有效, 优化
	// diagonal 数组大小均为 2*n-1 的原因是
	// r 行序号的范围是 [0,n-1]
	// c 列序号的范围是 [0,n-1]
	// r+c 的范围数值范围是 [0,2n-2],即 长度为 2n-1
	// r-c 的范围数值范围是 [-(n-1),n-1],即 长度为 2n-1
	// [-(n-1),n-1] 无法用数组序号表示，用 n-1 偏移调整一下
	// rc = r-c + n-1

	diagonal45 := make([]bool, 2*n-1)
	diagonal134 := make([]bool, 2*n-1)

	var bt func(int)
	bt = func(r int) {
		if r == n {
			board := make([]string, n)
			for i, c := range queens {
				// 第i行的皇后 在 第c列, c 在范围[0, n-1] 内
				// 在第 c 列 意味着前面 有 c 个 “.”
				// 后面 有  n-c-1 个 “.”
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-c-1)
			}
			res = append(res, board)
			return
		}
		// 遍历当前行 r 的 每一列，

		// 用 valid 确认 两个斜线时候上时候有其他 皇后
		for c, ok := range col {
			rc := r - c + n - 1
			if !ok && !diagonal45[r+c] && !diagonal134[rc] {
				// 从 col 中选出 剩余 没有被皇后占据的 列， c, ok := range col
				queens[r] = c
				col[c], diagonal45[r+c], diagonal134[rc] = true, true, true
				bt(r + 1)
				col[c], diagonal45[r+c], diagonal134[rc] = false, false, false
			}
		}
	}
	bt(0)
	return res
}

// @lc code=end
