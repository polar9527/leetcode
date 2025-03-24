/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 *
 * https://leetcode.cn/problems/candy/description/
 *
 * algorithms
 * Hard (48.47%)
 * Likes:    1633
 * Dislikes: 0
 * Total Accepted:    399.3K
 * Total Submissions: 824.3K
 * Testcase Example:  '[1,0,2]'
 *
 * n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
 *
 * 你需要按照以下要求，给这些孩子分发糖果：
 *
 *
 * 每个孩子至少分配到 1 个糖果。
 * 相邻两个孩子评分更高的孩子会获得更多的糖果。
 *
 *
 * 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：ratings = [1,0,2]
 * 输出：5
 * 解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
 *
 *
 * 示例 2：
 *
 *
 * 输入：ratings = [1,2,2]
 * 输出：4
 * 解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
 * ⁠    第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。
 *
 *
 *
 * 提示：
 *
 *
 * n == ratings.length
 * 1 <= n <= 2 * 10^4
 * 0 <= ratings[i] <= 2 * 10^4
 *
 *
 */

// @lc code=start
func candy(ratings []int) int {

	l := len(ratings)
	candies := make([]int, l)

	candies[0] = 1
	for i := 1; i < l; i++ {
		candies[i] = 1
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	sum := candies[l-1]
	for j := l - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			if candies[j] <= candies[j+1] {
				candies[j] = candies[j+1] + 1
			}
		}
		sum += candies[j]
	}
	return sum
}

// @lc code=end

