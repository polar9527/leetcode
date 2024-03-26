package go_case

/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 非递减子序列
 *
 * https://leetcode.cn/problems/non-decreasing-subsequences/description/
 *
 * algorithms
 * Medium (51.83%)
 * Likes:    776
 * Dislikes: 0
 * Total Accepted:    173.5K
 * Total Submissions: 334.8K
 * Testcase Example:  '[4,6,7,7]'
 *
 * 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
 *
 * 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,6,7,7]
 * 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,4,3,2,1]
 * 输出：[[4,4]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 15
 * -100 <= nums[i] <= 100
 *
 *
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {

	var res [][]int
	var path []int

	var backtracing func(int)
	backtracing = func(startIndex int) {
		if len(path) >= 2 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		used := make(map[int]bool)
		for i := startIndex; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || path[len(path)-1] <= nums[i] {
				used[nums[i]] = true
				path = append(path, nums[i])
				backtracing(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backtracing(0)
	return res
}

// @lc code=end
