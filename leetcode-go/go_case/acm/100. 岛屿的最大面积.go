package acm

import "fmt"

func main() {
	var rows, cols int
	fmt.Scan(&rows, &cols)

	vertexes := make([][]int, rows)
	for i := range vertexes {
		vertexes[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Scan(&vertexes[i][j])
		}
	}

	di := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	count := 0
	var dfs func(int, int)
	dfs = func(x, y int) {
		if visited[x][y] || vertexes[x][y] == 0 {
			return
		}
		visited[x][y] = true
		count++
		for i := 0; i < 4; i++ {
			xn := x + di[i][0]
			yn := y + di[i][1]
			if xn < 0 || xn >= rows || yn < 0 || yn >= cols {
				continue
			}
			if !visited[xn][yn] && vertexes[xn][yn] == 1 {
				dfs(xn, yn)
			}
		}
	}

	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] || vertexes[i][j] == 1 {
				count = 0
				dfs(i, j)
				if count > res {
					res = count
				}
			}
		}
	}
	fmt.Println(res)
}
