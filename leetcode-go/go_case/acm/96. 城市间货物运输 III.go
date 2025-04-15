package acm

import (
	"fmt"
	"math"
)

func bellmanford(n int, edges [][]int, src, dst, k int) int {
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
	minDist[src] = 0
	backup := make([]int, n+1)
	for i := 1; i <= k+1; i++ {
		copy(backup, minDist)
		for _, e := range edges {
			from, to, cost := e[0], e[1], e[2]
			if backup[from] == math.MaxInt64 {
				continue
			}
			minDist[to] = min(minDist[to], backup[from]+cost)
		}
	}
	return minDist[dst]
}

func spfa(n int, edges [][]int, src, dst, k int) int {

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
	minDist[src] = 0

	start := src

	inq := make([]bool, n+1)
	q := []int{start}

	inq[start] = true
	backup := make([]int, n+1)
	k = k + 1
	for len(q) > 0 && k > 0 {
		k--

		iLevel := len(q)
		// 每一层遍历完之后，先跳出内层循环，检查 k > 0
		for ; iLevel > 0; iLevel-- {
			cur := q[0]
			q = q[1:]
			inq[cur] = false
			copy(backup, minDist)
			for _, e := range matrix[cur] {
				from, to, cost := cur, e.to, e.cost
				if minDist[to] > backup[from]+cost {
					minDist[to] = backup[from] + cost
					if !inq[to] {
						q = append(q, to)
						inq[to] = true
					}
				}
			}
		}

	}
	return minDist[dst]
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
	var src, dst, k int
	fmt.Scan(&src, &dst, &k)

	result := spfa(v, edges, src, dst, k)

	if result == math.MaxInt64 {
		fmt.Println("unreachable")
	} else {
		fmt.Println(result)
	}
}
