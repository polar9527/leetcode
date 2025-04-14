package acm

import (
	"fmt"
)

func main() {
	MaxInt := 1<<31 - 1
	var stationNum, roadNum int
	fmt.Scan(&stationNum, &roadNum)

	// 有向图
	// init with max
	roads := make([][]int, stationNum+1)
	for i := 1; i <= stationNum; i++ {
		roads[i] = make([]int, stationNum+1)
		for j := 1; j <= stationNum; j++ {
			roads[i][j] = MaxInt
		}
	}

	//record the road map
	for i := 0; i < roadNum; i++ {
		var x, y, w int
		fmt.Scan(&x, &y, &w)
		roads[x][y] = w
	}

	start := 1
	end := stationNum

	minDistance := make([]int, stationNum+1)
	for i := range minDistance {
		minDistance[i] = MaxInt
	}
	minDistance[start] = 0
	visited := make([]bool, stationNum+1)

	for is := 1; is <= stationNum; is++ {
		// 1.找到 里start station 距离最近的 station
		cur := 0
		minDist := MaxInt
		for i := 1; i <= stationNum; i++ {
			if !visited[i] {
				if minDist > minDistance[i] {
					minDist = minDistance[i]
					cur = i
				}
			}
		}
		if cur == 0 {
			break
		}
		// 2.访问cur,将其加入路径
		visited[cur] = true

		// 3.通过 cur 更新 所有 未被访问的 station i 的最短路径, 且从 cur 到 i 必须相连

		for i := 1; i <= stationNum; i++ {
			if !visited[i] && roads[cur][i] != MaxInt {
				if minDistance[i] > minDistance[cur]+roads[cur][i] {
					minDistance[i] = minDistance[cur] + roads[cur][i]
				}
			}
		}
	}

	if minDistance[end] == MaxInt {
		fmt.Println(-1)
	} else {
		fmt.Println(minDistance[end])
	}
}
