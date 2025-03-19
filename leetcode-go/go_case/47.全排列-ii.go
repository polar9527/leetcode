package go_case

import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode.cn/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (65.69%)
 * Likes:    1555
 * Dislikes: 0
 * Total Accepted:    536.5K
 * Total Submissions: 816.8K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2]
 * 输出：
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 8
 * -10 <= nums[i] <= 10
 *
 *
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	var res [][]int
	var path []int
	used := make([]bool, l)
	var bt func()
	bt = func() {
		if len(path) == l {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < l; i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] {
				continue
			}

			if !used[i] {
				used[i] = true
				path = append(path, nums[i])
				bt()
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	bt()
	return res
}

// @lc code=end
