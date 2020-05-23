/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
package main

func longestPalindrome(s string) string {
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

// func longestPalindrome(s string) string {
// 	le := len(s)
// 	if le == 0 {
// 		return ""
// 	}
// 	ans := ""
// 	dp := make([][]int, le)
// 	for i := range dp {
// 		dp[i] = make([]int, le)
// 	}
// 	// [i:j] = dp[i][j]
// 	// 0 <= i <= j < le
// 	for l := 0; l < le; l++ {
// 		for i := 0; i + l < le; i++ {
// 			j := i + l
// 			if l == 0 {
// 				dp[i][j] = 1
// 			} else if l == 1 {
// 				if s[i] == s[j] {
// 					dp[i][j] = 1
// 				}
// 			} else {
// 				// j - i >= 2
// 				// j - 1 >= i - 1
// 				if s[i] == s[j] {
// 					dp[i][j] = dp[i+1][j-1]
// 				}
// 			}
// 			if dp[i][j] == 1 && l + 1 > len(ans) {
// 				ans = s[i:j+1]
// 			}
// 		}
// 	}
//
// 	return ans
// }

// func longestPalindrome(s string) string {
// 	if len(s) == 0 || len(s) == 1 {
// 		return s
// 	}
//
// 	length := len(s)
// 	dp := make([][]int, length)
//
// 	max := 1
// 	start := 0
//
// 	for i := 0; i < length; i++ {
// 		if dp[i] == nil {
// 			dp[i] = make([]int, length)
// 		}
// 		dp[i][i] = 1
// 		if i < length-1 && s[i] == s[i+1] {
// 			dp[i][i+1] = 1
// 			start = i
// 			max = 2
// 		}
// 	}
//
// 	for l := 3; l <= length; l++ {
// 		for i := 0; (i + l - 1) < length; i++ {
// 			j := l + i - 1
// 			if s[i] == s[j] && dp[i+1][j-1] == 1 {
// 				dp[i][j] = 1
// 				start = i
// 				max = l
// 			}
// 		}
// 	}
//
// 	return s[start : start+max]
// }

// @lc code=end
