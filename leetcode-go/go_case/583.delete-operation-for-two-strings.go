package go_case

/*
 * @lc app=leetcode id=583 lang=golang
 *
 * [583] Delete Operation for Two Strings
 *
 * https://leetcode.com/problems/delete-operation-for-two-strings/description/
 *
 * algorithms
 * Medium (61.51%)
 * Likes:    5684
 * Dislikes: 87
 * Total Accepted:    247.1K
 * Total Submissions: 401.6K
 * Testcase Example:  '"sea"\n"eat"'
 *
 * Given two strings word1 and word2, return the minimum number of steps
 * required to make word1 and word2 the same.
 *
 * In one step, you can delete exactly one character in either string.
 *
 *
 * Example 1:
 *
 *
 * Input: word1 = "sea", word2 = "eat"
 * Output: 2
 * Explanation: You need one step to make "sea" to "ea" and another step to
 * make "eat" to "ea".
 *
 *
 * Example 2:
 *
 *
 * Input: word1 = "leetcode", word2 = "etco"
 * Output: 4
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= word1.length, word2.length <= 500
 * word1 and word2 consist of only lowercase English letters.
 *
 *
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	//初始化
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+2)
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

// func minDistance(word1 string, word2 string) int {
// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
// 	t1 := len(word1)
// 	t2 := len(word2)
// 	dp := make([][]int, t1+1)
// 	for i := range dp {
// 		dp[i] = make([]int, t2+1)
// 	}

// 	for i := 1; i <= t1; i++ {
// 		for j := 1; j <= t2; j++ {
// 			if word1[i-1] == word2[j-1] {
// 				dp[i][j] = dp[i-1][j-1] + 1
// 			} else {
// 				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
// 			}
// 		}
// 	}
// 	return len(word1) + len(word2) - 2*dp[t1][t2]
// }

// @lc code=end
