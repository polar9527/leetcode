/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (44.37%)
 * Likes:    3128
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 2.3M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 整数数组 nums 按升序排列，数组中的值 互不相同 。
 *
 * 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k],
 * nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始
 * 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
 *
 * 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
 *
 * 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 0
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 3
 * 输出：-1
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1], target = 0
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * nums 中的每个值都 独一无二
 * 题目数据保证 nums 在预先未知的某个下标上进行了旋转
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start
func search(nums []int, target int) int {

	n := len(nums)
	// nums[x] 是否 为 target， 或者 在target 右侧
	check := func(x int) bool {
		// 如果 nums[x] > nums[n-1]，说明 x 在旋转点的左侧（较大的部分）
		if nums[x] > nums[n-1] {
			// 目标值target必须大于 nums[n-1] 且小于等于 nums[x]
			return target > nums[n-1] && nums[x] >= target
		}
		// 否则，x 在旋转点的右侧（较小的部分）
		// 目标值target必须大于 nums[n-1] 或小于等于 nums[x]
		return target > nums[n-1] || nums[x] >= target
	}
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)>>1

		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// [l:] 这部分 >= target, l 最大 为 n-1
	if l < n && nums[l] == target {
		return l
	}
	return -1
}

// @lc code=end

