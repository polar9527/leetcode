/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (39.78%)
 * Likes:    382
 * Dislikes: 0
 * Total Accepted:    61.2K
 * Total Submissions: 152.2K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * 输出: [1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2:
 *
 * 输入:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 */
package leetcode

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	y := len(matrix)
	if y == 0 {
		return nil
	}
	x := len(matrix[0])

	ans := make([]int, x*y)
	top := 0
	bottom := y - 1
	left := 0
	right := x - 1
	index := 0
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			ans[index] = matrix[top][i]
			index++
		}
		for i := top + 1; i <= bottom; i++ {
			ans[index] = matrix[i][right]
			index++
		}
		if top < bottom && left < right {
			for i := right - 1; i >= left; i-- {
				ans[index] = matrix[bottom][i]
				index++
			}
			for i := bottom - 1; i > top; i-- {
				ans[index] = matrix[i][left]
				index++
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return ans
}

// @lc code=end
