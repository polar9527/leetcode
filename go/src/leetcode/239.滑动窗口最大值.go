/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode-cn.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (47.32%)
 * Likes:    356
 * Dislikes: 0
 * Total Accepted:    47.4K
 * Total Submissions: 100.1K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 *
 * 返回滑动窗口中的最大值。
 *
 *
 *
 * 进阶：
 *
 * 你能在线性时间复杂度内解决此题吗？
 *
 *
 *
 * 示例:
 *
 * 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
 * 输出: [3,3,5,5,6,7]
 * 解释:
 *
 * ⁠ 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 *
 *
 */
package leetcode

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	// 窗口内最大值后面更小的数要存起来
	ret := make([]int, 0)
	d := new(Dqueue)
	for i := 0; i < k; i++ {
		for d.len() != 0 && nums[d.tail()] <= nums[i] {
			d.popBack()
		}
		d.pushBack(i)
	}
	for i := k; i < len(nums); i++ {
		ret = append(ret, nums[d.head()])
		for d.len() != 0 && nums[d.tail()] <= nums[i] {
			d.popBack()
		}
		if d.len() != 0 && i-k >= d.head() {
			d.popFront()
		}
		d.pushBack(i)
	}
	ret = append(ret, nums[d.head()])
	return ret
}

type Dqueue struct {
	data []int
}

func (d *Dqueue) len() int {
	return len(d.data)
}

func (d *Dqueue) popFront() {
	if d.len() != 0 {
		d.data = d.data[1:]
	}
}
func (d *Dqueue) popBack() {
	if d.len() != 0 {
		d.data = d.data[:len(d.data)-1]
	}
}

func (d *Dqueue) pushBack(n int) {
	d.data = append(d.data, n)
}

func (d *Dqueue) tail() int {
	if len(d.data) > 0 {
		return d.data[d.len()-1]
	}
	panic("Dqueue index out of range")
}

func (d *Dqueue) head() int {
	if len(d.data) > 0 {
		return d.data[0]
	}
	panic("Dqueue index out of range")
}

// @lc code=end
