package main

import (
	"fmt"
)

func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	col := len(grid)
	row := len(grid[0])
	dp := make([]int, col*row)
	dp[0] = grid[0][0]
	for c := 0; c < col; c++ {
		for r := 0; r < row; r++ {
			left := 0
			up := 0
			if c > 0 {
				left = dp[(c-1)*row+r]
			}
			if r > 0 {
				up = dp[c*row+(r-1)]
			}
			dp[c*row+r] = grid[c][r] + max(left, up)
		}
	}
	return dp[col*row-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	g := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	ret := maxValue(g)
	fmt.Println(ret)
}
