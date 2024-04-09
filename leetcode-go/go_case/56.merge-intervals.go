package go_case

import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (47.17%)
 * Likes:    21833
 * Dislikes: 760
 * Total Accepted:    2.3M
 * Total Submissions: 4.9M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given an array of intervals where intervals[i] = [starti, endi], merge all
 * overlapping intervals, and return an array of the non-overlapping intervals
 * that cover all the intervals in the input.
 *
 *
 * Example 1:
 *
 *
 * Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlap, merge them into
 * [1,6].
 *
 *
 * Example 2:
 *
 *
 * Input: intervals = [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^4
 *
 *
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if right < intervals[i][0] {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = max(right, intervals[i][1])
		}
	}
	res = append(res, []int{left, right}) // 将最后一个区间放入
	return res
}

// @lc code=end
