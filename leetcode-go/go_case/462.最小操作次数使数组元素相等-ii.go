/*
 * @lc app=leetcode.cn id=462 lang=golang
 *
 * [462] 最小操作次数使数组元素相等 II
 *
 * https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/description/
 *
 * algorithms
 * Medium (62.62%)
 * Likes:    320
 * Dislikes: 0
 * Total Accepted:    59K
 * Total Submissions: 94.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最小操作数。
 *
 * 在一次操作中，你可以使数组中的一个元素加 1 或者减 1 。
 *
 * 测试用例经过设计以使答案在 32 位 整数范围内。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：2
 * 解释：
 * 只需要两次操作（每次操作指南使一个元素加 1 或减 1）：
 * [1,2,3]  =>  [2,2,3]  =>  [2,2,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,10,2,9]
 * 输出：16
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

// @lc code=start
func minMoves2(nums []int) (ans int) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	sort.Ints(nums)
	x := nums[len(nums)/2]
	for _, num := range nums {
		ans += abs(num - x)
	}
	return
}

// @lc code=end

