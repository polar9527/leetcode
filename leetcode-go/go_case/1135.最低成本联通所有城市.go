package go_case

import "sort"

/*
* 想象一下你是个城市基建规划者，地图上有 n 座城市，它们按以 1 到 n 的次序编号。
*
* 给你整数 n 和一个数组 conections，其中 connections[i] = [xi, yi, costi] 表示将城市 xi 和城市 yi 连接所要的costi（连接是双向的）。
*
* 返回连接所有城市的最低成本，每对城市之间至少有一条路径。如果无法连接所有 n 个城市，返回 -1
*
* 该 最小成本 应该是所用全部连接成本的总和。
*
*
*
* 示例 1：
*
*
*
* 输入：n = 3, conections = [[1,2,5],[1,3,6],[2,3,1]]
* 输出：6
* 解释：选出任意 2 条边都可以连接所有城市，我们从中选取成本最小的 2 条。
* 示例 2： 示例 2：
*
*
*
* 输入：n = 4, conections = [[1,2,3],[3,4,4]]
* 输出：-1
* 解释：即使连通所有的边，也无法连接所有城市。
*
*
* 提示：
*
* 1 <= n <= 104 1 <= n <= 10
* 1 <= connections.length <= 104 1 <= connections.length <= 10
* connections[i].length == 3 connections[i].length == 3
* 1 <= xi, yi <= n
* 1 <= 习，yi <= n
* xi != yi 习 ！= 易
* 0 <= costi <= 105 0 <= 成本 <= 10
 */

// @lc code=start

func minimumCost(n int, connections [][]int) (ans int) {

	sort.Slice(connections, func(i, j int) bool {
		if connections[i][2] < connections[j][2] {
			return true
		}
		return false
	})

	ancestor := make([]int, n+1)

	for i := 1; i <= n; i++ {
		ancestor[i] = i
	}
	var find func(k int) int {
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

	for i := 0; i < len(connections); i++ {
		xr := find(connections[i][0])
		yr := find(connections[i][1])
		if xr ! = yr {
			join(connections[i][0], connections[i][1])
			ans += connections[i][2]
			n--
			if n == 1 {
				return ans
			}
		}
	}
	return -1
}

// @lc code=end
