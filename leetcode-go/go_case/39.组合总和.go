package go_case

import "slices"

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode.cn/problems/combination-sum/description/
 *
 * algorithms
 * Medium (72.44%)
 * Likes:    2745
 * Dislikes: 0
 * Total Accepted:    888.4K
 * Total Submissions: 1.2M
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target
 * 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
 *
 * candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
 *
 * 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：candidates = [2,3,6,7], target = 7
 * 输出：[[2,2,3],[7]]
 * 解释：
 * 2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
 * 7 也是一个候选， 7 = 7 。
 * 仅有这两种组合。
 *
 * 示例 2：
 *
 *
 * 输入: candidates = [2,3,5], target = 8
 * 输出: [[2,2,2,2],[2,3,3],[3,5]]
 *
 * 示例 3：
 *
 *
 * 输入: candidates = [2], target = 1
 * 输出: []
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= candidates.length <= 30
 * 2 <= candidates[i] <= 40
 * candidates 的所有元素 互不相同
 * 1 <= target <= 40
 *
 *
 */

// @lc code=start
// func combinationSum(candidates []int, target int) [][]int {
// 	var res [][]int
// 	var path []int

// 	var backtracing func(int, int)

// 	backtracing = func(startIndex, sum int) {
// 		if sum > target {
// 			return
// 		}

// 		if sum == target {
// 			res = append(res, append([]int{}, path...))
// 			return
// 		}

// 		for i := startIndex; i < len(candidates); i++ {

//				if sum+candidates[i] > target {
//					continue
//				}
//				path = append(path, candidates[i])
//				backtracing(i, sum+candidates[i])
//				path = path[:len(path)-1]
//			}
//		}
//		backtracing(0, 0)
//		return res
//	}

// 枚举所有可能的结果
// func combinationSum(candidates []int, target int) [][]int {

// 	var res [][]int
// 	var path []int
// 	var bt func(int, int)
// 	bt = func(start, sum int) {
// 		if sum > target {
// 			return
// 		}
// 		if sum == target {
// 			res = append(res, append([]int{}, path...))
// 			return
// 		}

// 		for i := start; i < len(candidates); i++ {
// 			path = append(path, candidates[i])
// 			bt(i, sum+candidates[i])
// 			path = path[:len(path)-1]
// 		}

// 	}
// 	bt(0, 0)
// 	return res
// }

// 枚举的基础上做剪枝优化
func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	var res [][]int
	var path []int
	var bt func(int, int)
	bt = func(start, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := start; i < len(candidates) && sum+candidates[i] <= target; i++ {
			path = append(path, candidates[i])
			bt(i, sum+candidates[i])
			path = path[:len(path)-1]
		}

	}
	bt(0, 0)
	return res
}

// @lc code=end
