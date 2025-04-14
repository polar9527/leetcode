package go_case

import (
	"math"
	"slices"
)

/*
 * @lc app=leetcode.cn id=743 lang=golang
 *
 * [743] 网络延迟时间
 *
 * https://leetcode.cn/problems/network-delay-time/description/
 *
 * algorithms
 * Medium (57.95%)
 * Likes:    844
 * Dislikes: 0
 * Total Accepted:    163.8K
 * Total Submissions: 281.7K
 * Testcase Example:  '[[2,1,1],[2,3,1],[3,4,1]]\n4\n2'
 *
 * 有 n 个网络节点，标记为 1 到 n。
 *
 * 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点，
 * wi 是一个信号从源节点传递到目标节点的时间。
 *
 * 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：times = [[1,2,1]], n = 2, k = 1
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：times = [[1,2,1]], n = 2, k = 2
 * 输出：-1
 *
 *
 *
 *
 * 提示：

 *
 *
 * 1 <= k <= n <= 100
 * 1 <= times.length <= 6000
 * times[i].length == 3
 * 1 <= ui, vi <= n
 * ui != vi
 * 0 <= wi <= 100
 * 所有 (ui, vi) 对都 互不相同（即，不含重复边）
 *
 *
 */

// @lc code=start
// dijkstra
func networkDelayTime(times [][]int, n int, k int) int {

	matrix := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		matrix[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			matrix[i][j] = math.MaxInt
		}
	}

	for i := 0; i < len(times); i++ {
		x, y, w := times[i][0], times[i][1], times[i][2]
		// 有向边
		matrix[x][y] = w
	}
	// init
	visited := make([]bool, n+1)
	// 每个节点到 k 节点的 最近距离
	minDistance := make([]int, n+1)
	for i := 0; i <= n; i++ {
		minDistance[i] = math.MaxInt
	}
	minDistance[k] = 0

	for x := 1; x <= n; x++ {
		// 1.找到 离起始节点 k 距离 最近的 未访问 节点
		minDis := math.MaxInt
		curNode := 1
		for i := 1; i <= n; i++ {
			if !visited[i] && minDistance[i] < minDis {
				minDis = minDistance[i]
				curNode = i
			}
		}
		// 此时选出的 curNode 是 离起始节点 k 距离 最近的 未访问 节点
		// 2.访问curNode
		// 当 上一层循环完毕，意味着所有 的 节点都访问完毕
		visited[curNode] = true

		// 3.更新 未访问节点 离起始节点 k 的最近距离
		for i := 1; i <= n; i++ {
			if !visited[i] && matrix[curNode][i] != math.MaxInt {
				// 对于 所有未访问节点 i
				// minDistance[i],  k -> i 的最短距离
				// minDistance[curNode] + matrix[curNode][i], k -> curNode -> i 的 最短距离
				// minDistance[curNode], k -> curNode的最短距离
				// 所有 i 的 k -> curNode -> i 和 k -> curNode 这些数据中取最小值
				if minDistance[i] > minDistance[curNode]+matrix[curNode][i] {
					minDistance[i] = minDistance[curNode] + matrix[curNode][i]
				}
			}
		}
	}
	res := slices.Max(minDistance[1:])
	if res == math.MaxInt {
		return -1
	}
	return res
}

// @lc code=end
