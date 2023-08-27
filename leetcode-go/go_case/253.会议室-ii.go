/*
 * @lc app=leetcode.cn id=253 lang=golang
 *
 * [253] 会议室 II
 *
 * https://leetcode.cn/problems/meeting-rooms-ii/description/
 *
 * algorithms
 * Medium (52.26%)
 * Likes:    548
 * Dislikes: 0
 * Total Accepted:    66.6K
 * Total Submissions: 127.4K
 * Testcase Example:  '[[0,30],[5,10],[15,20]]'
 *
 * 给你一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi]
 * ，返回 所需会议室的最小数量 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[0,30],[5,10],[15,20]]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[7,10],[2,4]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= intervals.length <= 10^4
 * 0 <= starti < endi <= 10^6
 *
 *
 */

// @lc code=start
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
func minMeetingRooms(intervals [][]int) int {
	var nums []int
	for _, v := range intervals {
		// 把 会议开始和会议结束时间，通过个位奇偶数来区分，同时放在一个组中排序
		nums = append(nums, v[0]*10+2)
		nums = append(nums, v[1]*10+1)
	}
	sort.Ints(nums)
	maxRoom := 0
	curNeedRoom := 0
	for _, v := range nums {
		if v%10 == 1 {
			curNeedRoom--
		} else {
			curNeedRoom++
		}
		if curNeedRoom > maxRoom {
			maxRoom = curNeedRoom
		}
	}
	return maxRoom
}

// 作者：珊璞桑
// 链接：https://leetcode.cn/problems/meeting-rooms-ii/solutions/241147/golang-shang-xia-che-fa-by-bloodborne/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// @lc code=end

// @lc code=end
