package go_case

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 *
 * https://leetcode.cn/problems/word-search/description/
 *
 * algorithms
 * Medium (48.50%)
 * Likes:    1997
 * Dislikes: 0
 * Total Accepted:    683.1K
 * Total Submissions: 1.4M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false
 * 。
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "ABCCED"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "SEE"
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "ABCB"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n = board[i].length
 * 1
 * 1
 * board 和 word 仅由大小写英文字母组成
 *
 *
 *
 *
 * 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
 *
 */

// @lc code=start
func exist(board [][]byte, word string) bool {

	wordLen := len(word)
	m := len(board)
	n := len(board[0])
	dir := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	var bt func(int, int, int) bool
	bt = func(x, y, index int) bool {
		if index == wordLen {
			return true
		}
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[index] {
			return false
		}
		temp := board[x][y]
		board[x][y] = 'v'
		for _, d := range dir {
			nextx := x + d[0]
			nexty := y + d[1]

			if bt(nextx, nexty, index+1) {
				return true
			}
		}
		board[x][y] = temp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if bt(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// @lc code=end
