package go_case

import "slices"

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 *
 * https://leetcode.cn/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (57.02%)
 * Likes:    3935
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.1M
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
 *
 * 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7]
 * 的子序列。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,9,2,5,3,7,101,18]
 * 输出：4
 * 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,0,3,2,3]
 * 输出：4
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [7,7,7,7,7,7,7]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2500
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
 *
 *
 */

// @lc code=start
// func lengthOfLIS(nums []int) int {

// 	// dp[i] 以nums[i] 为结尾的 最长严格递增子序列长度
// 	// 如果 nums[i] > nums[j]，那么
// 	// dp[i] = dp[j] + 1

// 	l := len(nums)

// 	if l <= 1 {
// 		return l
// 	}
// 	// if l == 1 {
// 	// 	return 1
// 	// }
// 	dp := make([]int, l)
// 	for i := range dp {
// 		dp[i] = 1
// 	}
// 	res := 0
// 	for i := 1; i < l; i++ {
// 		for j := 0; j < i; j++ {
// 			if nums[j] < nums[i] {
// 				dp[i] = max(dp[i], dp[j]+1)
// 			}
// 		}
// 		if dp[i] > res {
// 			res = dp[i]
// 		}
// 	}

// 	return res
// }

// DFS
// func lengthOfLIS(nums []int) int {
// 	n := len(nums)

// 	cache := make([]int, n)
// 	for i := range cache {
// 		cache[i] = -1
// 	}

// 	// 以 nums[i] 为结尾的最长递增子序列长度
// 	var dfs func(int) int
// 	dfs = func(i int) (res int) {
// 		if cache[i] != -1 {
// 			return cache[i]
// 		}
// 		defer func() {
// 			cache[i] = res
// 		}()
// 		for j := range nums[:i] {
// 			if nums[j] < nums[i] {
// 				res = max(res, dfs(j))
// 			}
// 		}
// 		return res + 1

// 	}
// 	maxRes := 0
// 	for i := 0; i < n; i++ {
// 		maxRes = max(maxRes, dfs(i))
// 	}
// 	return maxRes
// }

// DP
func lengthOfLIS(nums []int) int {
	n := len(nums)

	dp := make([]int, n)

	for i, x := range nums {
		for j, y := range nums[:i] {
			if y < x {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i]++
	}
	return slices.Max(dp)
}

// 二分 + 贪心
// g[i] 长度为i+1的严格递增子序列末尾的最小值
// func lengthOfLIS(nums []int) int {
// 	g := []int{}
// 	for _, x := range nums {
// 		j := sort.SearchInts(g, x)
// 		if j == len(g) { // >=x 的 g[j] 不存在
// 			g = append(g, x)
// 		} else {
// 			g[j] = x
// 		}
// 	}
// 	return len(g)
// }

// @lc code=end
