/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 *
 * https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/description/
 *
 * algorithms
 * Easy (72.49%)
 * Likes:    679
 * Dislikes: 0
 * Total Accepted:    367.1K
 * Total Submissions: 500.4K
 * Testcase Example:  '"abbaca"'
 *
 * 给出由小写字母组成的字符串 s，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
 *
 * 在 s 上反复执行重复项删除操作，直到无法继续删除。
 *
 * 在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。
 *
 *
 *
 * 示例：
 *
 *
 * 输入："abbaca"
 * 输出："ca"
 * 解释：
 * 例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串
 * "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * s 仅由小写英文字母组成。
 *
 *
 */

// @lc code=start
func removeDuplicates(s string) string {
	st := make([]rune, 0)
	for _, r := range s {
		if len(st) > 0 && r == st[len(st)-1] {
			st = st[:len(st)-1]
		} else {
			st = append(st, r)
		}
	}
	return string(st)
}

// @lc code=end

