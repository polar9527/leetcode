/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
package main

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	length := len(s)
	dp := make([][]int, length)

	max := 1
	start := 0

	for i := 0; i < length; i++ {
		if dp[i] == nil {
			dp[i] = make([]int, length)
		}
		dp[i][i] = 1
		if i < length-1 && s[i] == s[i+1] {
			dp[i][i+1] = 1
			start = i
			max = 2
		}
	}

	for l := 3; l <= length; l++ {
		for i := 0; (i + l - 1) < length; i++ {
			j := l + i - 1
			if s[i] == s[j] && dp[i+1][j-1] == 1 {
				dp[i][j] = 1
				start = i
				max = l
			}
		}
	}

	return s[start : start+max]
}

// @lc code=end
