package offer2

/*
 * @lc app=leetcode.cn id=346 lang=golang
 * 示例：
 *
 * 输入：
 * ["MovingAverage", "next", "next", "next", "next"]
 * [[3], [1], [10], [3], [5]]
 * 输出：
 * [null, 1.0, 5.5, 4.66667, 6.0]
 *
 * 解释：
 * MovingAverage movingAverage = new MovingAverage(3);
 * movingAverage.next(1); // 返回 1.0 = 1 / 1
 * movingAverage.next(10); // 返回 5.5 = (1 + 10) / 2
 * movingAverage.next(3); // 返回 4.66667 = (1 + 10 + 3) / 3
 * movingAverage.next(5); // 返回 6.0 = (10 + 3 + 5) / 3
 *
 *
 * 提示：
 *
 * 1 <= size <= 1000
 * -105 <= val <= 105
 * 最多调用 next 方法 104 次
 * 通过次数
 * 20.5K
 * 提交次数
 * 28.2K
 * 通过率
 * 72.5%
 */

// @lc code=start
type MovingAverage struct {
	size, sum int
	q         []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (m *MovingAverage) Next(val int) float64 {
	if len(m.q) == m.size {
		m.sum -= m.q[0]
		m.q = m.q[1:]
	}
	m.sum += val
	m.q = append(m.q, val)
	return float64(m.sum) / float64(len(m.q))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
// @lc code=end
