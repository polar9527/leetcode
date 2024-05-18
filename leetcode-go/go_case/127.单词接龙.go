package go_case

/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 *
 * https://leetcode.cn/problems/word-ladder/description/
 *
 * algorithms
 * Hard (48.77%)
 * Likes:    1372
 * Dislikes: 0
 * Total Accepted:    216.4K
 * Total Submissions: 443.6K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 ->
 * s2 -> ... -> sk：
 *
 *
 * 每一对相邻的单词只差一个字母。
 * 对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
 * sk == endWord
 *
 *
 * 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列
 * 中的 单词数目 。如果不存在这样的转换序列，返回 0 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log","cog"]
 * 输出：5
 * 解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
 *
 *
 * 示例 2：
 *
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log"]
 * 输出：0
 * 解释：endWord "cog" 不在字典中，所以无法进行转换。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= beginWord.length <= 10
 * endWord.length == beginWord.length
 * 1 <= wordList.length <= 5000
 * wordList[i].length == beginWord.length
 * beginWord、endWord 和 wordList[i] 由小写英文字母组成
 * beginWord != endWord
 * wordList 中的所有字符串 互不相同
 *
 *
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {

	wordMap := make(map[string]int)
	for i, word := range wordList {
		wordMap[word] = i
	}

	if _, ok := wordMap[endWord]; !ok {
		return 0
	}

	isOneDiff := func(s1, s2 string) bool {
		if len(s1) != len(s2) {
			return false
		}

		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				return s1[i+1:] == s2[i+1:]
			}
		}
		return false
	}

	getCandidates := func(word string, wordList []string) []string {

		var res []string
		for _, wItem := range wordList {
			if isOneDiff(word, wItem) {
				res = append(res, wItem)
			}
		}
		return res
	}

	// 用26个英文字母分别替换掉各个位置的字母，生成一个结果集
	// getCandidates := func(word string) []string {
	// 	var res []string
	// 	for i := 0; i < 26; i++ {
	// 		for j := 0; j < len(word); j++ {
	// 			if word[j] != byte(int('a')+i) {
	// 				res = append(res, word[:j]+string(int('a')+i)+word[j+1:])
	// 			}
	// 		}
	// 	}
	// 	return res
	// }

	queue := []string{beginWord}
	depth := 0

	for len(queue) > 0 {
		depth++
		// 广度遍历第i层 有多少个节点 需要遍历
		iLevel := len(queue)
		for i := 0; i < iLevel; i++ {
			curWord := queue[0]
			queue = queue[1:]
			candidates := getCandidates(curWord, wordList)
			// candidates := getCandidates(curWord)
			for _, candidate := range candidates {
				if _, exist := wordMap[candidate]; exist {
					if candidate == endWord {
						return depth + 1
					}
					// 防止有向图循环
					delete(wordMap, candidate)

					queue = append(queue, candidate)
				}

			}

		}

	}
	return 0

}

// @lc code=end
