package acm

import (
	"fmt"
	"strings"
)

func main() {
	var rows, cols int
	fmt.Scan(&rows, &cols)

	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Scan(&grid[i][j])
		}
	}
	dir := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	// count := 0
	var bfs func(int, int)
	bfs = func(x, y int) {
		q := [][]int{[]int{x, y}}
		grid[x][y] = 2
		for len(q) > 0 {
			curx := q[0][0]
			cury := q[0][1]
			q = q[1:]
			for i := 0; i < 4; i++ {
				xn := curx + dir[i][0]
				yn := cury + dir[i][1]
				if xn < 0 || xn >= rows || yn < 0 || yn >= cols {
					continue
				}
				if grid[xn][yn] == 1 {
					q = append(q, []int{xn, yn})
					grid[xn][yn] = 2
				}
			}

		}
	}

	for i := 0; i < rows; i++ {
		if grid[i][0] == 1 {
			bfs(i, 0)
		}
		if grid[i][cols-1] == 1 {
			bfs(i, cols-1)
		}
	}

	for j := 0; j < cols; j++ {
		if grid[0][j] == 1 {
			bfs(0, j)
		}
		if grid[rows-1][j] == 1 {
			bfs(rows-1, j)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				grid[i][j] = 1
			} else if grid[i][j] == 1 {
				grid[i][j] = 0
			}
		}
	}

	for _, p := range grid {
		var s []string
		for _, vertex := range p {
			s = append(s, fmt.Sprint(vertex))
		}
		fmt.Println(strings.Join(s, " "))
	}
}
