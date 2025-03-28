package go_case

/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 *
 * https://leetcode.cn/problems/ones-and-zeroes/description/
 *
 * algorithms
 * Medium (67.08%)
 * Likes:    1231
 * Dislikes: 0
 * Total Accepted:    255.2K
 * Total Submissions: 380.4K
 * Testcase Example:  '["10","0001","111001","1","0"]\n5\n3'
 *
 * 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
 *
 *
 * 请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
 *
 * 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
 * 输出：4
 * 解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
 * 其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1
 * ，大于 n 的值 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：strs = ["10", "0", "1"], m = 1, n = 1
 * 输出：2
 * 解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strs.length <= 600
 * 1 <= strs[i].length <= 100
 * strs[i] 仅由 '0' 和 '1' 组成
 * 1 <= m, n <= 100
 *
 *
 */

// @lc code=start
// 1D array
// func findMaxForm(strs []string, m int, n int) int {
// 	get := func(s string) (ones, zeroes int) {
// 		for _, c := range s {
// 			if c == '0' {
// 				zeroes++
// 			} else {
// 				ones++
// 			}
// 		}
// 		return
// 	}

// 	dp := make([][]int, m+1)
// 	for i, _ := range dp {
// 		dp[i] = make([]int, n+1)
// 	}

// 	// init
// 	// dp[0][0] = 0

// 	for _, s := range strs {
// 		n1, n0 := get(s)
// 		// m -> 1
// 		// n -> 0
// 		for i := m; i >= n0; i-- {
// 			for j := n; j >= n1; j-- {
// 				dp[i][j] = max(dp[i][j], dp[i-n0][j-n1]+1)
// 			}
// 		}
// 	}
// 	return dp[m][n]
// }

// 2D array
func findMaxForm(strs []string, m int, n int) int {
	get := func(s string) (ones, zeroes int) {
		for _, c := range s {
			if c == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		return
	}

	dp := make([][][]int, len(strs))
	for i, _ := range dp {
		dp[i] = make([][]int, m+1)
		for j, _ := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	// init
	n1, n0 := get(strs[0])
	for i := n0; i <= m; i++ {
		for j := n1; j <= n; j++ {
			dp[0][i][j] = 1
		}
	}

	for i := 1; i < len(strs); i++ {
		n1, n0 := get(strs[i])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if j-n0 < 0 || k-n1 < 0 {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-n0][k-n1]+1)
				}
			}
		}
	}
	return dp[len(strs)-1][m][n]
}

// @lc code=end
