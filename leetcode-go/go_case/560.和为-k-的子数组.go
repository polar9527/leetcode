package go_case

/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (44.72%)
 * Likes:    1993
 * Dislikes: 0
 * Total Accepted:    333.7K
 * Total Submissions: 742.2K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,1], k = 2
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3], k = 3
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -1000 <= nums[i] <= 1000
 * -10^7 <= k <= 10^7
 *
 *
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	n := len(nums)
	count, preSum := 0, 0
	m := map[int]int{}
	// preSum 与 k 相等，那么前缀和 preSum 自己所代表的 子数组也要 计算在内
	m[0] = 1
	for i := 0; i < n; i++ {
		preSum += nums[i]
		if _, ok := m[preSum-k]; ok {
			// 此处 是计算 preSum[i]和preSum[j] 之间的 前缀和差值，也就是 子数组之和k,
			// 而  m[preSum-k] 的值的含义是，所有满足 preSum-k 的前缀和的个数，
			// 表示此处可能有 j1,j2,...,jN ，一共N个已经存在字典m中的前缀和满足当前 与preSum-k相等的条件
			count += m[preSum-k]
		}
		// 而每次都要往字典的这个键值m[preSum]上累加，因为后续的 某个前缀和 需要在这个 累加值上 计算 满足条件
		// 的子数组的个数
		m[preSum] += 1
	}

	return count
}

// @lc code=end
