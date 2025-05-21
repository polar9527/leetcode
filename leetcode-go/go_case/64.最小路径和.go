/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode.cn/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (71.69%)
 * Likes:    1825
 * Dislikes: 0
 * Total Accepted:    776.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 200
 * 0 <= grid[i][j] <= 200
 *
 *
 */

// @lc code=start
// dfs
// func minPathSum(grid [][]int) int {
// 	m, n := len(grid), len(grid[0])
// 	memo := make([][]int, m)
// 	for i := range memo {
// 		memo[i] = make([]int, n)
// 		for j := range memo[i] {
// 			memo[i][j] = -1 // -1 表示没有计算过
// 		}
// 	}
// 	var dfs func(int, int) int
// 	dfs = func(i, j int) int {
// 		if i < 0 || j < 0 {
// 			return math.MaxInt
// 		}
// 		if i == 0 && j == 0 {
// 			return grid[i][j]
// 		}
// 		p := &memo[i][j]
// 		if *p == -1 { // 没有计算过
// 			*p = min(dfs(i, j-1), dfs(i-1, j)) + grid[i][j]
// 		}
// 		return *p
// 	}
// 	return dfs(m-1, n-1)
// }

// DP
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
		if i == 0 {
			for j := range f[0] {
				f[0][j] = math.MaxInt
			}
		}
		f[i][0] = math.MaxInt
	}

	for i, row := range grid {
		for j, x := range row {
			if i == 0 && j == 0 {
				f[1][1] = x
			} else {
				f[i+1][j+1] = min(f[i+1][j], f[i][j+1]) + x
			}
		}
	}
	return f[m][n]
}

// @lc code=end

