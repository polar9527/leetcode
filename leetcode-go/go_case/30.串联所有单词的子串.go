package go_case

/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 *
 * https://leetcode.cn/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (38.90%)
 * Likes:    1200
 * Dislikes: 0
 * Total Accepted:    243.7K
 * Total Submissions: 637.1K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * 给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。
 *
 * s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。
 *
 *
 * 例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"，
 * "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
 *
 *
 * 返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "barfoothefoobarman", words = ["foo","bar"]
 * 输出：[0,9]
 * 解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
 * 子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
 * 子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
 * 输出顺序无关紧要。返回 [9,0] 也是可以的。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
 * 输出：[]
 * 解释：因为 words.length == 4 并且 words[i].length == 4，所以串联子串的长度必须为 16。
 * s 中没有子串长度为 16 并且等于 words 的任何顺序排列的连接。
 * 所以我们返回一个空数组。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
 * 输出：[6,9,12]
 * 解释：因为 words.length == 3 并且 words[i].length == 3，所以串联子串的长度必须为 9。
 * 子串 "foobarthe" 开始位置是 6。它是 words 中以 ["foo","bar","the"] 顺序排列的连接。
 * 子串 "barthefoo" 开始位置是 9。它是 words 中以 ["bar","the","foo"] 顺序排列的连接。
 * 子串 "thefoobar" 开始位置是 12。它是 words 中以 ["the","foo","bar"] 顺序排列的连接。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * 1 <= words.length <= 5000
 * 1 <= words[i].length <= 30
 * words[i] 和 s 由小写英文字母组成
 *
 *
 */

// @lc code=start
// func findSubstring(s string, words []string) []int {
// 	if len(s) == 0 || len(words) == 0 || len(words[0]) == 0 {
// 		return []int{}
// 	}

// 	wordsNum := len(words)
// 	wordLen := len(words[0])
// 	wordMap := make(map[string]int)
// 	for i := range words {
// 		wordMap[words[i]]++
// 	}
// 	//
// 	targetWordsLen := wordsNum * wordLen
// 	res := []int{}

// 	for i := 0; i <= len(s)-targetWordsLen; i++ {
// 		// 判断 在区间 [i:i+targeWordLen) 这个 window 中
// 		// 字符串 出现的次数 和 wordMap 一致
// 		subString := s[i : i+targetWordsLen]
// 		match := true
// 		window := make(map[string]int)

// 		for j := 0; j < len(subString); j += wordLen {
// 			word := subString[j : j+wordLen]
// 			if wordMap[word] == 0 {
// 				match = false
// 				break
// 			}

// 			window[word]++
// 			if window[word] > wordMap[word] {
// 				match = false
// 				break
// 			}
// 		}
// 		if match {
// 			res = append(res, i)
// 		}

// 	}
// 	return res
// }

// optimize
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	wordsNum := len(words)
	wordLen := len(words[0])
	wordMap := make(map[string]int)
	for i := range words {
		wordMap[words[i]]++
	}
	//
	// targetWordsLen := wordsNum * wordLen
	res := []int{}

	for i := 0; i < wordLen; i++ {
		// 子串的起始点
		window := make(map[string]int)
		count := 0

		// left 窗口第一个 word 的起始位置， right 窗口最后一个 word 的起始位置
		for left, right := i, i; right <= len(s)-wordLen; right += wordLen {
			word := s[right : right+wordLen]
			if wordMap[word] > 0 {
				window[word]++
				count++

				// 情况 1，窗口内单词超过需的数量
				// 当前单词次数超出需求，需要将窗口最左侧第一个相同的单词，以及之前的单词
				// 从窗口移除，
				// 并将窗口从左侧缩小至当前单词的下一个
				for window[word] > wordMap[word] {
					// word 和 word 到 当前窗口 最左边的 单词依次弹出窗口
					// 直到将该单词从当前窗口移除
					window[s[left:left+wordLen]]--
					left += wordLen
					count--
				}
				//  找到子串
				if count == wordsNum {
					res = append(res, left)
					window[s[left:left+wordLen]]--
					left += wordLen
					count--
				}

			} else {
				// 情况 2，遇见不在需求范围内的单词
				window = make(map[string]int)
				count = 0
				left = right + wordLen

			}
		}

	}
	return res
}

// @lc code=end
