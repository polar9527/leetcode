/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (35.67%)
 * Likes:    1326
 * Dislikes: 140
 * Total Accepted:    286.4K
 * Total Submissions: 798.2K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 *
 *
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * Output: false
 *
 */
package main

import (
	"fmt"
)

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// fmt.Println(matrix)
	// fmt.Println(target)
	row := len(matrix)
	if row == 0 {
		return false
	}
	column := len(matrix[0])
	if column == 0 {
		return false
	}

	for r, c := 0, column-1; r < row && c >= 0; {
		val := matrix[r][c]
		if target == val {
			return true
		} else if target < val {
			c--
		} else {
			r++
		}
	}
	return false
}

// @lc code=end

func main() {
	// m := [][]int{
	// 	{1,   3,  5,  7},
	// 	{10, 11, 16, 20},
	// 	{23, 30, 34, 50},
	// }
	m := [][]int{
		{1, 1},
	}
	t := 16
	ret := searchMatrix(m, t)
	fmt.Println(ret)
}
