package go_case

/*
 * @lc app=leetcode id=718 lang=golang
 *
 * [718] Maximum Length of Repeated Subarray
 *
 * https://leetcode.com/problems/maximum-length-of-repeated-subarray/description/
 *
 * algorithms
 * Medium (50.95%)
 * Likes:    6741
 * Dislikes: 166
 * Total Accepted:    297.8K
 * Total Submissions: 584.4K
 * Testcase Example:  '[1,2,3,2,1]\n[3,2,1,4,7]'
 *
 * Given two integer arrays nums1 and nums2, return the maximum length of a
 * subarray that appears in both arrays.
 *
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
 * Output: 3
 * Explanation: The repeated subarray with maximum length is [3,2,1].
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
 * Output: 5
 * Explanation: The repeated subarray with maximum length is [0,0,0,0,0].
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums1.length, nums2.length <= 1000
 * 0 <= nums1[i], nums2[i] <= 100
 *
 *
 */

// @lc code=start
func findLength(A []int, B []int) int {
	m, n := len(A), len(B)
	res := 0
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

// 滚动数组
// func findLength(nums1 []int, nums2 []int) int {
//     n, m, res := len(nums1), len(nums2), 0
//     dp := make([]int, m+1)
//     for i := 1; i <= n; i++ {
//         for j := m; j >= 1; j-- {
//             if nums1[i-1] == nums2[j-1] {
//                 dp[j] = dp[j-1] + 1
//             } else {
//                 dp[j] = 0  // 注意这里不相等要赋值为0，供下一层使用
//             }
//             res = max(res, dp[j])
//         }
//     }
//     return res
// }
// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

// @lc code=end
