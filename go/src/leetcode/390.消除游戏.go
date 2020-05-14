/*
 * @lc app=leetcode.cn id=390 lang=golang
 *
 * [390] 消除游戏
 *
 * https://leetcode-cn.com/problems/elimination-game/description/
 *
 * algorithms
 * Medium (43.08%)
 * Likes:    63
 * Dislikes: 0
 * Total Accepted:    3K
 * Total Submissions: 6.8K
 * Testcase Example:  '9'
 *
 * 给定一个从1 到 n 排序的整数列表。
 * 首先，从左到右，从第一个数字开始，每隔一个数字进行删除，直到列表的末尾。
 * 第二步，在剩下的数字中，从右到左，从倒数第一个数字开始，每隔一个数字进行删除，直到列表开头。
 * 我们不断重复这两步，从左到右和从右到左交替进行，直到只剩下一个数字。
 * 返回长度为 n 的列表中，最后剩下的数字。
 *
 * 示例：
 *
 *
 * 输入:
 * n = 9,
 * 1 2 3 4 5 6 7 8 9
 * 2 4 6 8
 * 2 6
 * 6
 *
 * 输出:
 * 6
 *
 */
package main

// @lc code=start
func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n%2 != 0 {
		return lastRemaining(n - 1)
	}
	return 2 * (n/2 + 1 - lastRemaining(n/2))
}

// @lc code=end
