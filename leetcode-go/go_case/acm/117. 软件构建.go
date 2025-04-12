package acm

import (
	"fmt"
	"strings"
)

func main() {
	var fn, dn int
	fmt.Scan(&fn, &dn)

	inDegree := make(map[int]int)
	dependencies := make(map[int][]int)
	for i := 0; i < dn; i++ {
		var from, to int
		fmt.Scan(&from, &to)
		inDegree[to]++
		dependencies[from] = append(dependencies[from], to)
	}

	que := make([]int, 0)
	for i := 0; i < fn; i++ {
		if inDegree[i] == 0 {
			que = append(que, i)
		}
	}

	res := []int{}
	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		res = append(res, cur)

		files := dependencies[cur]
		if len(files) != 0 {
			for _, f := range files {
				inDegree[f]--
				if inDegree[f] == 0 {
					que = append(que, f)
				}
			}
		}
	}
	if len(res) == fn {
		var s []string
		for _, r := range res {
			s = append(s, fmt.Sprint(r))
		}
		fmt.Println(strings.Join(s, " "))
		return
	}
	fmt.Println(-1)

}
