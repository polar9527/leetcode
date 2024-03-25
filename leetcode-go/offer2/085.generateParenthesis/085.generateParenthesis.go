package offer2

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.62%)
 * Likes:    3538
 * Dislikes: 0
 * Total Accepted:    813.8K
 * Total Submissions: 1M
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 8
 *
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {

	var res []string
	var path []byte
	var backtracing func(int, int)
	backtracing = func(leftOpen, rightClose int) {
		if len(path) == 2*n {
			leafResult := append([]byte{}, path...)
			res = append(res, string(leafResult))
			return
		}

		if leftOpen < n {
			path = append(path, '(')
			backtracing(leftOpen+1, rightClose)
			path = path[:len(path)-1]
		}

		if rightClose < leftOpen {
			path = append(path, ')')
			backtracing(leftOpen, rightClose+1)
			path = path[:len(path)-1]
		}
	}

	backtracing(0, 0)
	return res
}

// @lc code=end
