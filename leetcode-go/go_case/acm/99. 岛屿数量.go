package acm

import "fmt"

var dir = [][2]int{
	[2]int{0, 1},
	[2]int{0, -1},
	[2]int{1, 0},
	[2]int{-1, 0},
}

func dfs(matrix [][]int, visited [][]bool, n, m, x, y int) {
	visited[x][y] = true
	for i := 0; i < 4; i++ {
		xn, yn := x+dir[i][0], y+dir[i][1]
		if xn < 0 || xn >= n || yn < 0 || yn >= m {
			continue
		}
		if !visited[xn][yn] && matrix[xn][yn] == 1 {
			dfs(matrix, visited, n, m, xn, yn)
		}
	}

}

func bfs(matrix [][]int, visited [][]bool, n, m, x, y int) {
	q := [][2]int{[2]int{x, y}}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		x, y := node[0], node[1]
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			xn, yn := x+dir[i][0], y+dir[i][1]
			if xn < 0 || xn >= n || yn < 0 || yn >= m {
				continue
			}
			if !visited[xn][yn] && matrix[xn][yn] == 1 {
				q = append(q, [2]int{xn, yn})
			}
		}

	}
}

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] && matrix[i][j] == 1 {
				res++
				bfs(matrix, visited, n, m, i, j)
			}
		}
	}

	fmt.Println(res)
}
