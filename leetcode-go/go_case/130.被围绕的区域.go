package go_case

/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 *
 * https://leetcode.cn/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (46.44%)
 * Likes:    1122
 * Dislikes: 0
 * Total Accepted:    282K
 * Total Submissions: 607.2K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X'
 * 填充。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board =
 * [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
 * 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
 * 解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O'
 * 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
 *
 *
 * 示例 2：
 *
 *
 * 输入：board = [["X"]]
 * 输出：[["X"]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n == board[i].length
 * 1
 * board[i][j] 为 'X' 或 'O'
 *
 *
 *
 *
 */

// @lc code=start
var DIRECTIONS = [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])

	var dfs func(board [][]byte, i, j int)
	dfs = func(board [][]byte, i, j int) {
		board[i][j] = 'A'
		for _, d := range DIRECTIONS {
			x, y := i+d[0], j+d[1]
			if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
				continue
			}
			if board[x][y] == 'O' {
				dfs(board, x, y)
			}
		}
	}

	var bfs func(board [][]byte, i, j int)
	bfs = func(board [][]byte, i, j int) {
		queue := [][]int{{i, j}}
		board[i][j] = 'A'
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, d := range DIRECTIONS {
				x, y := cur[0]+d[0], cur[1]+d[1]
				if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
					continue
				}
				if board[x][y] == 'O' {
					board[x][y] = 'A'
					queue = append(queue, []int{x, y})
				}
			}
		}
	}

	// 列
	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][cols-1] == 'O' {
			dfs(board, i, cols-1)
		}
	}
	// 行
	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}
		if board[rows-1][j] == 'O' {
			dfs(board, rows-1, j)
		}
	}

	for _, r := range board {
		for j, c := range r {
			if c == 'A' {
				r[j] = 'O'
				continue
			}
			if c == 'O' {
				r[j] = 'X'
			}
		}
	}
}

// @lc code=end
