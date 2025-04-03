package go_case

/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 *
 * https://leetcode.cn/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (66.88%)
 * Likes:    1208
 * Dislikes: 0
 * Total Accepted:    288.5K
 * Total Submissions: 431.4K
 * Testcase Example:  '"abc"'
 *
 * 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
 *
 * 回文字符串 是正着读和倒过来读一样的字符串。
 *
 * 子字符串 是字符串中的由连续字符组成的一个序列。
 *
 * 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abc"
 * 输出：3
 * 解释：三个回文子串: "a", "b", "c"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "aaa"
 * 输出：6
 * 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 由小写英文字母组成
 *
 *
 */

// @lc code=start
// brutal
// func countSubstrings(s string) int {
// 	n := len(s)
// 	ans := 0
// 	for i := 0; i < 2*n-1; i++ {
// 		l, r := i/2, i/2+i%2
// 		for l >= 0 && r < n && s[l] == s[r] {
// 			l--
// 			r++
// 			ans++
// 		}
// 	}
// 	return ans
// }

// dp
func countSubstrings(s string) int {
	n := len(s)

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// dp含义
	// dp[i][j]  [i, j] 左闭右闭的区间内 是否是回文字符串

	// 递推公式
	// 当 s[i] == s[j] 时
	// i == j, 或者 i-j == 1, 说明是回文字符串
	// 当 i - j > 1 时， 要根据 dp[i+1][j-1] 来判断

	// 初始化
	// dp[i][0] = 0
	// 遍历顺序
	// 打印遍历数组

	res := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					res++
					dp[i][j] = true
				} else {
					if dp[i+1][j-1] {
						res++
						dp[i][j] = true
					}
				}
			}
		}
	}
	return res
}

// @lc code=end
