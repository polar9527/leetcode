package go_case

import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode.cn/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (63.43%)
 * Likes:    1202
 * Dislikes: 0
 * Total Accepted:    357.2K
 * Total Submissions: 563.2K
 * Testcase Example:  '[1,2,2]'
 *
 * 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。
 *
 * 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,2]
 * 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10
 * -10 <= nums[i] <= 10
 *
 *
 *
 *
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var track []int

	var backtracing func([]int, int, []int)
	backtracing = func(nums []int, startIndex int, track []int) {
		res = append(res, append([]int{}, track...))
		// if startIndex == len(nums) {
		//     return
		// }
		for i := startIndex; i < len(nums); i++ {
			if i > startIndex && nums[i] == nums[i-1] {
				continue
			}
			track = append(track, nums[i])
			backtracing(nums, i+1, track)
			track = track[:len(track)-1]
		}
	}

	sort.Ints(nums)
	backtracing(nums, 0, track)
	return res
}

// @lc code=end
