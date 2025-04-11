package acm

import (
	"fmt"
	"math"
)

func main() {

	var v, e int

	fmt.Scan(&v, &e)

	grid := make(map[int]map[int]int, e+1)
	for i := range grid {
		grid[i] = make(map[int]int, e+1)
		for j := range grid[i] {
			grid[i][j] = math.MaxInt
		}
	}

	for i := 1; i <= e; i++ {
		var x, y, w int
		fmt.Scan(&x, &y, &w)
		grid[x][y] = w
		grid[y][x] = w
	}

	prim := func() int {
		distances := make([]int, v+1)
		for i := range distances {
			distances[i] = math.MaxInt
		}
		isInPrimTree := make([]bool, v+1)

		for i := 1; i <= v; i++ {
			// 1 选择一个距离prim树最近的顶点vertex
			cur := -1
			minDistance := math.MaxInt
			for j := 1; j <= v; j++ {
				if !isInPrimTree[j] && distances[j] < minDistance {
					minDistance = distances[j]
					cur = j
				}
			}

			// 2 将 vertex 加入 prim树， 同时累加边的权重
			isInPrimTree[cur] = true

			// 3 更新 还未加入prim树的其他顶点 到 prim 树的距离
			for i := 1; i <= v; i++ {
				if !isInPrimTree[i] && grid[cur][i] < distances[i] {
					distances[i] = grid[cur][i]
				}
			}

		}
		res := 0
		for i := 2; i <= v; i++ {
			res += distances[i]
		}
		return res
	}

	fmt.Println(prim())

}
