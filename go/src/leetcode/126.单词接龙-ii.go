/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 *
 * https://leetcode-cn.com/problems/word-ladder-ii/description/
 *
 * algorithms
 * Hard (31.48%)
 * Likes:    245
 * Dislikes: 0
 * Total Accepted:    17.3K
 * Total Submissions: 45.5K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord
 * 的最短转换序列。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换过程中的中间单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回一个空列表。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出:
 * [
 * ⁠ ["hit","hot","dot","dog","cog"],
 * ["hit","hot","lot","log","cog"]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: []
 *
 * 解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。
 *
 */
package main

import (
	"math"
)

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// 字典化 wordlist, 每个string 对应一个 id
	ids := make(map[string]int)
	for i, word := range wordList {
		ids[word] = i
	}

	// 将 beginWord 加入字典，beginWord之后将作为起始节点
	if _, ok := ids[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		ids[beginWord] = len(wordList) - 1
	}
	// 如果endWord不在字典中，可以判定无结果
	if _, ok := ids[endWord]; !ok {
		return [][]string{}
	}

	// 为每个节点建立有向弧 i -> {x,y,z...}
	n := len(wordList)
	edges := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if transformCheck(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}

	var ans [][]string
	cost := make([]int, n)
	// queue 中每一项存放的是 起始节点 到 某个节点的路径
	queue := [][]int{[]int{ids[beginWord]}}
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[ids[beginWord]] = 0

	for i := 0; i < len(queue); i++ {
		now := queue[i]
		last := now[len(now)-1]
		if last == ids[endWord] {
			var tmp []string
			for _, index := range now {
				tmp = append(tmp, wordList[index])
			}
			ans = append(ans, tmp)
		} else {
			for _, to := range edges[last] {
				// cost[to] < cost[last] + 1
				// 说明 to 所在的相对于起始节点而言的level 小于或者等于 last, 所以to不能成为last节点的后续路径
				if cost[to] >= cost[last]+1 {
					// cost[to] >= cost[last] + 1 只有两种可能，
					// 1. cost[to] == math.MaxInt32 > cost[last] + 1
					// 此时cost[to]还没有被遍历到
					// 2. cost[to] == cost[last] + 1
					// 此时cost[to]已经被和last处于同一level的其他节点遍历到了
					cost[to] = cost[last] + 1
					tmp := make([]int, len(now))
					copy(tmp, now)
					tmp = append(tmp, to)
					queue = append(queue, tmp)
				}
			}
		}
	}
	return ans
}

func transformCheck(from, to string) bool {
	for i := 0; i < len(from); i++ {
		if from[i] != to[i] {
			return from[i+1:] == to[i+1:]
		}
	}
	return false
}

// @lc code=end
