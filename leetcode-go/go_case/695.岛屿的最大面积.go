package go_case

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
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	di := [][2]int{
		[2]int{0, 1},
		[2]int{0, -1},
		[2]int{1, 0},
		[2]int{-1, 0},
	}
	var count int

	var dfs func(int, int)
	dfs = func(x, y int) {
		visited[x][y] = true
		count++
		for i := 0; i < 4; i++ {
			xn := x + di[i][0]
			yn := y + di[i][1]
			if xn < 0 || xn >= m || yn < 0 || yn >= n {
				continue
			}
			if !visited[xn][yn] && grid[xn][yn] == 1 {
				dfs(xn, yn)
			}
		}
	}

	var bfs func(int, int)
	bfs = func(x, y int) {
		q := [][]int{[]int{x, y}}
		visited[x][y] = true
		count++
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			curx := cur[0]
			cury := cur[1]

			for i := 0; i < 4; i++ {
				xn := curx + di[i][0]
				yn := cury + di[i][1]
				if xn < 0 || xn >= m || yn < 0 || yn >= n {
					continue
				}
				if !visited[xn][yn] && grid[xn][yn] == 1 {
					visited[xn][yn] = true
					count++
					q = append(q, []int{xn, yn})
				}
			}

		}

	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				count = 0
				bfs(i, j)
				res = max(res, count)
			}
		}
	}
	return res
}

// @lc code=end
