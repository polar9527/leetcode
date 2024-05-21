package go_case

/*
 * @lc app=leetcode.cn id=685 lang=golang
 *
 * [685] 冗余连接 II
 *
 * https://leetcode.cn/problems/redundant-connection-ii/description/
 *
 * algorithms
 * Hard (42.18%)
 * Likes:    405
 * Dislikes: 0
 * Total Accepted:    36.3K
 * Total Submissions: 86K
 * Testcase Example:  '[[1,2],[1,3],[2,3]]'
 *
 * 在本问题中，有根树指满足以下条件的 有向
 * 图。该树只有一个根节点，所有其他节点都是该根节点的后继。该树除了根节点之外的每一个节点都有且只有一个父节点，而根节点没有父节点。
 *
 * 输入一个有向图，该图由一个有着 n 个节点（节点值不重复，从 1 到 n）的树及一条附加的有向边构成。附加的边包含在 1 到 n
 * 中的两个不同顶点间，这条附加的边不属于树中已存在的边。
 *
 * 结果图是一个以边组成的二维数组 edges 。 每个元素是一对 [ui, vi]，用以表示 有向 图中连接顶点 ui 和顶点 vi 的边，其中 ui 是
 * vi 的一个父节点。
 *
 * 返回一条能删除的边，使得剩下的图是有 n 个节点的有根树。若有多个答案，返回最后出现在给定二维数组的答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：edges = [[1,2],[1,3],[2,3]]
 * 输出：[2,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：edges = [[1,2],[2,3],[3,4],[4,1],[1,5]]
 * 输出：[4,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == edges.length
 * 3 <= n <= 1000
 * edges[i].length == 2
 * 1 <= ui, vi <= n
 *
 *
 */

// @lc code=start
// 全局变量

func findRedundantDirectedConnection(edges [][]int) []int {

	n := len(edges)
	father := make([]int, n+1)

	// 并查集初始化
	initialize := func() {
		for i := 1; i < n+1; i++ {
			father[i] = i
		}
	}

	// 并查集里寻根的过程
	var find func(int) int
	find = func(u int) int {
		if u == father[u] {
			return u
		}
		father[u] = find(father[u])
		return father[u]
	}

	// 将v->u 这条边加入并查集
	join := func(u, v int) {
		u = find(u)
		v = find(v)
		if u == v {
			return
		}
		father[v] = u
	}

	// 判断 u 和 v是否找到同一个根，本题用不上
	same := func(u, v int) bool {
		u = find(u)
		v = find(v)
		return u == v
	}

	// getRemoveEdge 在有向图里找到删除的那条边，使其变成树
	getRemoveEdge := func(edges [][]int) []int {
		initialize()
		for i := 0; i < len(edges); i++ { // 遍历所有的边
			if same(edges[i][0], edges[i][1]) { // 构成有向环了，就是要删除的边
				return edges[i]
			}
			join(edges[i][0], edges[i][1])
		}
		return []int{}
	}

	// isTreeAfterRemoveEdge 删一条边之后判断是不是树
	isTreeAfterRemoveEdge := func(edges [][]int, deleteEdge int) bool {
		initialize()
		for i := 0; i < len(edges); i++ {
			if i == deleteEdge {
				continue
			}
			if same(edges[i][0], edges[i][1]) { // 构成有向环了，一定不是树
				return false
			}
			join(edges[i][0], edges[i][1])
		}
		return true
	}

	inDegree := make([]int, len(father))
	for i := 0; i < len(edges); i++ {
		// 统计入度
		inDegree[edges[i][1]] += 1
	}
	// 记录入度为2的边（如果有的话就两条边）
	// 找入度为2的节点所对应的边，注意要倒序，因为优先返回最后出现在二维数组中的答案
	twoDegree := make([]int, 0)
	for i := len(edges) - 1; i >= 0; i-- {
		if inDegree[edges[i][1]] == 2 {
			twoDegree = append(twoDegree, i)
		}
	}

	// 处理图中情况1 和 情况2
	// 如果有入度为2的节点，那么一定是两条边里删一个，看删哪个可以构成树
	if len(twoDegree) > 0 {
		if isTreeAfterRemoveEdge(edges, twoDegree[0]) {
			return edges[twoDegree[0]]
		}
		return edges[twoDegree[1]]
	}

	// 处理图中情况3
	// 明确没有入度为2的情况，那么一定有有向环，找到构成环的边返回就可以了
	return getRemoveEdge(edges)
}

// @lc code=end
