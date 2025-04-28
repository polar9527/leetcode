/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode.cn/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (48.94%)
 * Likes:    3043
 * Dislikes: 0
 * Total Accepted:    787.9K
 * Total Submissions: 1.6M
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 *
 * 返回 滑动窗口中的最大值 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
 * 输出：[3,3,5,5,6,7]
 * 解释：
 * 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1], k = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 *
 *
 */

// @lc code=start

func maxSlidingWindow(nums []int, k int) []int {

	q := []int{}
	res := []int{}
	for i, num := range nums {
		// push item
		// 元素加入队列 q 前，要保证 队列单调减
		for len(q) != 0 && nums[q[len(q)-1]] <= num {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		// pop item
		// 窗口的长度刚好大于 k 的时候
		if i-q[0]+1 > k {
			q = q[1:]
		}

		if i >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}

// @lc code=end


