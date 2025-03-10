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

func NewMoQueue() *MoQueue {
	return &MoQueue{
		q: make([]int, 0),
	}
}

type MoQueue struct {
	q []int
}

func (m *MoQueue) Front() int {
	return m.q[0]
}

func (m *MoQueue) Back() int {
	return m.q[len(m.q)-1]
}

func (m *MoQueue) Push(v int) {
	for !m.Empty() && v > m.Back() {
		m.q = m.q[0 : len(m.q)-1]
	}
	m.q = append(m.q, v)
}

func (m *MoQueue) Pop(v int) {
	if !m.Empty() && v == m.Front() {
		m.q = m.q[1:]
	}
}

func (m *MoQueue) Empty() bool {
	return len(m.q) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMoQueue()
	res := make([]int, 0)

	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	res = append(res, queue.Front())

	l := len(nums)
	for i := k; i < l; i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		res = append(res, queue.Front())
	}
	return res
}

// @lc code=end


