/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (32.73%)
 * Likes:    1920
 * Dislikes: 520
 * Total Accepted:    330.4K
 * Total Submissions: 1M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given a matrix of m x n elements (m rows, n columns), return all elements of
 * the matrix in spiral order.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 * Input:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 */
package main

// @lc code=start
func spiralOrder(matrix [][]int) []int {

	var ret []int
	rows := len(matrix)
	if rows == 0 {
		return ret
	}
	cols := len(matrix[0])
	if cols == 0 {
		return ret
	}
	var cycle []int
	start := 0
	for rows > start*2 && cols > start*2 {
		cycle = getCycle(matrix, start)
		ret = append(ret, cycle...)
		start++
	}

	return ret
}

func getCycle(matrix [][]int, start int) []int {

	var ret []int
	rows := len(matrix)
	cols := len(matrix[0])
	rowEnd := rows - 1 - start
	colEnd := cols - 1 - start

	for i := start; i <= colEnd; i++ {
		ret = append(ret, matrix[start][i])
	}

	if rowEnd > start {
		for i := start + 1; i <= rowEnd; i++ {
			ret = append(ret, matrix[i][colEnd])
		}
	}

	if rowEnd > start && colEnd > start {
		for i := colEnd - 1; i >= start; i-- {
			ret = append(ret, matrix[rowEnd][i])
		}
	}
	if rowEnd-1 > start && colEnd > start {
		for i := rowEnd - 1; i >= start+1; i-- {
			ret = append(ret, matrix[i][start])
		}
	}

	return ret
}

// @lc code=end
