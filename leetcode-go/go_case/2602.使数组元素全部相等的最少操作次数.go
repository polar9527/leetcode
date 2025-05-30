package go_case

import (
	"slices"
	"sort"
)

/*
 * @lc app=leetcode.cn id=2602 lang=golang
 *
 * [2602] 使数组元素全部相等的最少操作次数
 *
 * https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/description/
 *
 * algorithms
 * Medium (39.12%)
 * Likes:    61
 * Dislikes: 0
 * Total Accepted:    12.3K
 * Total Submissions: 31.1K
 * Testcase Example:  '[3,1,6,8]\n[1,5]'
 *
 * 给你一个正整数数组 nums 。
 *
 * 同时给你一个长度为 m 的整数数组 queries 。第 i 个查询中，你需要将 nums 中所有元素变成 queries[i] 。你可以执行以下操作
 * 任意 次：
 *
 *
 * 将数组里一个元素 增大 或者 减小 1 。
 *
 *
 * 请你返回一个长度为 m 的数组 answer ，其中 answer[i]是将 nums 中所有元素变成 queries[i] 的 最少 操作次数。
 *
 * 注意，每次查询后，数组变回最开始的值。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [3,1,6,8], queries = [1,5]
 * 输出：[14,10]
 * 解释：第一个查询，我们可以执行以下操作：
 * - 将 nums[0] 减小 2 次，nums = [1,1,6,8] 。
 * - 将 nums[2] 减小 5 次，nums = [1,1,1,8] 。
 * - 将 nums[3] 减小 7 次，nums = [1,1,1,1] 。
 * 第一个查询的总操作次数为 2 + 5 + 7 = 14 。
 * 第二个查询，我们可以执行以下操作：
 * - 将 nums[0] 增大 2 次，nums = [5,1,6,8] 。
 * - 将 nums[1] 增大 4 次，nums = [5,5,6,8] 。
 * - 将 nums[2] 减小 1 次，nums = [5,5,5,8] 。
 * - 将 nums[3] 减小 3 次，nums = [5,5,5,5] 。
 * 第二个查询的总操作次数为 2 + 4 + 1 + 3 = 10 。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [2,9,6,3], queries = [10]
 * 输出：[20]
 * 解释：我们可以将数组中所有元素都增大到 10 ，总操作次数为 8 + 1 + 4 + 7 = 20 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * m == queries.length
 * 1 <= n, m <= 10^5
 * 1 <= nums[i], queries[i] <= 10^9
 *
 *
 */

// @lc code=start
func minOperations(nums []int, queries []int) []int64 {

	// 先排序 nums 方便求出前缀和
	// sort.Slice(nums, func(i, j int) bool {
	// 	if nums[i] < nums[j] {
	// 		return true
	// 	}
	// 	return false
	// })
	slices.Sort(nums)
	n := len(nums)
	prefixSum := make([]int, n+1)
	prefixSum[0] = 0
	// for i := 1; i <= n; i++ {
	// 	prefixSum[i] += prefixSum[i-1] + nums[i-1]
	// }
	for i := range nums {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	res := make([]int64, len(queries))
	// sort.SearchInts(nums)
	for i, q := range queries {
		// q 可以插入到 nums[j] 的位置
		// 所有 q 前面 有 [0,j-1] 个序号，共 j 个数字
		j := sort.SearchInts(nums, q)
		left := q*j - prefixSum[j]
		// [j, n] 之间的 值的总和，
		right := prefixSum[n] - prefixSum[j] - q*(n-j)
		res[i] = int64(left) + int64(right)
	}

	return res
}

// @lc code=end
