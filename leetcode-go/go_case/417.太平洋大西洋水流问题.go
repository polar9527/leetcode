package go_case

/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 *
 * https://leetcode.cn/problems/pacific-atlantic-water-flow/description/
 *
 * algorithms
 * Medium (56.14%)
 * Likes:    688
 * Dislikes: 0
 * Total Accepted:    100.6K
 * Total Submissions: 179.4K
 * Testcase Example:  '[[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]'
 *
 * 有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。 “太平洋” 处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。
 *
 * 这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵 heights ， heights[r][c] 表示坐标 (r, c)
 * 上单元格 高于海平面的高度 。
 *
 * 岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。
 *
 * 返回网格坐标 result 的 2D 列表 ，其中 result[i] = [ri, ci] 表示雨水从单元格 (ri, ci) 流动
 * 既可流向太平洋也可流向大西洋 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
 * 输出: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
 *
 *
 * 示例 2：
 *
 *
 * 输入: heights = [[2,1],[1,2]]
 * 输出: [[0,0],[0,1],[1,0],[1,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == heights.length
 * n == heights[r].length
 * 1 <= m, n <= 200
 * 0 <= heights[r][c] <= 10^5
 *
 *
 */

// @lc code=start

// func pacificAtlantic(heights [][]int) [][]int {
// 	var DIRECTIONS = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 	var dfs func([][]int, [][]bool, int, int)
// 	dfs = func(heights [][]int, visited [][]bool, i, j int) {
// 		visited[i][j] = true
// 		for _, d := range DIRECTIONS {
// 			x, y := i+d[0], j+d[1]
// 			if x < 0 || x >= len(heights) || y < 0 || y >= len(heights[0]) || heights[i][j] > heights[x][y] || visited[x][y] {
// 				continue
// 			}

// 			dfs(heights, visited, x, y)
// 		}
// 	}

// 	// var bfs func([][]int, [][]bool, int, int)
// 	// bfs = func(heights [][]int, visited [][]bool, i, j int) {
// 	// 	queue := make([][]int, 0)
// 	// 	queue = append(queue, []int{i, j})
// 	// 	visited[i][j] = true
// 	// 	for len(queue) > 0 {
// 	// 		cur := queue[0]
// 	// 		queue = queue[1:]
// 	// 		for _, d := range DIRECTIONS {
// 	// 			x, y := cur[0]+d[0], cur[1]+d[1]
// 	// 			if x < 0 || x >= len(heights) || y < 0 || y >= len(heights[0]) || heights[cur[0]][cur[1]] > heights[x][y] || visited[x][y] {
// 	// 				continue
// 	// 			}
// 	// 			queue = append(queue, []int{x, y})
// 	// 			visited[x][y] = true
// 	// 		}
// 	// 	}
// 	// }

// 	res := make([][]int, 0)
// 	pacific := make([][]bool, len(heights))
// 	atlantic := make([][]bool, len(heights))
// 	for i := 0; i < len(heights); i++ {
// 		pacific[i] = make([]bool, len(heights[0]))
// 		atlantic[i] = make([]bool, len(heights[0]))
// 	}

// 	// 列
// 	for i := 0; i < len(heights); i++ {
// 		dfs(heights, pacific, i, 0)
// 		dfs(heights, atlantic, i, len(heights[0])-1)
// 	}
// 	// 行
// 	for j := 0; j < len(heights[0]); j++ {
// 		dfs(heights, pacific, 0, j)
// 		dfs(heights, atlantic, len(heights)-1, j)
// 	}

// 	for i := 0; i < len(heights); i++ {
// 		for j := 0; j < len(heights[0]); j++ {
// 			if pacific[i][j] && atlantic[i][j] {
// 				res = append(res, []int{i, j})
// 			}
// 		}
// 	}

// 	return res
// }

func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])
	// pacific
	// atlantic

	dir := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var dfs func(int, int, [][]bool)
	dfs = func(x, y int, v [][]bool) {
		if v[x][y] {
			return
		}
		v[x][y] = true
		for d := 0; d < 4; d++ {
			xn := x + dir[d][0]
			yn := y + dir[d][1]
			if xn < 0 || xn >= rows || yn < 0 || yn >= cols {
				continue
			}
			if !v[xn][yn] && heights[x][y] <= heights[xn][yn] {
				dfs(xn, yn, v)
			}
		}
	}

	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	for i := 0; i < cols; i++ {
		// 上边，pacific, x = 0, y = i
		dfs(0, i, pacific)
		// 下边, atlantic, x = rows - 1, y = i
		dfs(rows-1, i, atlantic)

	}

	for j := 0; j < rows; j++ {
		// 左边，pacific, x = j, y = 0
		dfs(j, 0, pacific)
		// 右边, atlantic, x = j, y = cols - 1
		dfs(j, cols-1, atlantic)
	}
	res := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

// @lc code=end
