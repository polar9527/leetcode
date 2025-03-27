/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 *
 * https://leetcode.cn/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (53.20%)
 * Likes:    2276
 * Dislikes: 0
 * Total Accepted:    712.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,5,11,5]'
 *
 * 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,5,11,5]
 * 输出：true
 * 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,5]
 * 输出：false
 * 解释：数组不能分割成两个元素和相等的子集。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */

// @lc code=
// 2D array DP
// func canPartition(nums []int) bool {
// 	max := func(x, y int) int {
// 		if x > y {
// 			return x
// 		}
// 		return y
// 	}

// 	l := len(nums)

// 	sum := 0
// 	for _, n := range nums {
// 		sum += n
// 	}

// 	if sum%2 == 1 {
// 		return false
// 	}
// 	target := sum / 2

// 	// d[i][j]
// 	// nums [0, i] 的元素中 对于 背包容量j 的 最大值,
// 	// 这个最大值的上限就是j,因为 容量 和 价值 在这里是等价的。
// 	// 所以装满了背包j 就能得到最大值 ，既为 j ( dp[i][j] == j)
// 	dp := make([][]int, l)
// 	for i := 0; i < l; i++ {
// 		dp[i] = make([]int, target+1)
// 	}

// 	// init
// 	// for i := 0; i < l; i++ {
// 	// 	dp[i][0] = 0
// 	// }
// 	for j := nums[0]; j <= target; j++ {
// 		dp[0][j] = nums[0]
// 	}

// 	for i := 1; i < l; i++ {
// 		for j := 0; j <= target; j++ {
// 			if j-nums[i] < 0 {
// 				dp[i][j] = dp[i-1][j]
// 			} else {
// 				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
// 			}
// 		}

// 	}

// 	if dp[l-1][target] == target {
// 		return true
// 	}
// 	return false
// }

// 1D array DP
func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2

	dp := make([]int, target+1)
	// dp[0] = 0

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if dp[target] == target {
		return true
	}
	return false
}

// @lc code=end

