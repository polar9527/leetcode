package go_case

/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 *
 * https://leetcode.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (38.50%)
 * Likes:    18865
 * Dislikes: 1184
 * Total Accepted:    1.8M
 * Total Submissions: 4.7M
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * You are given an integer array nums. You are initially positioned at the
 * array's first index, and each element in the array represents your maximum
 * jump length at that position.
 *
 * Return true if you can reach the last index, or false otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,3,1,1,4]
 * Output: true
 * Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last
 * index.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [3,2,1,0,4]
 * Output: false
 * Explanation: You will always arrive at index 3 no matter what. Its maximum
 * jump length is 0, which makes it impossible to reach the last index.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^4
 * 0 <= nums[i] <= 10^5
 *
 *
 */

// @lc code=start
func canJump(nums []int) bool {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cover := 0
	n := len(nums) - 1
	for i := 0; i <= cover; i++ { // 每次与覆盖值比较
		cover = max(i+nums[i], cover) //每走一步都将 cover 更新为最大值
		if cover >= n {
			return true
		}
	}
	return false
}

// @lc code=end
