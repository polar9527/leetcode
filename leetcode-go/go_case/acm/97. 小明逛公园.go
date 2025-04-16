package acm

import (
	"fmt"
	"math"
)

func floyd(n int, edges [][]int, plans [][]int) []int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	adjm := make([][]int, n+1)
	for i := range adjm {
		adjm[i] = make([]int, n+1)
		for j := range adjm[i] {
			adjm[i][j] = math.MaxInt64
		}
	}

	for _, e := range edges {
		from, to, val := e[0], e[1], e[2]
		adjm[from][to] = val
		adjm[to][from] = val
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if adjm[i][k] != math.MaxInt64 && adjm[k][j] != math.MaxInt64 {
					adjm[i][j] = min(adjm[i][j], adjm[i][k]+adjm[k][j])
				}
			}
		}
	}

	res := []int{}
	for _, p := range plans {
		start, end := p[0], p[1]
		s := adjm[start][end]
		if s == math.MaxInt64 {
			res = append(res, -1)
		} else {
			res = append(res, s)
		}
	}

	return res
}

func main() {
	var n, en int
	fmt.Scan(&n, &en)

	edges := [][]int{}
	for i := 0; i < en; i++ {
		var f, t, w int
		fmt.Scan(&f, &t, &w)
		edges = append(edges, []int{f, t, w})
	}
	var pn int
	fmt.Scan(&pn)
	plans := [][]int{}
	for i := 0; i < pn; i++ {
		var start, end int
		fmt.Scan(&start, &end)
		plans = append(plans, []int{start, end})
	}

	solutions := floyd(n, edges, plans)

	for i := range solutions {
		fmt.Println(solutions[i])
	}
}
