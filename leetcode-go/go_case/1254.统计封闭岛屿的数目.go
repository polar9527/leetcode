package go_case

/*
 * @lc app=leetcode.cn id=1254 lang=golang
 *
 * [1254] 统计封闭岛屿的数目
 *
 * https://leetcode.cn/problems/number-of-closed-islands/description/
 *
 * algorithms
 * Medium (64.40%)
 * Likes:    319
 * Dislikes: 0
 * Total Accepted:    73.4K
 * Total Submissions: 113.8K
 * Testcase Example:  '[[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]'
 *
 * 二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。
 *
 * 请返回 封闭岛屿 的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：grid =
 * [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
 * 输出：2
 * 解释：
 * 灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：grid = [[1,1,1,1,1,1,1],
 * [1,0,0,0,0,0,1],
 * [1,0,1,1,1,0,1],
 * [1,0,1,0,1,0,1],
 * [1,0,1,1,1,0,1],
 * [1,0,0,0,0,0,1],
 * ⁠            [1,1,1,1,1,1,1]]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= grid.length, grid[0].length <= 100
 * 0 <= grid[i][j] <=1
 *
 *
 */

// @lc code=start
func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	// visited := make([][]int, m)
	// for i := range visited {
	// 	visited[i] = make([]int, n)
	// }

	bi := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var dfs func(int, int)
	dfs = func(x, y int) {
		grid[x][y] = 1
		for i := 0; i < 4; i++ {
			xn := x + bi[i][0]
			yn := y + bi[i][1]
			if xn < 0 || xn >= m || yn < 0 || yn >= n {
				continue
			}
			if grid[xn][yn] == 0 {
				dfs(xn, yn)
			}
		}
	}
	// 矩阵左右两边
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			dfs(i, 0)
		}
		if grid[i][n-1] == 0 {
			dfs(i, n-1)
		}
	}

	// 矩阵上下两边
	for j := 0; j < n; j++ {
		if grid[0][j] == 0 {
			dfs(0, j)
		}
		if grid[m-1][j] == 0 {
			dfs(m-1, j)
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

// @lc code=end
