package go_case

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 *
 * https://leetcode.cn/problems/word-break/description/
 *
 * algorithms
 * Medium (57.62%)
 * Likes:    2711
 * Dislikes: 0
 * Total Accepted:    759.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
 *
 * 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
 *
 *
 * 示例 2：
 *
 *
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
 * 注意，你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 *
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 300
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 20
 * s 和 wordDict[i] 仅由小写英文字母组成
 * wordDict 中的所有字符串 互不相同
 *
 *
 */

// @lc code=start
// backtrace
// func wordBreak(s string, wordDict []string) bool {
// 	wordSet := make(map[string]bool)
// 	for _, s := range wordDict {
// 		wordSet[s] = true
// 	}
// 	memory := make(map[int]bool)
// 	var bt func(int) bool
// 	bt = func(start int) bool {
// 		if start >= len(s) {
// 			return true
// 		}
// 		if _, ok := memory[start]; ok {
// 			return memory[start]
// 		}
// 		for i := start; i < len(s); i++ {
// 			substr := s[start : i+1]
// 			if wordSet[substr] {
// 				if bt(i + 1) {
// 					memory[i+1] = true
// 					return true
// 				}
// 			}
// 		}
// 		memory[start] = false
// 		return false
// 	}
// 	return bt(0)
// }

// dp
// func wordBreak(s string, wordDict []string) bool {

// 	// 背包位 s
// 	// s 中的 子串，且这个子串必须在 wordDict中才算物品
// 	// dp[i] 为 长度为i 的 字符串， 是否可以由 wordDict 中的字符串组成
// 	// 完全背包问题求排列
// 	// i > j
// 	// dp[i] = true , if [j, i) 在 wordDict， 且 dp[j] = true

// 	// 求排列 ，先背包后物品

// 	wordSet := make(map[string]bool)
// 	for _, s := range wordDict {
// 		wordSet[s] = true
// 	}
// 	dp := make([]bool, len(s)+1)
// 	dp[0] = true
// 	for i := 1; i <= len(s); i++ {
// 		for j := 0; j < i; j++ {
// 			substr := s[j:i]
// 			if _, ok := wordSet[substr]; ok {
// 				if dp[j] {
// 					dp[i] = true
// 				}
// 			}
// 		}
// 	}
// 	return dp[len(s)]
// }

// dfs
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	maxLen := 0
	wordSet := make(map[string]bool)
	for _, s := range wordDict {
		wordSet[s] = true
		maxLen = max(maxLen, len(s))
	}

	cache := make([]int8, n+1)
	for i := range cache {
		cache[i] = -1
	}

	var dfs func(int) int8
	dfs = func(i int) (res int8) {
		if i == 0 {
			return 1
		}

		if cache[i] != -1 {
			return cache[i]
		}
		defer func() {
			cache[i] = res
		}()
		// 0, i-maxLen, i-1, i
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if wordSet[s[j:i]] && dfs(j) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(n) == 1
}

// DP
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	maxLen := 0
	wordSet := make(map[string]bool)
	for _, s := range wordDict {
		wordSet[s] = true
		maxLen = max(maxLen, len(s))
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// @lc code=end
