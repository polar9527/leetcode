package go_case

/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 *
 * https://leetcode.cn/problems/island-perimeter/description/
 *
 * algorithms
 * Easy (70.22%)
 * Likes:    753
 * Dislikes: 0
 * Total Accepted:    159.8K
 * Total Submissions: 227.5K
 * Testcase Example:  '[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]'
 *
 * 给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。
 *
 * 网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
 *
 * 岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100
 * 。计算这个岛屿的周长。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
 * 输出：16
 * 解释：它的周长是上面图片中的 16 个黄色的边
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[1]]
 * 输出：4
 *
 *
 * 示例 3：
 *
 *
 * 输入：grid = [[1,0]]
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 * row == grid.length
 * col == grid[i].length
 * 1
 * grid[i][j] 为 0 或 1
 *
 *
 */

// @lc code=start
// func islandPerimeter(grid [][]int) int {
// 	DIRECTIONS := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
// 	rows := len(grid)
// 	if rows == 0 {
// 		return 0
// 	}
// 	cols := len(grid[0])

// 	visited := make([][]bool, rows)
// 	for i := 0; i < rows; i++ {
// 		visited[i] = make([]bool, cols)
// 	}

// 	res := 0
// 	var dfs func(int, int)
// 	dfs = func(i, j int) {
// 		if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == 0 || visited[i][j] {
// 			return
// 		}
// 		visited[i][j] = true
// 		for _, d := range DIRECTIONS {
// 			xn := i + d[0]
// 			yn := j + d[1]
// 			if xn < 0 || xn >= rows || yn < 0 || yn >= cols || grid[xn][yn] == 0 {
// 				res++
// 				continue
// 			} else if !visited[xn][yn] {
// 				dfs(xn, yn)
// 			}
// 		}
// 	}

//		for i := 0; i < rows; i++ {
//			for j := 0; j < cols; j++ {
//				if grid[i][j] == 1 || !visited[i][j] {
//					dfs(i, j)
//				}
//			}
//		}
//		return res
//	}
func islandPerimeter(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	dir := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				for k := 0; k < 4; k++ {
					x := i + dir[k][0]
					y := j + dir[k][1]
					if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == 0 {
						res++
					}
				}
			}
		}
	}
	return res
}

// @lc code=end
