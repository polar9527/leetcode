package go_case

import "slices"

/*
 * @lc app=leetcode.cn id=632 lang=golang
 *
 * [632] 最小区间
 *
 * https://leetcode.cn/problems/smallest-range-covering-elements-from-k-lists/description/
 *
 * algorithms
 * Hard (64.18%)
 * Likes:    507
 * Dislikes: 0
 * Total Accepted:    42K
 * Total Submissions: 65.4K
 * Testcase Example:  '[[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]'
 *
 * 你有 k 个 非递减排列 的整数列表。找到一个 最小 区间，使得 k 个列表中的每个列表至少有一个数包含在其中。
 *
 * 我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
 * 输出：[20,24]
 * 解释：
 * 列表 1：[4, 10, 15, 24, 26]，24 在区间 [20,24] 中。
 * 列表 2：[0, 9, 12, 20]，20 在区间 [20,24] 中。
 * 列表 3：[5, 18, 22, 30]，22 在区间 [20,24] 中。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [[1,2,3],[1,2,3],[1,2,3]]
 * 输出：[1,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums.length == k
 * 1 <= k <= 3500
 * 1 <= nums[i].length <= 50
 * -10^5 <= nums[i][j] <= 10^5
 * nums[i] 按非递减顺序排列
 *
 *
 *
 *
 */

// @lc code=start
func smallestRange(nums [][]int) []int {

	type pair struct{ num, i int }
	var pairs []pair
	for i, list := range nums {
		for _, n := range list {
			pairs = append(pairs, pair{n, i})
		}
	}

	slices.SortFunc(pairs, func(a, b pair) int {
		return a.num - b.num
	})

	resL, resR := pairs[0].num, pairs[len(pairs)-1].num

	count := len(nums)
	indexCnt := make([]int, count)
	left := 0
	for _, p := range pairs {
		r, i := p.num, p.i
		if indexCnt[i] == 0 {
			// p.num 所属的列表首次出现， 增加到窗口 indexCnt 中
			// count 记录还剩多少列表没有出现
			count--
		}
		// 不论列表是否第一次出现，都增加到窗口 indexCnt 中
		indexCnt[i]++
		for count == 0 {
			l, i := pairs[left].num, pairs[left].i
			// 此时，所属的列表都至少有一个出现在窗口 indexCnt 中
			// 先记录候选结果
			if r-l < resR-resL {
				resL, resR = l, r
			}
			// 从左侧缩小窗口
			indexCnt[i]--
			left++
			if indexCnt[i] == 0 {
				// 这个最左侧的 pair 对应的 list（nums[i]）, 在窗口中没有其他对应的 pair了
				// 所以剩余列表计数 count需要加1
				count++
			}

		}

	}
	return []int{resL, resR}
}

// @lc code=end
