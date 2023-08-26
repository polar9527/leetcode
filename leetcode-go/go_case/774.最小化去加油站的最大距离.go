package go_case

/*
 * @lc app=leetcode.cn id=774 lang=golang
 *
 * [774] 最小化去加油站的最大距离
 *
 * https://leetcode.cn/problems/minimize-max-distance-to-gas-station/description/
 *
 * algorithms
 * Hard (63.65%)
 * Likes:    63
 * Dislikes: 0
 * Total Accepted:    2.3K
 * Total Submissions: 3.7K
 * Testcase Example:  '[1,2,3,4,5,6,7,8,9,10]\n9'
 *
 * 整数数组 stations 表示 水平数轴 上各个加油站的位置。给你一个整数 k 。
 *
 * 请你在数轴上增设 k 个加油站，新增加油站可以位于 水平数轴 上的任意位置，而不必放在整数位置上。
 *
 * 设 penalty() 是：增设 k 个新加油站后，相邻 两个加油站间的最大距离。
 * 请你返回 penalty() 可能的最小值。与实际答案误差在 10^-6 范围内的答案将被视作正确答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：stations = [1,2,3,4,5,6,7,8,9,10], k = 9
 * 输出：0.50000
 *
 *
 * 示例 2：
 *
 *
 * 输入：stations = [23,24,36,39,46,56,57,65,84,98], k = 1
 * 输出：14.00000
 *
 *
 *
 *
 * 提示：
 *
 *
 * 10
 * 0
 * stations 按 严格递增 顺序排列
 * 1
 *
 *
 */
// https://www.bbsmax.com/A/WpdK4prNzV/
// @lc code=start
func minmaxGasDist(stations []int, k int) float64 {
	var maxDistance int
	for i := 1; i < len(stations); i++ {
		maxDistance = max(maxDistance, stations[i]-stations[i-1])
	}
	l, r := 0.0, float64(maxDistance)
	for (r - l) > 1e-6 {
		mid := (l + r) / 2
		cnt := 0
		for i := 1; i < len(stations); i++ {
			cnt += int(float64(stations[i]-stations[i-1]) / mid)
		}
		// 要找到 mid 的最小值，满足 penalty(stations, mid) <= float64(k)
		if cnt <= k {
			r = mid
		} else {
			l = mid
		}
	}

	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
