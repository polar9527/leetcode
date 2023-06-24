/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {

	wordSet := make(map[string]struct{})
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	n := len(s)
	// 对于每个s[i], s[i:] 中所有的 能拆分成wordSet中的单词的组合，
	// 每个组合就是一个[]string, 所有的组合就是[][]string
	dp := make([][][]string, n)

	var backTrace func(int) [][]string

	backTrace = func(index int) [][]string {
		// index must < n
		if dp[index] != nil {
			return dp[index]
		}

		wordLists := [][]string{}
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextWords := range backTrace(i) {
					wordLists = append(wordLists, append([]string{word}, nextWords...))
				}
			}
		}

		word := s[index:]
		if _, has := wordSet[word]; has {
			wordLists = append(wordLists, []string{word})
		}

		dp[index] = wordLists
		return wordLists
	}

	var sentences []string
	for _, words := range backTrace(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}

	return sentences
}

// @lc code=end

