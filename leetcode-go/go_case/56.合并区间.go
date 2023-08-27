package go_case

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode.cn/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (49.43%)
 * Likes:    2059
 * Dislikes: 0
 * Total Accepted:    692.7K
 * Total Submissions: 1.4M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]
 * 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 输出：[[1,6],[8,10],[15,18]]
 * 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[1,4],[4,5]]
 * 输出：[[1,5]]
 * 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^4
 *
 *
 */

// @lc code=start
type gaps [][]int

func (g gaps) Len() int {
	return len(g)
}

func (g gaps) Less(i, j int) bool {
	return g[i][0] < g[j][0]
}

func (g gaps) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	var ans [][]int
	sort.Sort(gaps(intervals))
	for i := 0; i < len(intervals); i++ {
		if len(ans) == 0 || ans[len(ans)-1][1] < intervals[i][0] {
			ans = append(ans, intervals[i])
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], intervals[i][1])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
