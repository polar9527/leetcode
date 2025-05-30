/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode.cn/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (45.75%)
 * Likes:    1051
 * Dislikes: 0
 * Total Accepted:    320.3K
 * Total Submissions: 698.5K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的 排列。如果是，返回 true ；否则，返回 false 。
 *
 * 换句话说，s1 的排列之一是 s2 的 子串 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "ab" s2 = "eidbaooo"
 * 输出：true
 * 解释：s2 包含 s1 的排列之一 ("ba").
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1= "ab" s2 = "eidboaoo"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s1.length, s2.length <= 10^4
 * s1 和 s2 仅包含小写字母
 *
 *
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1 > n2 {
		return false
	}
	count1 := [26]int{}
	count2 := [26]int{}

	for i, c := range s1 {
		count1[c-'a']++
		count2[s2[i]-'a']++
	}
	if count1 == count2 {
		return true
	}

	for i := n1; i < n2; i++ {
		count2[s2[i]-'a']++
		count2[s2[i-n1]-'a']--
		if count1 == count2 {
			return true
		}
	}
	return false
}

// @lc code=end

˝