package go_case

/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 *
 * https://leetcode.cn/problems/delete-operation-for-two-strings/description/
 *
 * algorithms
 * Medium (68.19%)
 * Likes:    727
 * Dislikes: 0
 * Total Accepted:    192.5K
 * Total Submissions: 282K
 * Testcase Example:  '"sea"\n"eat"'
 *
 * 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
 *
 * 每步 可以删除任意一个字符串中的一个字符。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: word1 = "sea", word2 = "eat"
 * 输出: 2
 * 解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
 *
 *
 * 示例  2:
 *
 *
 * 输入：word1 = "leetcode", word2 = "etco"
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 * 1 <= word1.length, word2.length <= 500
 * word1 和 word2 只包含小写英文字母
 *
 *
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	// dp[i][j] 以 word1[i-1] 为结尾的 word1[0:i-1]字符串 和 以word2[j-1] 为结尾的 word2[0:j-1]字符串
	// 删除任意一个字符至相同， 所需的最小步数。

	// 当 word1[i-1] == word2[j-1]时，
	// dp[i][j] = dp[i-1][j-1]

	// 当 word1[i-1] != word2[j-1]时，
	// 1.删除word1[i-1], dp[i-1][j]+1
	// 2.删除word2[j-1], dp[i-1][j]+1
	// 3.删除word1[i-1] 和 word2[j-1],  dp[i-1][j-1]+2
	// 三者当中的最小值
	// dp[i][j] = min(dp[i-1][j-1]+1, min(dp[i-1][j]+1, dp[i][j-1]+1))

	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	// init
	for i := 0; i <= l1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+2, min(dp[i-1][j]+1, dp[i][j-1]+1))
			}
		}
	}
	return dp[l1][l2]
}

// @lc code=end
