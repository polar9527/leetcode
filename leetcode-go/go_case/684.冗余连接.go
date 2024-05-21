package go_case

/*
 * @lc app=leetcode.cn id=684 lang=golang
 *
 * [684] 冗余连接
 *
 * https://leetcode.cn/problems/redundant-connection/description/
 *
 * algorithms
 * Medium (67.89%)
 * Likes:    628
 * Dislikes: 0
 * Total Accepted:    116.3K
 * Total Submissions: 171.3K
 * Testcase Example:  '[[1,2],[1,3],[2,3]]'
 *
 * 树可以看成是一个连通且 无环 的 无向 图。
 *
 * 给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n
 * 中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai
 * 和 bi 之间存在一条边。
 *
 * 请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的那个。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入: edges = [[1,2], [1,3], [2,3]]
 * 输出: [2,3]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入: edges = [[1,2], [2,3], [3,4], [1,4], [1,5]]
 * 输出: [1,4]
 *
 *
 *
 *
 * 提示:
 *
 *
 * n == edges.length
 * 3 <= n <= 1000
 * edges[i].length == 2
 * 1 <= ai < bi <= edges.length
 * ai != bi
 * edges 中无重复元素
 * 给定的图是连通的
 *
 *
 */

// @lc code=start
func findRedundantConnection(edges [][]int) []int {

	n := len(edges)
	// 因为题目中节点值是从 1 到 n
	father := make([]int, n+1)
	// 并查集初始化
	for i := 1; i < n+1; i++ {
		father[i] = i
	}

	var find func(int) int // 并查集里寻根的过程
	find = func(x int) int {

		if x == father[x] {
			return x
		}
		// 路径压缩
		father[x] = find(father[x])
		return father[x]
	}
	var join func(u, v int) // 将 v->u 这条边加入并查集
	join = func(u, v int) {
		ur := find(u)
		vr := find(v)
		if ur == vr {
			return
		}
		father[vr] = ur
	}

	// 判断 u 和 v是否找到同一个根，本题用不上
	same := func(u, v int) bool {
		ur := find(u)
		vr := find(v)
		return ur == vr
	}

	// 给定往一棵 n 个节点 (节点值 1～n) 的树中添加   一条边  后的图,
	// 如果有多个答案,意味着 只可能是一个完整的环
	for i := 0; i < len(edges); i++ {
		if same(edges[i][0], edges[i][1]) {
			return edges[i]
		}
		join(edges[i][0], edges[i][1])
	}
	return nil
}

// @lc code=end
