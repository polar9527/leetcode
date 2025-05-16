package go_case

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

	// 在 2n 的 位置上 选择 n 个位置放置左括号（
	var res []string
	var bt func(int, int, string)
	bt = func(left, right int, path string) {
		if len(path) == 2*n {

			res = append(res, path)
		}
		// 选择左括号
		// 左括号的数量小于 n, 还可以继续添加左括号
		if left < n {
			bt(left+1, right, path+"(")
		}
		// 不选左括号，选右括号
		// 同时要满足条件，即前缀的右括号数量要小于左括号
		// 否则就是错误结果
		if right < left {
			bt(left, right+1, path+")")
		}
	}
	bt(0, 0, "")
	return res
}

// @lc code=end
