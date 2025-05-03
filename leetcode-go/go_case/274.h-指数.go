package go_case

/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 *
 * https://leetcode.cn/problems/h-index/description/
 *
 * algorithms
 * Medium (46.27%)
 * Likes:    580
 * Dislikes: 0
 * Total Accepted:    282.6K
 * Total Submissions: 610K
 * Testcase Example:  '[3,0,6,1,5]'
 *
 * 给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。
 *
 * 根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h
 * 篇论文被引用次数大于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：citations = [3,0,6,1,5]
 * 输出：3
 * 解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
 * 由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。
 *
 * 示例 2：
 *
 *
 * 输入：citations = [1,3,1]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == citations.length
 * 1 <= n <= 5000
 * 0 <= citations[i] <= 1000
 *
 *
 */

// @lc code=start
// func hIndex(citations []int) int {
// 	// define
// 	// 1.研究者最多有 h 篇 论文 至少被引用了 h 次
// 	// 2.剩余的 n-h 篇 论文引用 次数少于 h 次
// 	// 排序法
// 	// 将 citation 降序排列
// 	// c[0] >= c[1] >= c[2] >= ... >= c[n-1]
// 	// 找到最大的 h, 使得 c[h-1] >= h
// 	// 由于是降序 c[0] >= c[1] >= c[2] >= ... >= c[h-1] >= h
// 	// 即 h 篇 论文引用的次数至少为 h
// 	// 剩余的 c[h], c[h+1], ... c[n-1] 都小于 h, 否则 h 可以更大
// 	n := len(citations)
// 	sort.Slice(citations, func(i, j int) bool {
// 		return citations[i] > citations[j]
// 	})
// 	h := 0
// 	for i := 0; i < n; i++ {
// 		if citations[i] >= i+1 {
// 			h = i + 1
// 		} else {
// 			break
// 		}
// 	}
// 	return h
// }

func hIndex(citations []int) int {
	n := len(citations)
	count := make([]int, n+1)

	for _, c := range citations {
		if c >= n {
			count[n]++
		} else {
			count[c]++
		}
	}
	total := 0
	for k := n; k >= 0; k-- {
		total += count[k]
		// 引用次数大于或等于 k 的 文章的总数 total
		if total >= k {
			return k
		}
	}
	return 0
}

// @lc code=end
