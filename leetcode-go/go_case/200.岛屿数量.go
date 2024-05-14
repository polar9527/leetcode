package go_case

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode.cn/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (60.50%)
 * Likes:    2486
 * Dislikes: 0
 * Total Accepted:    834.2K
 * Total Submissions: 1.4M
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 *
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 *
 * 此外，你可以假设该网格的四条边均被水包围。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1
 * grid[i][j] 的值为 '0' 或 '1'
 *
 *
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	res := 0
	var visited = make([][]bool, rows)
	for i, _ := range visited {
		visited[i] = make([]bool, cols)
	}

	di := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if visited[i][j] || grid[i][j] == '0' {
			return
		}

		visited[i][j] = true
		for k := 0; k < 4; k++ {
			xNext := i + di[k][0]
			yNext := j + di[k][1]

			if xNext >= 0 && xNext < rows && yNext >= 0 && yNext < cols {
				dfs(xNext, yNext)
			}
		}
		return
	}

	var bfs func(i, j int)
	bfs = func(i, j int) {
		var queue [][2]int
		queue = append(queue, [2]int{i, j})
		visited[i][j] = true
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for k := 0; k < 4; k++ {
				xNext := cur[0] + di[k][0]
				yNext := cur[1] + di[k][1]

				if xNext >= 0 && xNext < rows && yNext >= 0 && yNext < cols {
					if !visited[xNext][yNext] && grid[xNext][yNext] == '1' {
						queue = append(queue, [2]int{xNext, yNext})
						visited[xNext][yNext] = true
					}
				}
			}

		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				res++
				// dfs(i, j)
				bfs(i, j)
			}
		}
	}

	return res
}

// @lc code=end
