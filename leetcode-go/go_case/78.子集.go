package go_case

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode.cn/problems/subsets/description/
 *
 * algorithms
 * Medium (81.14%)
 * Likes:    2265
 * Dislikes: 0
 * Total Accepted:    753.5K
 * Total Submissions: 927.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
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
 * nums 中的所有元素 互不相同
 *
 *
 */

// @lc code=start
// func subsets(nums []int) [][]int {
// 	var res [][]int
// 	var backtracing func([]int, int, []int)
// 	backtracing = func(nums []int, startindex int, path []int) {
// 		res = append(res, append([]int{}, path...))
// 		// if (startindex >= len(nums)){
// 		// 	return
// 		// }
// 		for i := startindex; i < len(nums); i++ {
// 			path = append(path, nums[i])
// 			backtracing(nums, i+1, path)
// 			path = path[:len(path)-1]
// 		}
// 	}
// 	path := []int{}
// 	backtracing(nums, 0, path)
// 	return res

// }

func subsets(nums []int) [][]int {
	l := len(nums)
	var res [][]int
	var path []int
	var bt func(int)
	bt = func(start int) {
		res = append(res, append([]int{}, path...))

		// if start >= l {
		// 	return
		// }

		for i := start; i < l; i++ {
			path = append(path, nums[i])
			bt(i + 1)
			path = path[:len(path)-1]
		}
	}

	bt(0)
	return res
}

// @lc code=end
