/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {

    wordSet := make(map[string]bool)
    maxStringLen := -1
    for _, s := range wordDict {
        wordSet[s] = true
        if maxStringLen < len(s) {
            maxStringLen = len(s)
        }
    }

    dp := make([]bool,len(s)+1)
    dp[0] = true
    // dp[i]=dp[j] && check(s[j..i−1])
    // 当j=0 时， dp[i]=dp[0] && check(s[0..i−1]), 只要s[0:i]在字典中，dp[i]必须为true,所以反推
    // dp[0]=true
    for i := 1; i <= len(s); i++ {
        // dp[i] = false
        for j := i - 1; j >= 0; j-- {
            // 当 s[j:i] 字符串的长度已经大于 字典中最长字符串的长度
            // 那么就可以直接结束，而不用去查字典了，因为肯定查不到
            if i-j > maxStringLen {
                break
            }
            if dp[j] && wordSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[len(s)]
}
// @lc code=end

