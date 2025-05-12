/*
 * @lc app=leetcode.cn id=2246 lang=golang
 *
 * [2246] 相邻字符不同的最长路径
 *
 * https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/description/
 *
 * algorithms
 * Hard (50.91%)
 * Likes:    98
 * Dislikes: 0
 * Total Accepted:    13.6K
 * Total Submissions: 26.5K
 * Testcase Example:  '[-1,0,0,1,1,2]\n"abacbe"'
 *
 * 给你一棵 树（即一个连通、无向、无环图），根节点是节点 0 ，这棵树由编号从 0 到 n - 1 的 n 个节点组成。用下标从 0 开始、长度为 n
 * 的数组 parent 来表示这棵树，其中 parent[i] 是节点 i 的父节点，由于节点 0 是根节点，所以 parent[0] == -1 。
 *
 * 另给你一个字符串 s ，长度也是 n ，其中 s[i] 表示分配给节点 i 的字符。
 *
 * 请你找出路径上任意一对相邻节点都没有分配到相同字符的 最长路径 ，并返回该路径的长度。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：parent = [-1,0,0,1,1,2], s = "abacbe"
 * 输出：3
 * 解释：任意一对相邻节点字符都不同的最长路径是：0 -> 1 -> 3 。该路径的长度是 3 ，所以返回 3 。
 * 可以证明不存在满足上述条件且比 3 更长的路径。
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：parent = [-1,0,0,0], s = "aabc"
 * 输出：3
 * 解释：任意一对相邻节点字符都不同的最长路径是：2 -> 0 -> 3 。该路径的长度为 3 ，所以返回 3 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == parent.length == s.length
 * 1 <= n <= 10^5
 * 对所有 i >= 1 ，0 <= parent[i] <= n - 1 均成立
 * parent[0] == -1
 * parent 表示一棵有效的树
 * s 仅由小写英文字母组成
 *
 *
 */

// @lc code=start
func longestPath(parent []int, s string) int {
	// 计算节点数量
	n := len(parent)
	// 计算每个节点的子节点
	children := make([][]int, n)
	// i == 0 时， parent[0] == -1 没有父节点，不用计算
	for i := 1; i < n; i++ {
		p := parent[i]
		children[p] = append(children[p], i)
	}

	var res int
	// 返回从 node 出发往下的最长链表
	var dfs func(int) int
	dfs = func(node int) int {
		var maxLen int
		// 当 node 是一个叶子节点，children[node]为空数组
		//
		// 直接返回maxLen的默认值 0
		// 遍历树的每个子节点
		for _, child := range children[node] {
			// 经过 child 这个子节点的最大链长
			length := dfs(child) + 1
			// 求解最长路径(边)
			if s[child] != s[node] {
				res = max(res, length+maxLen)
				maxLen = max(maxLen, length)
			}
		}
		return maxLen
	}
	dfs(0)
	// 求解最长路径(节点)
	return res + 1
}

// @lc code=end

