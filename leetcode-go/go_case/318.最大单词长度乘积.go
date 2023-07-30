package go_case

/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 *
 * https://leetcode.cn/problems/maximum-product-of-word-lengths/description/
 *
 * algorithms
 * Medium (72.59%)
 * Likes:    417
 * Dislikes: 0
 * Total Accepted:    68.3K
 * Total Submissions: 94.1K
 * Testcase Example:  '["abcw","baz","foo","bar","xtfn","abcdef"]'
 *
 * 给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j])
 * 的最大值，并且这两个单词不含有公共字母。如果不存在这样的两个单词，返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：words = ["abcw","baz","foo","bar","xtfn","abcdef"]
 * 输出：16
 * 解释：这两个单词为 "abcw", "xtfn"。
 *
 * 示例 2：
 *
 *
 * 输入：words = ["a","ab","abc","d","cd","bcd","abcd"]
 * 输出：4
 * 解释：这两个单词为 "ab", "cd"。
 *
 * 示例 3：
 *
 *
 * 输入：words = ["a","aa","aaa","aaaa"]
 * 输出：0
 * 解释：不存在这样的两个单词。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= words.length <= 1000
 * 1 <= words[i].length <= 1000
 * words[i] 仅包含小写字母
 *
 *
 */

// @lc code=start
func maxProduct(words []string) int {
	ret := 0
	bitMap := make(map[int]int)
	for _, word := range words {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		if len(word) > bitMap[mask] {
			bitMap[mask] = len(word)
		}
	}

	for xk, xl := range bitMap {
		for yk, yl := range bitMap {
			if xk&yk == 0 && xl*yl > ret {
				ret = xl * yl
			}
		}
	}
	return ret
}

// @lc code=end
