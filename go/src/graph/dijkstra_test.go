package graph

import (
	"fmt"
	"strconv"
	"testing"
)

// V - S = T
type Dijkstra struct {
	Visit bool   // 表示是否访问
	Val   int    // 表示距离
	Path  string // 路径的显示
}

const (
	IntMax = 1<<32 - 1
)

func getShortPathByDijkstra(begin int, vertex [][]int) []Dijkstra {
	if 0 == len(vertex) || 0 == len(vertex[0]) || len(vertex) != len(vertex[0]) {
		return []Dijkstra{}
	}
	d := make([]Dijkstra, len(vertex))

	for i := 0; i < len(vertex); i++ {
		d[i].Visit = false
		d[i].Val = vertex[begin-1][i]
		d[i].Path = "V" + strconv.Itoa(begin) + " -> V" + strconv.Itoa(i+1)
	}

	d[begin-1].Visit = true
	d[begin-1].Val = 0
	count := 1
	for count < len(vertex) { // 从源点到目的点依次加入最短路径节点进去
		min := IntMax
		temp := 0
		for i := 0; i < len(vertex); i++ { // 找源点到集合T中最短路径的点加入到S中
			if !d[i].Visit && d[i].Val < min {
				min = d[i].Val
				temp = i
			}
		}
		d[temp].Visit = true
		for i := 0; i < len(vertex); i++ { // 进行松弛,即更新源点到各节点的最短路径

			if !d[i].Visit && vertex[temp][i] != IntMax && d[temp].Val+vertex[temp][i] < d[i].Val {
				d[i].Val = d[temp].Val + vertex[temp][i]
				d[i].Path = d[temp].Path + " -> V" + strconv.Itoa(i+1)
			}
		}
		count++
	}

	return d
}

func TestDij(t *testing.T) {
	var vertex [][]int
	for i := 0; i < 8; i++ {
		tmp := make([]int, 8)
		vertex = append(vertex, tmp)
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			vertex[i][j] = IntMax
		}
	}
	vertex[0][1] = 4
	vertex[0][3] = 6
	vertex[0][6] = 7
	vertex[1][2] = 12
	vertex[2][3] = 9
	vertex[2][4] = 1
	vertex[2][5] = 2
	vertex[2][7] = 10
	vertex[3][4] = 13
	vertex[3][6] = 2
	vertex[4][5] = 5
	vertex[4][6] = 11
	vertex[4][7] = 8
	vertex[5][7] = 7
	vertex[6][7] = 14

	d := getShortPathByDijkstra(1, vertex)
	fmt.Println(d)
}
