/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 *
 * https://leetcode.cn/problems/distinct-subsequences/description/
 *
 * algorithms
 * Hard (53.07%)
 * Likes:    1334
 * Dislikes: 0
 * Total Accepted:    218.7K
 * Total Submissions: 411.6K
 * Testcase Example:  '"rabbbit"\n"rabbit"'
 *
 * 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数，结果需要对 10^9 + 7 取模。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "rabbbit", t = "rabbit"
 * 输出：3
 * 解释：
 * 如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
 * rabbbit
 * rabbbit
 * rabbbit
 *
 * 示例 2：
 *
 *
 * 输入：s = "babgbag", t = "bag"
 * 输出：5
 * 解释：
 * 如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
 * babgbag
 * babgbag
 * babgbag
 * babgbag
 * babgbag
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length, t.length <= 1000
 * s 和 t 由英文字母组成
 *
 *
 */

// @lc code=start
func numDistinct(s string, t string) int {
	ls := len(s)
	lt := len(t)
	// dp[i][j] 以 s[i-1]结尾的s的 子序列中 出现 以 t[j-1] 结尾 的t 的个数

	// 如果 s[i] == t[j],
	// 以s[i-1]为结尾来匹配t, 不以s[i-1]为结尾来匹配t
	// dp[i][j] = dp[i-1][j-1] + d[i-1][j]

	// 如果 s[i] != t[j],
	// dp[i][j] = dp[i-1][j]
	dp := make([][]int, ls+1)
	for i := range dp {
		dp[i] = make([]int, lt+1)
	}

	// init
	// dp[i][0], t 为 空字符串， s[i-1]结尾的s的 子序列中 出现 以 空字符串t 的个数为1
	// dp[0][j], s 为 空字符串， 空字符串 出现 以 t[j-1] 结尾 的t 的个数 为0
	for i := 0; i <= ls; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= ls; i++ {
		for j := 1; j <= lt; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[ls][lt]
}

// @lc code=end

