/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 *
 * https://leetcode.cn/problems/maximum-length-of-repeated-subarray/description/
 *
 * algorithms
 * Medium (56.82%)
 * Likes:    1167
 * Dislikes: 0
 * Total Accepted:    298.4K
 * Total Submissions: 525.2K
 * Testcase Example:  '[1,2,3,2,1]\n[3,2,1,4,7]'
 *
 * 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
 * 输出：3
 * 解释：长度最长的公共子数组是 [3,2,1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length, nums2.length <= 1000
 * 0 <= nums1[i], nums2[i] <= 100
 *
 *
 */

// @lc code=start
// dp[i][j] 表示 以nums[i-1] 和 nums[j-1] 为结尾的 最长公共子序列
// func findLength(nums1 []int, nums2 []int) int {
// 	l1 := len(nums1)
// 	l2 := len(nums2)

// 	dp := make([][]int, l1+1)
// 	for i := range dp {
// 		dp[i] = make([]int, l2+1)
// 	}
// 	// dp[i][j] 表示 以nums[i-1] 和 nums[j-1] 为结尾的 最长公共子序列
// 	// 所以 当 nums1[i] == nums2[0]的时候
// 	// 或者 当 nums1[0] == nums2[j]的时候
// 	res := 0
// 	for i := 1; i <= l1; i++ {
// 		for j := 1; j <= l2; j++ {
// 			if nums1[i-1] == nums2[j-1] {
// 				dp[i][j] = dp[i-1][j-1] + 1
// 			}
// 			if dp[i][j] > res {
// 				res = dp[i][j]
// 			}
// 		}
// 	}
// 	return res
// }

// dp[i][j] 表示 以nums[i] 和 nums[j] 为结尾的 最长公共子序列
func findLength(nums1 []int, nums2 []int) int {
	l1 := len(nums1)
	l2 := len(nums2)

	dp := make([][]int, l1)
	for i := range dp {
		dp[i] = make([]int, l2)
	}
	// init
	for i := 0; i < l1; i++ {
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
		}
	}
	for j := 0; j < l2; j++ {
		if nums1[0] == nums2[j] {
			dp[0][j] = 1
		}
	}

	res := 0
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if nums1[i] == nums2[j] && i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

// @lc code=end

