/*
 * @lc app=leetcode.cn id=1905 lang=golang
 *
 * [1905] 统计子岛屿
 *
 * https://leetcode.cn/problems/count-sub-islands/description/
 *
 * algorithms
 * Medium (67.87%)
 * Likes:    140
 * Dislikes: 0
 * Total Accepted:    35.6K
 * Total Submissions: 52.5K
 * Testcase Example:  '[[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]]\n' +
  '[[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]'
 *
 * 给你两个 m x n 的二进制矩阵 grid1 和 grid2 ，它们只包含 0 （表示水域）和 1 （表示陆地）。一个 岛屿 是由 四个方向
 * （水平或者竖直）上相邻的 1 组成的区域。任何矩阵以外的区域都视为水域。
 *
 * 如果 grid2 的一个岛屿，被 grid1 的一个岛屿 完全 包含，也就是说 grid2 中该岛屿的每一个格子都被 grid1
 * 中同一个岛屿完全包含，那么我们称 grid2 中的这个岛屿为 子岛屿 。
 *
 * 请你返回 grid2 中 子岛屿 的 数目 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]],
 * grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
 * 输出：3
 * 解释：如上图所示，左边为 grid1 ，右边为 grid2 。
 * grid2 中标红的 1 区域是子岛屿，总共有 3 个子岛屿。
 *
 *
 * 示例 2：
 *
 * 输入：grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]],
 * grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
 * 输出：2
 * 解释：如上图所示，左边为 grid1 ，右边为 grid2 。
 * grid2 中标红的 1 区域是子岛屿，总共有 2 个子岛屿。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid1.length == grid2.length
 * n == grid1[i].length == grid2[i].length
 * 1 <= m, n <= 500
 * grid1[i][j] 和 grid2[i][j] 都要么是 0 要么是 1 。
 *
 *
*/

// @lc code=start

func countSubIslands(grid1 [][]int, grid2 [][]int) int {

	rows, cols := len(grid2), len(grid2[0])

	dir := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var check bool
	var bfs func(int, int) bool
	bfs = func(x, y int) bool {

		q := [][]int{{x, y}}
		grid2[x][y] = 0
		if grid1[x][y] == 0 {
			check = false
		}
		for len(q) > 0 {
			n := q[0]
			q = q[1:]
			for k := 0; k < 4; k++ {
				xn := n[0] + dir[k][0]
				yn := n[1] + dir[k][1]
				if xn < 0 || xn >= rows || yn < 0 || yn >= cols || grid2[xn][yn] == 0 {
					continue
				}

				q = append(q, []int{xn, yn})
				grid2[xn][yn] = 0
				if grid1[xn][yn] == 0 {
					check = false
				}

			}

		}
		return check
	}

	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid2[i][j] == 1 {
				check = true
				if bfs(i, j) {
					res++
				}
			}
		}
	}
	return res
}

// @lc code=end

