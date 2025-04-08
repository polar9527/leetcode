package main

import (
	"fmt"
	"sync"
)

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

	var dir = [][2]int{
		[2]int{0, 1},
		[2]int{0, -1},
		[2]int{1, 0},
		[2]int{-1, 0},
	}

	count := 0
	var dfs func(int, int, int)
	dfs = func(x, y, key int) {
		if visited[x][y] || matrix[x][y] == 0 {
			return
		}
		matrix[x][y] = key
		visited[x][y] = true
		count++
		for i := 0; i < 4; i++ {
			xn, yn := x+dir[i][0], y+dir[i][1]
			if xn < 0 || xn >= n || yn < 0 || yn >= m {
				continue
			}
			dfs(xn, yn, key)
		}

	}

	var bfs func(int, int, int)
	bfs = func(x, y, key int) {
		q := [][2]int{[2]int{x, y}}
		visited[x][y] = true
		matrix[x][y] = key
		count++
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			x, y := node[0], node[1]

			for i := 0; i < 4; i++ {
				xn, yn := x+dir[i][0], y+dir[i][1]
				if xn < 0 || xn >= n || yn < 0 || yn >= m {
					continue
				}
				if !visited[xn][yn] && matrix[xn][yn] == 1 {
					q = append(q, [2]int{xn, yn})
					visited[xn][yn] = true
					matrix[xn][yn] = key
					count++
				}
			}

		}
	}
	isALlGrid := true
	matrixNum := make(map[int]int)
	key := 2
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				isALlGrid = false
			}
			if !visited[i][j] && matrix[i][j] == 1 {
				count = 0
				bfs(i, j, key)
				matrixNum[key] = count
				key++
			}
		}
	}

	if isALlGrid {
		fmt.Println(n * m)
		return
	}

	bufferPool := &sync.Pool{
		New: func() interface{} {
			return make(map[int]bool)
		},
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count = 1
			matrixNumVisited := bufferPool.Get().(map[int]bool)
			if matrix[i][j] == 0 {
				for k := 0; k < 4; k++ {
					xn, yn := i+dir[k][0], j+dir[k][1]
					if xn < 0 || xn >= n || yn < 0 || yn >= m {
						continue
					}

					key = matrix[xn][yn]
					if _, ok := matrixNumVisited[key]; !ok {
						count += matrixNum[key]
						matrixNumVisited[key] = true
					}

				}
			}
			if res < count {
				res = count
			}
			for k := range matrixNumVisited {
				delete(matrixNumVisited, k)
			}
			bufferPool.Put(matrixNumVisited)
		}
	}

	fmt.Println(res)
}
