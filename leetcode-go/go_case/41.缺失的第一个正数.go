/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 *
 * https://leetcode.cn/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (46.57%)
 * Likes:    2344
 * Dislikes: 0
 * Total Accepted:    539.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,0]'
 *
 * 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
 * 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,0]
 * 输出：3
 * 解释：范围 [1,2] 中的数字都在数组中。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,4,-1,1]
 * 输出：2
 * 解释：1 在数组中，但 2 没有。
 *
 * 示例 3：
 *
 *
 * 输入：nums = [7,8,9,11,12]
 * 输出：1
 * 解释：最小的正数 1 没有出现。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 */

// @lc code=start
// func firstMissingPositive(nums []int) int {

// 	set := make(map[int]bool)
// 	for _, num := range nums {
// 		set[num] = true
// 	}

// 	for i := 1; i <= len(nums); i++ {
// 		if !set[i] {
// 			return i
// 		}
// 	}
// 	return len(nums) + 1
// }

func firstMissingPositive(nums []int) int {
	n := len(nums)
	i := 0
	for i < n {
		num := nums[i]
		// 如果当前数字是正整数且在数组范围内，且不在正确的位置上，则交换
		if num > 0 && num <= n && nums[num-1] != num {
			nums[i], nums[num-1] = nums[num-1], nums[i]
		} else {
			i++
		}
	}

	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

// @lc code=end

