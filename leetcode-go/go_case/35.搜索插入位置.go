package go_case

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 *
 * https://leetcode.cn/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (45.16%)
 * Likes:    2118
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.7M
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *
 * 请必须使用时间复杂度为 O(log n) 的算法。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,3,5,6], target = 5
 * 输出: 2
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,3,5,6], target = 2
 * 输出: 1
 *
 *
 * 示例 3:
 *
 *
 * 输入: nums = [1,3,5,6], target = 7
 * 输出: 4
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 为 无重复元素 的 升序 排列数组
 * -10^4 <= target <= 10^4
 *
 *
 */
// https://leetcode.cn/problems/N6YdxV/solutions/1228864/jian-zhi-offer-ii-068-cha-zhao-cha-ru-we-acsn/
// @lc code=start
func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	l, r := 0, length
	for l < r {
		mid := l + (r-l)>>1
		if target == nums[mid] {
			r = mid // lower_bound 查找第一个大于或等于 target 的数字，不直接返回，收缩右边界
		} else if target < nums[mid] {
			r = mid // 左闭右开,所以这里取 mid 而不是 mid + 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// @lc code=end
