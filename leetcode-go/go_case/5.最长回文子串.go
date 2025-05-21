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
// func longestPalindrome(s string) string {

// 	n := len(s)
// 	dp := make([][]bool, n)
// 	for i := range dp {
// 		dp[i] = make([]bool, n)
// 	}

// 	for i := 0; i < n; i++ {
// 		dp[i][i] = true
// 	}
// 	maxlen := 0
// 	res := [2]int{0, 0}
// 	for i := n - 1; i >= 0; i-- {
// 		for j := i; j < n; j++ {
// 			if s[i] == s[j] {
// 				if j-i <= 1 {
// 					dp[i][j] = true
// 				} else {
// 					if dp[i+1][j-1] {
// 						dp[i][j] = true
// 					} else {
// 						dp[i][j] = false
// 					}
// 				}
// 			} else {
// 				dp[i][j] = false
// 			}
// 			if dp[i][j] && (j-i+1 > maxlen) {
// 				maxlen = j - i + 1
// 				res = [2]int{i, j}
// 			}
// 		}
// 	}

// 	return s[res[0] : res[1]+1]

// }

// 中心扩展检测
// func longestPalindrome(s string) string {
// 	n := len(s)
// 	if n == 0 {
// 		return ""
// 	}

// 	start, end := 0, 0

// 	// 定义检查回文的DFS函数
// 	var expandAroundCenter func(left, right int) (int, int)
// 	expandAroundCenter = func(left, right int) (int, int) {
// 		for left >= 0 && right < n && s[left] == s[right] {
// 			left--
// 			right++
// 		}
// 		return left + 1, right - 1
// 	}

// 	for i := 0; i < n; i++ {
// 		// 奇数长度回文
// 		l1, r1 := expandAroundCenter(i, i)
// 		// 偶数长度回文
// 		l2, r2 := expandAroundCenter(i, i+1)

// 		// 更新最长回文
// 		if r1-l1 > end-start {
// 			start, end = l1, r1
// 		}
// 		if r2-l2 > end-start {
// 			start, end = l2, r2
// 		}
// 	}

// 	return s[start : end+1]
// }

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	// dp[i][j] 表示s[i..j]是否是回文
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLen := 0, 1

	// 所有长度为1的子串都是回文
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// 检查长度为2的子串
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLen = 2
		}
	}

	// 检查长度大于2的子串
	for length := 3; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			// 首尾字符相同，且中间子串是回文
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if length > maxLen {
					start = i
					maxLen = length
				}
			}
		}
	}

	return s[start : start+maxLen]
}

// @lc code=end

