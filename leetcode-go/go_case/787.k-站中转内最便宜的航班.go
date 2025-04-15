package go_case

/*
 * @lc app=leetcode.cn id=787 lang=golang
 *
 * [787] K 站中转内最便宜的航班
 *
 * https://leetcode.cn/problems/cheapest-flights-within-k-stops/description/
 *
 * algorithms
 * Medium (40.25%)
 * Likes:    700
 * Dislikes: 0
 * Total Accepted:    84.4K
 * Total Submissions: 209.5K
 * Testcase Example:  '4\n[[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]]\n0\n3\n1'
 *
 * 有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei]
 * ，表示该航班都从城市 fromi 开始，以价格 pricei 抵达 toi。
 *
 * 现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的
 * 价格最便宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入:
 * n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src =
 * 0, dst = 3, k = 1
 * 输出: 700
 * 解释: 城市航班图如上
 * 从城市 0 到城市 3 经过最多 1 站的最佳路径用红色标记，费用为 100 + 600 = 700。
 * 请注意，通过城市 [0, 1, 2, 3] 的路径更便宜，但无效，因为它经过了 2 站。
 *
 *
 * 示例 2：
 *
 *
 * 输入:
 * n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
 * 输出: 200
 * 解释:
 * 城市航班图如上
 * 从城市 0 到城市 2 经过最多 1 站的最佳路径标记为红色，费用为 100 + 100 = 200。
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
 * 输出：500
 * 解释：
 * 城市航班图如上
 * 从城市 0 到城市 2 不经过站点的最佳路径标记为红色，费用为 500。
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 100
 * 0 <= flights.length <= (n * (n - 1) / 2)
 * flights[i].length == 3
 * 0 <= fromi, toi < n
 * fromi != toi
 * 1 <= pricei <= 10^4
 * 航班没有重复，且不存在自环
 * 0 <= src, dst, k < n
 * src != dst
 *
 *
 */

// @lc code=start
// bellman-ford with edge limit
// func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

// 	// n个节点，
// 	// k次中转，也就是 k+2 个节点，k+1 条边

// 	// minDist[0] 不使用
// 	minDist := make([]int, n)
// 	backup := make([]int, n)
// 	// init all
// 	for i := range minDist {
// 		minDist[i] = math.MaxInt64
// 	}
// 	// init the src node
// 	minDist[src] = 0

// 	// relax k+1 tines
// 	for i := 1; i <= k+1; i++ {
// 		copy(backup, minDist)
// 		for _, f := range flights {
// 			from, to, cost := f[0], f[1], f[2]
// 			if backup[from] == math.MaxInt64 {
// 				continue
// 			}
// 			if minDist[to] > backup[from]+cost {
// 				minDist[to] = backup[from] + cost
// 			}
// 		}
// 	}
// 	if minDist[dst] == math.MaxInt64 {
// 		return -1
// 	}

// 	return minDist[dst]
// }

// dfs + memory
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	n += 10
	memo := make([][]int, n)
	g := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		g[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	for _, e := range flights {
		g[e[0]][e[1]] = e[2]
	}
	inf := int(1e6)
	var dfs func(u, k int) int
	dfs = func(u, k int) int {
		if memo[u][k] != -1 {
			return memo[u][k]
		}
		if u == dst {
			return 0
		}
		if k <= 0 {
			return inf
		}
		ans := inf
		for v, p := range g[u] {
			if p > 0 {
				ans = min(ans, dfs(v, k-1)+p)
			}
		}
		memo[u][k] = ans
		return ans
	}
	ans := dfs(src, k+1)
	if ans >= inf {
		return -1
	}
	return ans
}

// @lc code=end
