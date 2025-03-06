package go_case

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode.cn/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (51.10%)
 * Likes:    1878
 * Dislikes: 0
 * Total Accepted:    680.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1
 * -100
 *
 *
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	t, b := 0, len(matrix)-1
	l, r := 0, len(matrix[0])-1
	amount := (b + 1) * (r + 1)
	nums := make([]int, amount)
	num := 1
	// matrix[i][j]
	for num <= amount {
		for j := l; j <= r && num <= amount; j++ {
			nums[num-1] = matrix[t][j]
			num++
		}
		t++
		for i := t; i <= b && num <= amount; i++ {
			nums[num-1] = matrix[i][r]
			num++
		}
		r--
		for j := r; j >= l && num <= amount; j-- {
			nums[num-1] = matrix[b][j]
			num++
		}
		b--
		for i := b; i >= t && num <= amount; i-- {
			nums[num-1] = matrix[i][l]
			num++
		}
		l++

	}
	return nums
}

// @lc code=end
