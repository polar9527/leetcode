/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (29.56%)
 * Likes:    2234
 * Dislikes: 0
 * Total Accepted:    276.9K
 * Total Submissions: 908K
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */
package main

// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					// l >= 2 && i + l < n
					// i+1 < n;  j-1 >= 0
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && len(ans) < l+1 {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

func longestPalindromeExpend(s string) string {
	le := len(s)
	if le == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < le; i++ {
		left1, right1 := expend(s, i, i)
		left2, right2 := expend(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expend(s string, start, end int) (int, int) {
	for ; start >= 0 && end < len(s) && s[start] == s[end]; start, end = start-1, end+1 {
	}
	return start + 1, end - 1
}

// @lc code=end
