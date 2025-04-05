package acm

import (
	"fmt"
	"strings"
)

func main() {
	// vertex numbers, edge numbers
	var vn, en int
	fmt.Scan(&vn, &en)

	// adjacency matrix
	// vertex -> [1, n], so vn+1
	am := make([][]int, vn+1)
	for i := range am {
		am[i] = make([]int, vn+1)
	}

	for en > 0 {
		var out, in int
		fmt.Scan(&out, &in)
		am[out][in] = 1
		en--
	}

	var results [][]int
	var dfs func(int, []int)
	dfs = func(start int, path []int) {
		if start == vn {
			results = append(results, append([]int{}, path...))
			return
		}
		for next := range am[start] {
			if am[start][next] == 1 {
				dfs(next, append(path, next))
			}

		}
	}

	dfs(1, []int{1})

	if len(results) == 0 {
		fmt.Println(-1)
	} else {
		for _, result := range results {
			var s []string
			for _, num := range result {
				s = append(s, fmt.Sprint(num))
			}
			fmt.Println(strings.Join(s, " "))
		}
	}

}
