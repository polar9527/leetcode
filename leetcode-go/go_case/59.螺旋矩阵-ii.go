/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode.cn/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (70.86%)
 * Likes:    1419
 * Dislikes: 0
 * Total Accepted:    522.9K
 * Total Submissions: 738.2K
 * Testcase Example:  '3'
 *
 * 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：[[1,2,3],[8,9,4],[7,6,5]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	t, b := 0, n-1
	l, r := 0, n-1
	target := n * n
	num := 1
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for num <= target {
		for i := l; i <= r; i++ {
			matrix[t][i] = num
			num++
		}
		t++
		for i := t; i <= b; i++ {
			matrix[i][r] = num
			num++
		}
		r--
		for i := r; i >= l; i-- {
			matrix[b][i] = num
			num++
		}
		b--
		for i := b; i >= t; i-- {
			matrix[i][l] = num
			num++
		}
		l++
	}
	return matrix
}

// @lc code=end

