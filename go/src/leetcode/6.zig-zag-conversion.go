/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// @lc code=start
// package main

// import (
// 	"fmt"
// )

func convert(s string, numRows int) string {
	if numRows == 0 || numRows == 1 {
		return s
	}
	if len(s) < 3 {
		return s
	}
	var zmap = make([][]rune, numRows)
	numRows--
	var row int
	for i, _ := range zmap {
		zmap[i] = make([]rune, 0)
	}
	for i, r := range s {
		if (i/numRows)%2 == 0 {
			row = i % numRows
		} else {
			row = numRows - (i % numRows)
		}
		zmap[row] = append(zmap[row], r)
	}

	result := make([]rune, 0)

	for _, stringList := range zmap {
		for _, char := range stringList {
			result = append(result, char)
		}
	}

	return string(result)
}

// func main() {
// 	// s := "LEETCODEISHIRING"
// 	s1 := "PAYPALISHIRING"
// 	ret := convert(s1, 4)
// 	fmt.Println(ret)
// 	fmt.Println("LCIRETOESIIGEDHN")

// }

// @lc code=end
