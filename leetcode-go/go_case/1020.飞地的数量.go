package go_case

/*
 * @lc app=leetcode.cn id=1020 lang=golang
 *
 * [1020] 飞地的数量
 *
 * https://leetcode.cn/problems/number-of-enclaves/description/
 *
 * algorithms
 * Medium (61.43%)
 * Likes:    272
 * Dislikes: 0
 * Total Accepted:    81.9K
 * Total Submissions: 133.4K
 * Testcase Example:  '[[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]'
 *
 * 给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
 *
 * 一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
 *
 * 返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
 * 输出：3
 * 解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
 * 输出：0
 * 解释：所有 1 都在边界上或可以到达边界。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 500
 * grid[i][j] 的值为 0 或 1
 *
 *
 */

// @lc code=start
func numEnclaves(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	var visited = make([][]bool, rows)
	for i, _ := range visited {
		visited[i] = make([]bool, cols)
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var count = 0
	// var dfs func(int, int)
	// dfs = func(i, j int) {
	// 	if grid[i][j] == 0 || visited[i][j] {
	// 		return
	// 	}
	// 	count++
	// 	visited[i][j] = true
	// 	for d := range dirs {
	// 		xn := dirs[d][0] + i
	// 		yn := dirs[d][1] + j
	// 		if xn >= 0 && xn < rows && yn >= 0 && yn < cols {
	// 			dfs(xn, yn)
	// 		}
	// 	}
	// }
	var bfs func(int, int)
	bfs = func(i, j int) {
		if grid[i][j] == 0 || visited[i][j] {
			return
		}
		visited[i][j] = true
		count++
		var queue [][]int
		queue = append(queue, []int{i, j})
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for d := range dirs {
				xn := dirs[d][0] + cur[0]
				yn := dirs[d][1] + cur[1]
				if xn >= 0 && xn < rows && yn >= 0 && yn < cols {
					bfs(xn, yn)
				}
			}
		}
	}

	for i := 0; i < rows; i++ {
		j := 0
		if grid[i][j] == 1 && !visited[i][j] {
			// dfs(i, j)
			bfs(i, j)
		}
		j = cols - 1
		if grid[i][j] == 1 && !visited[i][j] {
			// dfs(i, j)
			bfs(i, j)
		}
	}
	for j := 0; j < cols; j++ {
		i := 0
		if grid[i][j] == 1 && !visited[i][j] {
			// dfs(i, j)
			bfs(i, j)
		}
		i = rows - 1
		if grid[i][j] == 1 && !visited[i][j] {
			// dfs(i, j)
			bfs(i, j)
		}

	}

	count = 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				// dfs(i, j)
				bfs(i, j)
			}
		}
	}
	return count
}

// @lc code=end
