package acm

import "fmt"

func main() {
	var rows, cols int

	fmt.Scan(&rows, &cols)

	g := make([][]int, rows)
	for i := range g {
		g[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Scan(&g[i][j])
		}
	}

	fmt.Println(islandPerimeter(g))
}

func islandPerimeter(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	dir := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				for k := 0; k < 4; k++ {
					x := i + dir[k][0]
					y := j + dir[k][1]
					if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == 0 {
						res++
					}
				}
			}
		}
	}
	return res
}
