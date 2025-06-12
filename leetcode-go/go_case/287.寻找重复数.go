/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 *
 * https://leetcode.cn/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (65.72%)
 * Likes:    2621
 * Dislikes: 0
 * Total Accepted:    505.1K
 * Total Submissions: 762.7K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
 *
 * 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
 *
 * 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,3,4,2,2]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,1,3,4,2]
 * 输出：3
 *
 *
 * 示例 3 :
 *
 *
 * 输入：nums = [3,3,3,3,3]
 * 输出：3
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^5
 * nums.length == n + 1
 * 1 <= nums[i] <= n
 * nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如何证明 nums 中至少存在一个重复的数字?
 * 你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？
 *
 *
 */

// @lc code=start
func findDuplicate(nums []int) int {
	s, f := 0, 0
	for s, f = nums[s], nums[nums[f]]; s != f; s, f = nums[s], nums[nums[f]] {

	}

	s = 0
	for ; s != f; s, f = nums[s], nums[f] {

	}
	return s
}

// @lc code=end

