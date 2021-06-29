/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 *
 * https://leetcode-cn.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (52.31%)
 * Likes:    410
 * Dislikes: 0
 * Total Accepted:    190.1K
 * Total Submissions: 363.1K
 * Testcase Example:  '"leetcode"'
 *
 * 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 *
 *
 *
 * 示例：
 *
 * s = "leetcode"
 * 返回 0
 *
 * s = "loveleetcode"
 * 返回 2
 *
 *
 *
 *
 * 提示：你可以假定该字符串只包含小写字母。
 *
 */
package leetcode

// @lc code=start
type pair struct {
	ch  byte
	pos int
}

// Priority Queue
func firstUniqChar(s string) int {
	queue := []pair{}
	dict := [26]int{}
	for i, c := range s {
		key := c - 'a'
		dict[key]++
		queue = append(queue, pair{byte(c), i})
		for len(queue) > 0 && dict[queue[0].ch-'a'] > 1 {
			queue = queue[1:]
		}
	}
	if len(queue) > 0 {
		return queue[0].pos
	}
	return -1
}

//
func firstUniqCharAlphabet(s string) int {
	// key = char's index in alphabet
	// value = char's position in ( value != 0 && value != -1)
	// value == 0, char does not appear in s
	// value == -1, char appear at least two times in s
	dict := [26]int{}
	// O(n)
	for i, c := range s {
		key := c - 'a'
		if dict[key] > 0 {
			// 不是第一次看见这个字母
			dict[key] = -1
		}
		if dict[key] == 0 {
			// 第一次看见这个字母
			dict[key] = i + 1
		}
		// 两个判断不能前后颠倒顺序
	}
	//O(alphabet)
	res := len(s) + 1
	for i := range dict {
		if dict[i] > 0 && dict[i] < res {
			res = dict[i]
		}
	}
	if res == len(s)+1 {
		return -1
	}
	return res - 1
}

// @lc code=end
