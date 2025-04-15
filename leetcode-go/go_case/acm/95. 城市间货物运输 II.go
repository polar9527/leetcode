package acm

import (
	"fmt"
	"math"
)

func bellmanford(n int, edges [][]int) (int, bool) {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	minDist := make([]int, n+1)
	for i := range minDist {
		minDist[i] = math.MaxInt64
	}
	minDist[1] = 0

	var isNegativeCycle bool
	for i := 1; i <= n-1; i++ {
		for _, e := range edges {
			from, to, cost := e[0], e[1], e[2]
			if minDist[from] == math.MaxInt64 {
				continue
			}
			if i < n {
				minDist[to] = min(minDist[to], minDist[from]+cost)
			} else if minDist[to] > min(minDist[to], minDist[from]+cost) {
				isNegativeCycle = true
			}
		}
	}
	return minDist[n], isNegativeCycle
}

func spfa(n int, edges [][]int) (int, bool) {

	type Edge struct {
		to, cost int
	}
	matrix := make([][]Edge, n+1)
	for _, e := range edges {
		matrix[e[0]] = append(matrix[e[0]], Edge{to: e[1], cost: e[2]})
	}

	minDist := make([]int, n+1)
	for i := range minDist {
		minDist[i] = math.MaxInt64
	}
	minDist[1] = 0

	start := 1
	count := make([]int, n+1)
	inq := make([]bool, n+1)
	q := []int{start}
	count[start]++
	inq[start] = true

	var isNegativeCycle bool
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		inq[cur] = false

		for _, e := range matrix[cur] {
			from, to, cost := cur, e.to, e.cost
			if minDist[to] > minDist[from]+cost {
				minDist[to] = minDist[from] + cost
				if !inq[to] {
					q = append(q, to)
					inq[to] = true
				}
				count[to]++
				if count[to] == n {
					isNegativeCycle = true
					for len(q) > 0 {
						q = q[1:]
					}
					break
				}
			}
		}
	}
	return minDist[n], isNegativeCycle
}

func main() {

	var v, e int
	fmt.Scan(&v, &e)

	edges := make([][]int, e)
	for i := 0; i < e; i++ {
		var f, t, v int
		fmt.Scan(&f, &t, &v)
		edges[i] = []int{f, t, v}
	}

	result, isNegativeCycle := spfa(v, edges)

	if isNegativeCycle {
		fmt.Println("circle")
		return
	}
	if result == math.MaxInt64 {
		fmt.Println("unconnected")
	} else {
		fmt.Println(result)
	}
}
