/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 *
 * https://leetcode.cn/problems/reverse-string/description/
 *
 * algorithms
 * Easy (80.32%)
 * Likes:    942
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 1.2M
 * Testcase Example:  '["h","e","l","l","o"]'
 *
 * 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
 *
 * 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = ["h","e","l","l","o"]
 * 输出：["o","l","l","e","h"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = ["H","a","n","n","a","h"]
 * 输出：["h","a","n","n","a","H"]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * s[i] 都是 ASCII 码表中的可打印字符
 *
 *
 */

// @lc code=start
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		// s[l], s[r] = s[r], s[l]
		// 构造 原s[l] ^ s[r] 存入当前 s[l]
		s[l] ^= s[r]
		// 当前s[l] ^ 原s[r] 存入 当前s[r]
		// 等价于 原s[l] ^ 原s[r] ^ 原s[r] => 原s[l] 存入 当前s[r]
		// 这时候 当前s[r] == 原s[l]； 且 当前s[l] == 原s[l] ^ 原s[r]

		s[r] ^= s[l]
		// 当前s[r] ^ 当前s[l] 存入 当前s[l]
		// 等价于 原s[l] ^ 原s[l] ^ 原s[r] =>  当前s[l]
		// 得到  原s[r] =>  当前s[l]
		s[l] ^= s[r]

		// 此时互换成功
		l++
		r--
	}
}

// @lc code=end

