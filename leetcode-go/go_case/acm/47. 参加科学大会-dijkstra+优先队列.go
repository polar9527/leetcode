package acm

import (
	"container/heap"
	"fmt"
)

type Road struct {
	t, v int
}

type Item struct {
	node, dist int
}

type PriorityQueue []*Item

func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q PriorityQueue) Less(i, j int) bool {
	if q[i].dist < q[j].dist {
		return true
	}
	return false
}

func (q *PriorityQueue) Push(e interface{}) {
	*q = append(*q, e.(*Item))
}

func (q *PriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	ret := old[n-1]
	*q = old[:n-1]
	return ret
}

func dijkstra(stationNum, roadNum int, edges [][]int, start, end int) int {
	MaxInt := 1<<31 - 1
	// 0 不用
	adjlist := make([][]Road, stationNum+1)
	for _, e := range edges {
		adjlist[e[0]] = append(adjlist[e[0]], Road{t: e[1], v: e[2]})
	}

	visited := make([]bool, stationNum+1)
	minDistance := make([]int, stationNum+1)
	for i := 0; i <= stationNum; i++ {
		minDistance[i] = MaxInt
	}
	minDistance[start] = 0

	q := new(PriorityQueue)
	heap.Init(q)
	heap.Push(q, &Item{node: start, dist: 0})

	for q.Len() > 0 {

		// 1.
		cur := heap.Pop(q).(*Item)

		// 2.
		visited[cur.node] = true

		// 3.
		nextNodes := adjlist[cur.node]
		for _, n := range nextNodes {
			if !visited[n.t] {
				if minDistance[n.t] > minDistance[cur.node]+n.v {
					minDistance[n.t] = minDistance[cur.node] + n.v
					heap.Push(q, &Item{node: n.t, dist: minDistance[n.t]})
				}
			}
		}

	}

	if minDistance[end] == MaxInt {
		return -1
	}
	return minDistance[end]
}

func main() {

	var stationNum, roadNum int
	fmt.Scan(&stationNum, &roadNum)

	edges := make([][]int, roadNum)
	for i := range edges {
		var f, t, v int
		fmt.Scan(&f, &t, &v)
		edges[i] = []int{f, t, v}
	}

	start := 1
	end := stationNum

	result := dijkstra(stationNum, roadNum, edges, start, end)

	fmt.Println(result)
}
