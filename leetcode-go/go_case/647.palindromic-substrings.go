package go_case

/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 *
 * https://leetcode.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (70.08%)
 * Likes:    10651
 * Dislikes: 232
 * Total Accepted:    802.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '"abc"'
 *
 * Given a string s, return the number of palindromic substrings in it.
 *
 * A string is a palindrome when it reads the same backward as forward.
 *
 * A substring is a contiguous sequence of characters within the string.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "abc"
 * Output: 3
 * Explanation: Three palindromic strings: "a", "b", "c".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "aaa"
 * Output: 6
 * Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 1000
 * s consists of lowercase English letters.
 *
 *
 */

// @lc code=start
func countSubstrings(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					res++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}
	return res
}

// func countSubstrings(s string) int {
// 	extend := func(i, j int) int {
// 		res := 0
// 		for i >= 0 && j < len(s) && s[i] == s[j] {
// 			i --
// 			j ++
// 			res ++
// 		}
// 		return res
// 	}
// 	res := 0
// 	for i := 0; i < len(s); i++ {
// 		res += extend(i, i)  // 以i为中心
// 		res += extend(i, i+1) // 以i和i+1为中心
// 	}
// 	return res
// }

// @lc code=end
