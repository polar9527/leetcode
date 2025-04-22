package go_case

/*
 * @lc app=leetcode.cn id=2615 lang=golang
 *
 * [2615] 等值距离和
 *
 * https://leetcode.cn/problems/sum-of-distances/description/
 *
 * algorithms
 * Medium (40.16%)
 * Likes:    40
 * Dislikes: 0
 * Total Accepted:    8.7K
 * Total Submissions: 21.4K
 * Testcase Example:  '[1,3,1,1,2]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。现有一个长度等于 nums.length 的数组 arr 。对于满足 nums[j] ==
 * nums[i] 且 j != i 的所有 j ，arr[i] 等于所有 |i - j| 之和。如果不存在这样的 j ，则令 arr[i] 等于 0 。
 *
 * 返回数组 arr 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,3,1,1,2]
 * 输出：[5,0,3,4,0]
 * 解释：
 * i = 0 ，nums[0] == nums[2] 且 nums[0] == nums[3] 。因此，arr[0] = |0 - 2| + |0 -
 * 3| = 5 。
 * i = 1 ，arr[1] = 0 因为不存在值等于 3 的其他下标。
 * i = 2 ，nums[2] == nums[0] 且 nums[2] == nums[3] 。因此，arr[2] = |2 - 0| + |2 -
 * 3| = 3 。
 * i = 3 ，nums[3] == nums[0] 且 nums[3] == nums[2] 。因此，arr[3] = |3 - 0| + |3 -
 * 2| = 4 。
 * i = 4 ，arr[4] = 0 因为不存在值等于 2 的其他下标。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,5,3]
 * 输出：[0,0,0]
 * 解释：因为 nums 中的元素互不相同，对于所有 i ，都有 arr[i] = 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^9
 *
 *
 */

// @lc code=start
// func distance(nums []int) []int64 {
// 	abs := func(a, b int) int {
// 		s := a - b
// 		if s > 0 {
// 			return s
// 		}
// 		return -s
// 	}
// 	mn := make(map[int][]int)
// 	for i, n := range nums {
// 		mn[n] = append(mn[n], i)
// 	}

// 	n := len(nums)
// 	arr := make([]int64, n)

// 	for i := 0; i < n; i++ {
// 		if indexes, ok := mn[nums[i]]; ok {
// 			sum := int64(0)
// 			for _, j := range indexes {
// 				// if j == i {
// 				// 	continue
// 				// }
// 				sum += int64(abs(i, j))

// 			}
// 			arr[i] = sum
// 		} else {
// 			arr[i] = 0
// 		}
// 	}
// 	return arr
// }

func distance(nums []int) []int64 {
	// 预处理：记录每个值的所有索引
	valueIndices := make(map[int][]int)
	for i, num := range nums {
		// 相当于排序了valueIndices[num]
		// i 是从小到大添加的
		valueIndices[num] = append(valueIndices[num], i)
	}

	n := len(nums)
	res := make([]int64, n)

	// 对每个值的索引数组计算前缀和
	for _, indices := range valueIndices {
		// 当 m == 0, 相当nums[i] 在 nums 中 没有其他相等的元素，
		// 也就不用做有任何计算
		m := len(indices)
		prefix := make([]int64, m+1)
		for i := 0; i < m; i++ {
			prefix[i+1] = prefix[i] + int64(indices[i])
		}

		// 计算每个索引的绝对差之和
		for i := 0; i < m; i++ {
			x := int64(indices[i])
			// indices 是有序的，
			// 因为在前面创建valueIndices[num]的时候,
			// index 就是从小到大添加的，所以不需要额外排序
			// 因为是有序的，所以 left 和 right 一定是正数
			left := x*int64(i) - prefix[i]
			right := (prefix[m] - prefix[i+1]) - x*int64(m-i-1)
			res[indices[i]] = left + right
		}
	}

	return res
}

// @lc code=end
