package go_case

import "sort"

/*
 * @lc app=leetcode id=435 lang=golang
 *
 * [435] Non-overlapping Intervals
 *
 * https://leetcode.com/problems/non-overlapping-intervals/description/
 *
 * algorithms
 * Medium (53.04%)
 * Likes:    7964
 * Dislikes: 214
 * Total Accepted:    548K
 * Total Submissions: 1M
 * Testcase Example:  '[[1,2],[2,3],[3,4],[1,3]]'
 *
 * Given an array of intervals intervals where intervals[i] = [starti, endi],
 * return the minimum number of intervals you need to remove to make the rest
 * of the intervals non-overlapping.
 *
 *
 * Example 1:
 *
 *
 * Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
 * Output: 1
 * Explanation: [1,3] can be removed and the rest of the intervals are
 * non-overlapping.
 *
 *
 * Example 2:
 *
 *
 * Input: intervals = [[1,2],[1,2],[1,2]]
 * Output: 2
 * Explanation: You need to remove two [1,2] to make the rest of the intervals
 * non-overlapping.
 *
 *
 * Example 3:
 *
 *
 * Input: intervals = [[1,2],[2,3]]
 * Output: 0
 * Explanation: You don't need to remove any of the intervals since they're
 * already non-overlapping.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= intervals.length <= 10^5
 * intervals[i].length == 2
 * -5 * 10^4 <= starti < endi <= 5 * 10^4
 *
 *
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	res := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i-1][1] {
			res++
		} else {
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
		}
	}
	return len(intervals) - res
}

// @lc code=end
