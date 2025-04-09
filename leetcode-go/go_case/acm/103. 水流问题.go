package acm

import (
	"fmt"
	"strings"
)

func main() {
	var rows, cols int

	fmt.Scan(&rows, &cols)

	heights := make([][]int, rows)
	for i := range heights {
		heights[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Scan(&heights[i][j])
		}
	}

	vLeftTop := make([][]bool, rows)
	vRightBottom := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		vLeftTop[i] = make([]bool, cols)
		vRightBottom[i] = make([]bool, cols)
	}

	dir := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var dfs func(int, int, [][]bool)
	dfs = func(x, y int, v [][]bool) {
		if v[x][y] {
			return
		}
		v[x][y] = true
		for d := 0; d < 4; d++ {
			xn := x + dir[d][0]
			yn := y + dir[d][1]
			if xn < 0 || xn >= rows || yn < 0 || yn >= cols {
				continue
			}
			if !v[xn][yn] && heights[x][y] <= heights[xn][yn] {
				dfs(xn, yn, v)
			}
		}
	}

	for i := 0; i < cols; i++ {
		// 上边
		dfs(0, i, vLeftTop)
		// 下边
		dfs(rows-1, i, vRightBottom)

	}

	for j := 0; j < rows; j++ {
		// 左边
		dfs(j, 0, vLeftTop)
		// 右边
		dfs(j, cols-1, vRightBottom)
	}
	res := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if vLeftTop[i][j] && vRightBottom[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	for i := range res {
		s := []string{}
		for _, num := range res[i] {
			s = append(s, fmt.Sprint(num))
		}
		fmt.Println(strings.Join(s, " "))
	}
}
