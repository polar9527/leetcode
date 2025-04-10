package go_case

/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 *
 * https://leetcode.cn/problems/word-ladder-ii/description/
 *
 * algorithms
 * Hard (36.58%)
 * Likes:    747
 * Dislikes: 0
 * Total Accepted:    63.7K
 * Total Submissions: 174.1K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，一个表示此过程的 转换序列 是形式上像 beginWord ->
 * s1 -> s2 -> ... -> sk 这样的单词序列，并满足：
 *
 *
 *
 *
 * 每对相邻的单词之间仅有单个字母不同。
 * 转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList
 * 中的单词。
 * sk == endWord
 *
 *
 * 给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回所有从 beginWord 到 endWord
 * 的 最短转换序列 ，如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk]
 * 的形式返回。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log","cog"]
 * 输出：[["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
 * 解释：存在 2 种最短的转换序列：
 * "hit" -> "hot" -> "dot" -> "dog" -> "cog"
 * "hit" -> "hot" -> "lot" -> "log" -> "cog"
 *
 *
 * 示例 2：
 *
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log"]
 * 输出：[]
 * 解释：endWord "cog" 不在字典 wordList 中，所以不存在符合要求的转换序列。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= beginWord.length <= 5
 * endWord.length == beginWord.length
 * 1 <= wordList.length <= 500
 * wordList[i].length == beginWord.length
 * beginWord、endWord 和 wordList[i] 由小写英文字母组成
 * beginWord != endWord
 * wordList 中的所有单词 互不相同
 *
 *
 *
 *
 */

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	wordMap := make(map[string]bool, 0)
	for _, word := range wordList {
		wordMap[word] = true
	}

	if !wordMap[endWord] {
		return [][]string{}
	}
	wordMap[beginWord] = true

	distance := make(map[string]int, 0)
	graph := make(map[string][]string)

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

	getCandidates := func(word string) []string {

		var res []string
		for wItem, _ := range wordMap {
			if isOneDiff(word, wItem) {
				res = append(res, wItem)
			}
		}
		return res
	}

	// 广度优先，从endWord 开始 向 beginWord 开始 层序遍历，同时建立邻接列表，
	// 并且 给每个word 节点，通过在上一个节点基础上累加的方式，标注相应的最短距离，
	// 层序遍历第一次碰到该节点时标注，因为此时是最短的。
	// 统计两个关键的数据
	// 1.邻接列表 graph
	// 2.最短距离 distance
	var bfs func(string)
	bfs = func(word string) {

		q := []string{word}
		distance[word] = 0

		for len(q) > 0 {
			iLevel := len(q)
			for i := 0; i < iLevel; i++ {
				curWord := q[0]
				q = q[1:]

				preWords := getCandidates(curWord)
				for _, preWord := range preWords {
					graph[preWord] = append(graph[preWord], curWord)
					if _, ok := distance[preWord]; !ok {
						distance[preWord] = distance[curWord] + 1
						q = append(q, preWord)
					}
				}

			}
		}
	}

	var res [][]string
	var path []string

	// 按照最短距离 用 回溯+深搜，遍历所有可能路径
	// 通过 distance[word] == distance[nextWord]+1  控制深搜方向
	var dfs func(string)
	dfs = func(word string) {
		if word == endWord {
			res = append(res, append([]string{}, path...))
			return
		}
		for _, nextWord := range graph[word] {
			if distance[word] == distance[nextWord]+1 {
				path = append(path, nextWord)
				dfs(nextWord)
				path = path[:len(path)-1]
			}

		}
	}

	bfs(endWord)
	path = append(path, beginWord)
	dfs(beginWord)
	return res
}

// @lc code=end
