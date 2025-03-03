/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 *
 * https://leetcode.cn/problems/ransom-note/description/
 *
 * algorithms
 * Easy (64.64%)
 * Likes:    952
 * Dislikes: 0
 * Total Accepted:    591.2K
 * Total Submissions: 889.4K
 * Testcase Example:  '"a"\n"b"'
 *
 * 给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
 *
 * 如果可以，返回 true ；否则返回 false 。
 *
 * magazine 中的每个字符只能在 ransomNote 中使用一次。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：ransomNote = "a", magazine = "b"
 * 输出：false
 *
 *
 * 示例 2：
 *
 *
 * 输入：ransomNote = "aa", magazine = "ab"
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：ransomNote = "aa", magazine = "aab"
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= ransomNote.length, magazine.length <= 10^5
 * ransomNote 和 magazine 由小写英文字母组成
 *
 *
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	// 遍历 magazine 的到 字母字典 records
	// 遍历 ransomNote 对 字母字典 做删减，如果 出现小于0的情况 直接返回 false
	if len(ransomNote) > len(magazine) {
		return false
	}
	records := [26]int{}
	for _, r := range magazine {
		records[r-'a']++
	}
	for _, r := range ransomNote {
		records[r-'a']--
		if records[r-'a'] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

