/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 *
 * https://leetcode.cn/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (42.13%)
 * Likes:    1401
 * Dislikes: 0
 * Total Accepted:    609.5K
 * Total Submissions: 1.4M
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 给定一个 m x n 的整数数组 grid。一个机器人初始位于 左上角（即 grid[0][0]）。机器人尝试移动到 右下角（即 grid[m -
 * 1][n - 1]）。机器人每次只能向下或者向右移动一步。
 *
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。机器人的移动路径中不能包含 任何 有障碍物的方格。
 *
 * 返回机器人能够到达右下角的不同路径数量。
 *
 * 测试用例保证答案小于等于 2 * 10^9。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
 * 输出：2
 * 解释：3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 *
 *
 * 示例 2：
 *
 *
 * 输入：obstacleGrid = [[0,1],[0,0]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == obstacleGrid.length
 * n == obstacleGrid[i].length
 * 1 <= m, n <= 100
 * obstacleGrid[i][j] 为 0 或 1
 *
 *
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rl, cl := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, rl)
	for r := 0; r < rl; r++ {
		dp[r] = make([]int, cl)
	}

	init := 1
	for r := 0; r < rl; r++ {
		if obstacleGrid[r][0] == 1 {
			init = 0
		}
		dp[r][0] = init
	}

	init = 1
	for c := 0; c < cl; c++ {
		if obstacleGrid[0][c] == 1 {
			init = 0
		}
		dp[0][c] = init
	}

	for r := 1; r < rl; r++ {
		for c := 1; c < cl; c++ {
			if obstacleGrid[r][c] == 0 {
				dp[r][c] = dp[r][c-1] + dp[r-1][c]
			}
		}
	}
	return dp[rl-1][cl-1]
}

// @lc code=end

