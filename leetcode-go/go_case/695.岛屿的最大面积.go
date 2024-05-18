package test

/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 *
 * https://leetcode.cn/problems/max-area-of-island/description/
 *
 * algorithms
 * Medium (68.05%)
 * Likes:    1070
 * Dislikes: 0
 * Total Accepted:    331.7K
 * Total Submissions: 487.4K
 * Testcase Example:  '[[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]'
 *
 * 给你一个大小为 m x n 的二进制矩阵 grid 。
 *
 * 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid
 * 的四个边缘都被 0（代表水）包围着。
 *
 * 岛屿的面积是岛上值为 1 的单元格的数目。
 *
 * 计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid =
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 * 输出：6
 * 解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[0,0,0,0,0,0,0,0]]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 50
 * grid[i][j] 为 0 或 1
 *
 *
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	res := 0
	count := 0
	var visited = make([][]bool, rows)
	for i, _ := range visited {
		visited[i] = make([]bool, cols)
	}

	di := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if visited[i][j] || grid[i][j] == 0 {
			return
		}

		visited[i][j] = true
		count++
		for k := 0; k < 4; k++ {
			xNext := i + di[k][0]
			yNext := j + di[k][1]

			if xNext >= 0 && xNext < rows && yNext >= 0 && yNext < cols {
				dfs(xNext, yNext)
			}
		}
		return
	}

	// var bfs func(i, j int)
	// bfs = func(i, j int) {
	// 	// var queue [][2]int
	// 	queue := [][2]int{{i, j}}
	// 	visited[i][j] = true
	// 	count++
	// 	for len(queue) > 0 {
	// 		cur := queue[0]
	// 		queue = queue[1:]
	// 		for k := 0; k < 4; k++ {
	// 			xNext := cur[0] + di[k][0]
	// 			yNext := cur[1] + di[k][1]

	// 			if xNext >= 0 && xNext < rows && yNext >= 0 && yNext < cols {
	// 				if !visited[xNext][yNext] && grid[xNext][yNext] == 1 {
	// 					queue = append(queue, [2]int{xNext, yNext})
	// 					visited[xNext][yNext] = true
	// 					count++
	// 				}
	// 			}
	// 		}

	// 	}
	// }
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				count = 0
				dfs(i, j)
				// bfs(i, j)
				res = max(res, count)
			}
		}
	}

	return res
}

// @lc code=end
