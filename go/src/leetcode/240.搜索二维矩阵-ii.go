/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package main

import "fmt"

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// matrix[rowIndex][columnIndex]
	var rowLength, columnLength int
	rowLength = len(matrix)
	if rowLength == 0 {
		return false
	}

	columnLength = len(matrix[0])
	if 0 == columnLength {
		return false
	}

	// if rowLength == 1 {
	// 	location := sort.SearchInts(matrix[0], target)
	// 	return location < columnLength && matrix[0][location] == target
	// }

	// if columnLength == 1 {
	// 	location := sort.Search(rowLength, func(i int) {
	// 		return matrix[i][0] >= target
	// 	})
	// 	return location < rowLength && matrix[location][0] == target
	// }
	if rowLength == 1 {
		if target < matrix[0][0] || matrix[0][columnLength-1] < target {
			return false
		} else {
			if ret := binSearch(matrix[0], target); ret >= 0 {
				return true
			} else {
				fmt.Printf("location: %d\n", ret)
				return false
			}
		}
	}

	if columnLength == 1 {
		if target < matrix[0][0] || matrix[rowLength-1][0] < target {
			return false
		} else {
			array := make([]int, rowLength)
			for i := range matrix {
				array[i] = matrix[i][0]
			}
			if ret := binSearch(array, target); ret >= 0 {
				return true
			} else {
				return false
			}
		}
	}

	if target < matrix[0][0] || matrix[rowLength-1][columnLength-1] < target {
		return false
	}

	for r, c := rowLength-1, 0; r >= 0 && c < columnLength; {
		if matrix[r][c] < target {
			c++
		} else if matrix[r][c] > target {
			r--
		} else {
			return true
		}
	}
	return false
}

func binSearch(a []int, target int) int {
	if len(a) == 0 {
		return -1
	}
	if len(a) == 1 {
		if a[0] == target {
			return 0
		}
		return -1
	}
	lo, hi, mi := 0, len(a)-1, len(a)/2
	for ; lo < hi && a[mi] != target; mi = lo + (hi-lo)/2 {
		if a[mi] < target {
			lo = mi + 1
		} else if a[mi] >= target {
			hi = mi
		}
	}
	if a[mi] >= target {
		return mi
	} else {
		return -1
	}
}

// @lc code=end
