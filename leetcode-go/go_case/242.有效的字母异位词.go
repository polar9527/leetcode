package go_case

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (64.02%)
 * Likes:    397
 * Dislikes: 0
 * Total Accepted:    249K
 * Total Submissions: 388.8K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 *
 * 示例 1:
 *
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "rat", t = "car"
 * 输出: false
 *
 * 说明:
 * 你可以假设字符串只包含小写字母。
 *
 * 进阶:
 * 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 *
 */

// @lc code=start
// func isAnagram(s string, t string) bool {
// 	return isAnagramHash(s, t)
// }

// func isAnagramSort(s, t string) bool {
// 	if s == t {
// 		return false
// 	}
// 	s1, s2 := []byte(s), []byte(t)
// 	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
// 	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
// 	return string(s1) == string(s2)
// }

// func isAnagramHash(s, t string) bool {
// 	if s == t {
// 		return false
// 	}
// 	var c1, c2 [26]int
// 	for _, ch := range s {
// 		c1[ch-'a']++
// 	}
// 	for _, ch := range t {
// 		c2[ch-'a']++
// 	}
// 	return c1 == c2
// }

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	ld := [26]int{}
// 	for _, r := range s {
// 		ld[r-'a']++
// 	}
// 	for _, r := range t {
// 		ld[r-'a']--
// 	}

// 	return ld == [26]int{}
// }

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	records := [26]int{}
	for index := 0; index < len(s); index++ {

		sCharIndex := s[index] - 'a'
		tCharIndex := t[index] - 'a'
		records[sCharIndex]++
		records[tCharIndex]--
	}
	return records == [26]int{}
}

// @lc code=end
