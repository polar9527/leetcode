package go_case

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1584 lang=golang
 *
 * [1584] 连接所有点的最小费用
 *
 * https://leetcode.cn/problems/min-cost-to-connect-all-points/description/
 *
 * algorithms
 * Medium (66.11%)
 * Likes:    343
 * Dislikes: 0
 * Total Accepted:    70K
 * Total Submissions: 105.7K
 * Testcase Example:  '[[0,0],[2,2],[3,10],[5,2],[7,0]]'
 *
 * 给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。
 *
 * 连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示
 * val 的绝对值。
 *
 * 请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
 * 输出：20
 * 解释：
 *
 * 我们可以按照上图所示连接所有点得到最小总费用，总费用为 20 。
 * 注意到任意两个点之间只有唯一一条路径互相到达。
 *
 *
 * 示例 2：
 *
 *
 * 输入：points = [[3,12],[-2,5],[-4,1]]
 * 输出：18
 *
 *
 * 示例 3：
 *
 *
 * 输入：points = [[0,0],[1,1],[1,0],[-1,1]]
 * 输出：4
 *
 *
 * 示例 4：
 *
 *
 * 输入：points = [[-1000000,-1000000],[1000000,1000000]]
 * 输出：4000000
 *
 *
 * 示例 5：
 *
 *
 * 输入：points = [[0,0]]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= points.length <= 1000
 * -10^6 <= xi, yi <= 10^6
 * 所有点 (xi, yi) 两两不同。
 *
 *
 */

// @lc code=start
// kruskal
func minCostConnectPoints(points [][]int) int {
	// '[[0,0],[2,2],[3,10],[5,2],[7,0]]'
	nv := len(points)
	abs := func(x, y int) int {
		if x-y > 0 {
			return x - y
		}
		return y - x
	}
	// edges = [][]int{ {x1,y1,val1}, {x2,y2,val2}}
	edges := [][]int{}
	for i := 0; i < nv-1; i++ {
		for j := i + 1; j < nv; j++ {
			e := []int{i, j, abs(points[i][0], points[j][0]) + abs(points[i][1], points[j][1])}
			edges = append(edges, e)
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		if edges[i][2] < edges[j][2] {
			return true
		}
		return false
	})

	// [0, nv-1]
	ancestor := make([]int, nv)
	for i := 0; i < nv; i++ {
		ancestor[i] = i
	}

	var find func(k int) int
	find = func(k int) int {
		if k == ancestor[k] {
			return k
		}
		kr := find(ancestor[k])
		ancestor[k] = kr
		return kr
	}

	join := func(x, y int) {
		xr := find(x)
		yr := find(y)

		if xr != yr {
			ancestor[xr] = yr
		}
	}

	cost := 0
	n := nv
	for i := 0; i < len(edges); i++ {
		xr := find(edges[i][0])
		yr := find(edges[i][1])
		if xr != yr {
			join(xr, yr)
			cost += edges[i][2]
			n--
			// 有 nv 个 point
			// 全部连上就有 nv-1个 边
			if n == 1 {
				return cost
			}
		}
	}
	return 0
}

// @lc code=end
