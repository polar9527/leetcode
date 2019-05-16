/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	n := len(matrix)

	a := 0
	b := n - 1

	for a < b {
		for i := 0; i < b-a; i++ {
			matrix[a][a+i], matrix[a+i][b], matrix[b][b-i], matrix[b-i][a] = matrix[b-i][a], matrix[a][a+i], matrix[a+i][b], matrix[b][b-i]
		}
		a++
		b--
	}
}

func main() {
	m := [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}

	rotate(m)
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}
