/*
 * @lc app=leetcode.cn id=327 lang=golang
 *
 * [327] 区间和的个数
 *
 * https://leetcode-cn.com/problems/count-of-range-sum/description/
 *
 * algorithms
 * Hard (43.08%)
 * Likes:    319
 * Dislikes: 0
 * Total Accepted:    21.5K
 * Total Submissions: 49.9K
 * Testcase Example:  '[-2,5,-1]\n-2\n2'
 *
 * 给定一个整数数组 nums 。区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。
 *
 * 请你以下标 i （0  ）为起点，元素个数逐次递增，计算子数组内的元素和。
 *
 * 当元素和落在范围 [lower, upper] （包含 lower 和 upper）之内时，记录子数组当前最末元素下标 j ，记作 有效 区间和
 * S(i, j) 。
 *
 * 求数组中，值位于范围 [lower, upper] （包含 lower 和 upper）之内的 有效 区间和的个数。
 *
 *
 *
 * 注意：
 * 最直观的算法复杂度是 O(n^2) ，请在此基础上优化你的算法。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：nums = [-2,5,-1], lower = -2, upper = 2,
 * 输出：3
 * 解释：
 * 下标 i = 0 时，子数组 [-2]、[-2,5]、[-2,5,-1]，对应元素和分别为 -2、3、2 ；其中 -2 和 2 落在范围 [lower
 * = -2, upper = 2] 之间，因此记录有效区间和 S(0,0)，S(0,2) 。
 * 下标 i = 1 时，子数组 [5]、[5,-1] ，元素和 5、4 ；没有满足题意的有效区间和。
 * 下标 i = 2 时，子数组 [-1] ，元素和 -1 ；记录有效区间和 S(2,2) 。
 * 故，共有 3 个有效区间和。
 *
 *
 *
 * 提示：
 *
 *
 * 0
 *
 *
 */
package leetcode

// @lc code=start
func countRangeSum(nums []int, lower int, upper int) int {

	var mergeCount func([]int) int
	mergeCount = func(arr []int) int {
		l := len(arr)
		if l <= 1 {
			return 0
		}

		nl := append([]int{}, arr[:l/2]...)
		nr := append([]int{}, arr[l/2:]...)

		cnt := mergeCount(nl) + mergeCount(nr)

		l, r := 0, 0
		for _, v := range nl {
			for l < len(nr) && nr[l]-v < lower {
				l++
			}
			for r < len(nr) && nr[r]-v <= upper {
				r++
			}
			cnt += r - l
		}

		pl, pr := 0, 0
		for i := range arr {
			if pl < len(nl) && (pr == len(nr) || nl[pl] < nr[pr]) {
				arr[i] = nl[pl]
				pl++
			} else {
				arr[i] = nr[pr]
				pr++
			}
		}

		return cnt

	}

	prefixSum := make([]int, len(nums)+1)
	for i, v := range nums {
		prefixSum[i+1] = prefixSum[i] + v
	}

	return mergeCount(prefixSum)
}

// @lc code=end
