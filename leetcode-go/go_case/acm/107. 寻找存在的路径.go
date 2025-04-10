package acm

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fatherMap := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		fatherMap[i] = i
	}

	var find func(int) int
	find = func(key int) int {
		if key == fatherMap[key] {
			return key
		}
		root := find(fatherMap[key])
		fatherMap[key] = root
		return root
	}

	isSame := func(u, v int) bool {
		uRoot := find(u)
		vRoot := find(v)
		return uRoot == vRoot
	}

	join := func(u, v int) {
		uRoot := find(u)
		vRoot := find(v)
		if uRoot == vRoot {
			return
		}
		fatherMap[vRoot] = uRoot
	}

	for i := 0; i < m; i++ {
		var s, f int
		fmt.Scan(&s, &f)
		join(s, f)
	}

	var source, destination int
	fmt.Scan(&source, &destination)

	if isSame(source, destination) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

}
