package acm

import (
	"fmt"
	"sort"
)

func main() {

	var v, e int
	fmt.Scan(&v, &e)

	type Edge struct {
		from, to, val int
	}
	edges := make([]Edge, e)
	for i := 0; i < e; i++ {
		var e Edge
		fmt.Scan(&(e.from), &(e.to), &(e.val))
		edges = append(edges, e)

	}

	ancestor := make([]int, v+1)
	init := func() {
		for i := 1; i <= v; i++ {
			ancestor[i] = i
		}
	}

	var find func(int) int
	find = func(k int) int {
		if k == ancestor[k] {
			return k
		}
		kRoot := find(ancestor[k])
		ancestor[k] = kRoot
		return kRoot
	}

	join := func(s, f int) {
		sr := find(s)
		fr := find(f)
		if sr != fr {
			ancestor[sr] = fr
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		if edges[i].val < edges[j].val {
			return true
		}
		return false
	})

	res := 0
	init()
	for _, e := range edges {
		fr := find(e.from)
		tr := find(e.to)
		if fr != tr {
			join(fr, tr)
			res += e.val
		}
	}
	fmt.Println(res)

}
