/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode.cn/problems/edit-distance/description/
 *
 * algorithms
 * Medium (63.47%)
 * Likes:    3645
 * Dislikes: 0
 * Total Accepted:    632K
 * Total Submissions: 994.7K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
 *
 * 你可以对一个单词进行如下三种操作：
 *
 *
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 *
 *
 * 示例 2：
 *
 *
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= word1.length, word2.length <= 500
 * word1 和 word2 由小写英文字母组成
 *
 *
 */

// @lc code=start
// func minDistance(word1 string, word2 string) int {
// 	l1 := len(word1)
// 	l2 := len(word2)
// 	// dp[i][j] 以word[i-1]结尾的 word1 转换成  以word2[j-1]为结尾的 word2 所使用的最少操作数

// 	dp := make([][]int, l1+1)
// 	for i := range dp {
// 		dp[i] = make([]int, l2+1)
// 	}

// 	for i := 0; i <= l1; i++ {
// 		dp[i][0] = i
// 	}
// 	for j := 0; j <= l2; j++ {
// 		dp[0][j] = j
// 	}

// 	for i := 1; i <= l1; i++ {
// 		for j := 1; j <= l2; j++ {
// 			if word1[i-1] == word2[j-1] {
// 				dp[i][j] = dp[i-1][j-1]
// 			} else {
// 				// 增加，删除，替换
// 				// 删除word1[i-1], dp[i-1][j] +1
// 				// 删除word2[j-1], dp[i][j-1] +1
// 				// 替换word1或者word2, dp[i-1][j-1] +1
// 				dp[i][j] = min(dp[i-1][j-1]+1, min(dp[i-1][j]+1, dp[i][j-1]+1))
// 			}
// 		}
// 	}
// 	return dp[l1][l2]
// }

// func minDistance(word1 string, word2 string) int {

// 	wl1 := len(word1)
// 	wl2 := len(word2)

// 	cache := make([][]int, wl1)
// 	for i := range cache {
// 		cache[i] = make([]int, wl2)
// 		for j := range cache[i] {
// 			cache[i][j] = -1
// 		}
// 	}

// 	var dfs func(int, int) int
// 	dfs = func(i, j int) (res int) {
// 		// i, j 是指 word1, word2 数组下标
// 		if i < 0 {
// 			return j + 1
// 		}
// 		if j < 0 {
// 			return i + 1
// 		}

// 		if cache[i][j] != -1 {
// 			return cache[i][j]
// 		}
// 		defer func() {
// 			cache[i][j] = res
// 		}()

// 		if word1[i] == word2[j] {
// 			return dfs(i-1, j-1)
// 		}
// 		// 删除 s[i]
// 		// 在s中插入 t[j], 等价于 删除 t[j],
// 		// 替换 将s[i] 替换为 t[j]
// 		return min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1
// 	}
// 	return dfs(wl1-1, wl2-1)
// }

func minDistance(word1 string, word2 string) int {

	wl1 := len(word1)
	wl2 := len(word2)
	dp := make([][]int, wl1+1)
	for i := range dp {
		dp[i] = make([]int, wl2+1)
	}

	for j := range wl2 {
		// j -> [0, wl2] 共 wl2+1 个数字，其中 [1:] 这 wl2 个数字 初始化为对应的下标
		dp[0][j+1] = j + 1
	}

	for i, w1 := range word1 {
		dp[i+1][0] = i + 1
		for j, w2 := range word2 {
			if w1 == w2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j], dp[i][j]) + 1
			}
		}
	}
	return dp[wl1][wl2]
}

// @lc code=end

