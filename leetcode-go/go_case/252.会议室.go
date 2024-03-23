package go_case

import "sort"

/*
 * @lc app=leetcode.cn id=252 lang=golang
 *
 * [252] 会议室
 *
 * https://leetcode.cn/problems/meeting-rooms/description/
 *
 * algorithms
 * Easy (57.85%)
 * Likes:    146
 * Dislikes: 0
 * Total Accepted:    25K
 * Total Submissions: 43.2K
 * Testcase Example:  '[[0,30],[5,10],[15,20]]'
 *
 * 给定一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi]
 * ，请你判断一个人是否能够参加这里面的全部会议。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[0,30],[5,10],[15,20]]
 * 输出：false
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[7,10],[2,4]]
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * intervals[i].length == 2
 * 0 i i
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

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}

	sort.Sort(gaps(intervals))
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}

// @lc code=end
