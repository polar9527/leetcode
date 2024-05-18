package go_case

/*
 * @lc app=leetcode.cn id=827 lang=golang
 *
 * [827] 最大人工岛
 *
 * https://leetcode.cn/problems/making-a-large-island/description/
 *
 * algorithms
 * Hard (47.17%)
 * Likes:    415
 * Dislikes: 0
 * Total Accepted:    53K
 * Total Submissions: 112.3K
 * Testcase Example:  '[[1,0],[0,1]]'
 *
 * 给你一个大小为 n x n 二进制矩阵 grid 。最多 只能将一格 0 变成 1 。
 *
 * 返回执行此操作后，grid 中最大的岛屿面积是多少？
 *
 * 岛屿 由一组上、下、左、右四个方向相连的 1 形成。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: grid = [[1, 0], [0, 1]]
 * 输出: 3
 * 解释: 将一格0变成1，最终连通两个小岛得到面积为 3 的岛屿。
 *
 *
 * 示例 2:
 *
 *
 * 输入: grid = [[1, 1], [1, 0]]
 * 输出: 4
 * 解释: 将一格0变成1，岛屿的面积扩大为 4。
 *
 * 示例 3:
 *
 *
 * 输入: grid = [[1, 1], [1, 1]]
 * 输出: 4
 * 解释: 没有0可以让我们变成1，面积依然为 4。
 *
 *
 *
 * 提示：
 *
 *
 * n == grid.length
 * n == grid[i].length
 * 1 <= n <= 500
 * grid[i][j] 为 0 或 1
 *
 *
 */

// @lc code=start
func largestIsland(grid [][]int) int {
	dir := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	n := len(grid)
	m := len(grid[0])
	area := 0
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	gridNum := make(map[int]int, 0) // 记录每一个岛屿的面积
	mark := 2                       // 记录每个岛屿的编号
	isAllGrid := true
	res := 0 // 标记是否整个地图都是陆地

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var dfs func(grid [][]int, visited [][]bool, x, y, mark int)
	dfs = func(grid [][]int, visited [][]bool, x, y, mark int) {
		// 终止条件：访问过的节点 或者 遇到海水
		if visited[x][y] || grid[x][y] == 0 {
			return
		}
		visited[x][y] = true // 标记访问过
		grid[x][y] = mark    // 给陆地标记新标签
		area++
		for i := 0; i < 4; i++ {
			nextX := x + dir[i][0]
			nextY := y + dir[i][1]
			if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
				continue
			}
			dfs(grid, visited, nextX, nextY, mark)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				isAllGrid = false
			}
			if !visited[i][j] && grid[i][j] == 1 {
				area = 0
				dfs(grid, visited, i, j, mark) // 将与其链接的陆地都标记上 true
				gridNum[mark] = area           // 记录每一个岛屿的面积
				mark++                         // 更新下一个岛屿编号
			}
		}
	}
	if isAllGrid {
		return n * m
	}
	// 根据添加陆地的位置，计算周边岛屿面积之和
	visitedGrid := make(map[int]struct{}) // 标记访问过的岛屿
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count := 1                           // 记录连接之后的岛屿数量
			visitedGrid = make(map[int]struct{}) // 每次使用时，清空
			if grid[i][j] == 0 {
				for k := 0; k < 4; k++ {
					// 计算相邻坐标
					nearI := i + dir[k][0]
					nearJ := j + dir[k][1]
					if nearI < 0 || nearI >= len(grid) || nearJ < 0 || nearJ >= len(grid[0]) {
						continue
					}
					// 添加过的岛屿不要重复添加
					if _, ok := visitedGrid[grid[nearI][nearJ]]; ok {
						continue
					}
					// 把相邻四面的岛屿数量加起来
					count += gridNum[grid[nearI][nearJ]]
					// 标记该岛屿已经添加过
					visitedGrid[grid[nearI][nearJ]] = struct{}{}
				}
			}
			res = max(res, count)
		}
	}
	return res
}

// @lc code=end
