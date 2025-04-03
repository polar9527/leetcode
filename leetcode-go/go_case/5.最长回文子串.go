/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode.cn/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (38.44%)
 * Likes:    7583
 * Dislikes: 0
 * Total Accepted:    1.9M
 * Total Submissions: 5M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的 回文 子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由数字和英文字母组成
 *
 *
 */

// @lc code=start
// dp
func longestPalindrome(s string) string {

	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	maxlen := 0
	res := [2]int{0, 0}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else {
					if dp[i+1][j-1] {
						dp[i][j] = true
					} else {
						dp[i][j] = false
					}
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] && (j-i+1 > maxlen) {
				maxlen = j - i + 1
				res = [2]int{i, j}
			}
		}
	}

	return s[res[0] : res[1]+1]

}

// @lc code=end

