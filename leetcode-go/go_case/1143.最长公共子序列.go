/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 *
 * https://leetcode.cn/problems/longest-common-subsequence/description/
 *
 * algorithms
 * Medium (66.28%)
 * Likes:    1724
 * Dislikes: 0
 * Total Accepted:    603.2K
 * Total Submissions: 908.7K
 * Testcase Example:  '"abcde"\n"ace"'
 *
 * 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
 *
 * 一个字符串的 子序列
 * 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
 *
 *
 * 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
 *
 *
 * 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：text1 = "abcde", text2 = "ace"
 * 输出：3
 * 解释：最长公共子序列是 "ace" ，它的长度为 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：text1 = "abc", text2 = "abc"
 * 输出：3
 * 解释：最长公共子序列是 "abc" ，它的长度为 3 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：text1 = "abc", text2 = "def"
 * 输出：0
 * 解释：两个字符串没有公共子序列，返回 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * text1 和 text2 仅由小写英文字符组成。
 *
 *
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	t1 := len(text1)
	t2 := len(text2)

	dp := make([][]int, t1+1)
	for i := range dp {
		dp[i] = make([]int, t2+1)
	}
	// dp[i][j] text1[0:i-1]  和 text[0:j-1] 的最长公共子序列
	for i := 1; i <= t1; i++ {
		for j := 1; j <= t2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 不考虑 text1[i-1] 或者 text2[j-1]
				// text1[0:i-2]  和 text[0:j-1] 的最长公共子序列
				// text1[0:i-1]  和 text[0:j-2] 的最长公共子序列
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[t1][t2]
}

// @lc code=end

