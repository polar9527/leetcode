package acm

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	graph := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < k; i++ {
		var v1, v2 int
		fmt.Scan(&v1, &v2)
		graph[v1] = append(graph[v1], v2)
	}

	visited := make([]bool, n+1)

	var dfs func(n int)
	dfs = func(n int) {
		if visited[n] {
			return
		}
		visited[n] = true
		for _, i := range graph[n] {
			dfs(i)
		}
	}

	dfs(1)
	for i := 1; i <= n; i++ {
		if !visited[i] {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(1)
}
