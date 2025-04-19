/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 *
 * https://leetcode.cn/problems/rotate-image/description/
 *
 * algorithms
 * Medium (78.03%)
 * Likes:    2045
 * Dislikes: 0
 * Total Accepted:    746.3K
 * Total Submissions: 953.1K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
 *
 * 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[[7,4,1],[8,5,2],[9,6,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
 * 输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == matrix.length == matrix[i].length
 * 1 <= n <= 20
 * -1000 <= matrix[i][j] <= 1000
 *
 *
 *
 *
 */

// @lc code=start
// func rotate(matrix [][]int) {
// 	n := len(matrix)
// 	tmp := make([][]int, n)
// 	for i := range tmp {
// 		tmp[i] = make([]int, n)
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			tmp[j][n-1-i] = matrix[i][j]
// 		}
// 	}
// 	copy(matrix, tmp)
// }

// 4个值 轮转交换
// 行号i转为逆序列号，列号转为行号
// i, j -> j, n-1-i
// top->right:     x,     y       -> y,      n-1-x
// right->bottom:  y,     n-1-x   -> n-1-x , n-1-y
// bottom->left:   n-1-x, n-1-y   -> n-1-y , x
// left->top:      n-1-y, x       -> x , y
//
// tmp = matrix[x][y]
// matrix[x][y] = matrix[n-1-y,][x]
// matrix[n-1-y,][x] = matrix[n-1-x][n-1-y]
// matrix[n-1-x][n-1-y] = matrix[y][n-1-x]
// matrix[y][n-1-x] = tmp

// n 为 偶数. n**2/4
// A = n/2 * n/2
// n 为 奇数。中心点可以不用反复4次轮转交换，(n-1)**2/4
// B = (n-1)/2 * (n+1)/2

func rotate(matrix [][]int) {
	m := matrix
	n := len(matrix)
	r := func(x, y int) {
		// tmp = m[x][y]
		// m[x][y] = m[n-1-y][x]
		// m[n-1-y][x] = m[n-1-x][n-1-y]
		// m[n-1-x][n-1-y] = m[y][n-1-x]
		// m[y][n-1-x] = tmp
		m[x][y], m[n-1-y][x], m[n-1-x][n-1-y], m[y][n-1-x] = m[n-1-y][x], m[n-1-x][n-1-y], m[y][n-1-x], m[x][y]
	}

	if n%2 == 0 {
		for i := 0; i < n/2; i++ {
			for j := 0; j < n/2; j++ {
				r(i, j)
			}
		}
	} else {
		for i := 0; i < (n-1)/2; i++ {
			for j := 0; j < (n+1)/2; j++ {
				r(i, j)
			}
		}
	}
}

// @lc code=end

