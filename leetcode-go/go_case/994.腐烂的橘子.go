/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 *
 * https://leetcode.cn/problems/rotting-oranges/description/
 *
 * algorithms
 * Medium (53.71%)
 * Likes:    1017
 * Dislikes: 0
 * Total Accepted:    295K
 * Total Submissions: 543.9K
 * Testcase Example:  '[[2,1,1],[1,1,0],[0,1,1]]'
 *
 * 在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
 *
 *
 * 值 0 代表空单元格；
 * 值 1 代表新鲜橘子；
 * 值 2 代表腐烂的橘子。
 *
 *
 * 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
 *
 * 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
 * 输出：-1
 * 解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。
 *
 *
 * 示例 3：
 *
 *
 * 输入：grid = [[0,2]]
 * 输出：0
 * 解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 10
 * grid[i][j] 仅为 0、1 或 2
 *
 *
 */

// @lc code=start
func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// 遍历整个矩阵，找到所有腐烂的橘子，推入队列
	// 同时记录新鲜橘子的个数
	freshCount := 0
	queue := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				freshCount++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	dir := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	// 遍历腐烂队列中的每个节点，每次添加周边被腐烂的新鲜橘子
	times := 0
	for len(queue) > 0 && freshCount > 0 {
		// 还有腐烂橘子的周边没有访问，
		// 还有新鲜橘子存在

		// 当前所有腐烂橘子往外腐烂一圈
		times++
		oneCircle := queue[:]
		queue = [][]int{}
		for _, loc := range oneCircle {

			for _, d := range dir {
				x := loc[0] + d[0]
				y := loc[1] + d[1]
				if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 {
					// 新鲜橘子腐烂
					grid[x][y] = 2
					freshCount--
					// 进入腐烂队列
					queue = append(queue, []int{x, y})
				}
			}
		}
	}
	if freshCount > 0 {
		return -1
	}
	return times
}

// @lc code=end

