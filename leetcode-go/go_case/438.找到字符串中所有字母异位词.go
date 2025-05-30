package go_case

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Medium (54.62%)
 * Likes:    1233
 * Dislikes: 0
 * Total Accepted:    304K
 * Total Submissions: 556.8K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
 *
 * 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "cbaebabacd", p = "abc"
 * 输出: [0,6]
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "abab", p = "ab"
 * 输出: [0,1,2]
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= s.length, p.length <= 3 * 10^4
 * s 和 p 仅包含小写字母
 *
 *
 */

// @lc code=start
func findAnagrams(s, p string) (ans []int) {
	sl := len(s)
	pl := len(p)
	if sl < pl {
		return ans
	}
	sm := [26]int{}
	pm := [26]int{}

	for i := range p {
		pm[p[i]-'a']++
		// sl >= pl, 所以不担心越界
		sm[s[i]-'a']++
	}

	if pm == sm {
		ans = append(ans, 0)
	}

	for i := 1; i < sl-pl+1; i++ {
		sm[s[i-1]-'a']--
		sm[s[i+pl-1]-'a']++
		if sm == pm {
			ans = append(ans, i)
		}
	}
	return ans
}

// @lc code=end
