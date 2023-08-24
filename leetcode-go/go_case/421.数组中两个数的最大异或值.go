package go_case

/*
 * @lc app=leetcode.cn id=421 lang=golang
 *
 * [421] 数组中两个数的最大异或值
 *
 * https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/description/
 *
 * algorithms
 * Medium (61.02%)
 * Likes:    552
 * Dislikes: 0
 * Total Accepted:    45.9K
 * Total Submissions: 75.3K
 * Testcase Example:  '[3,10,5,25,2,8]'
 *
 * 给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,10,5,25,2,8]
 * 输出：28
 * 解释：最大运算结果是 5 XOR 25 = 28.
 *
 * 示例 2：
 *
 *
 * 输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
 * 输出：127
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^5
 * 0 <= nums[i] <= 2^31 - 1
 *
 *
 *
 *
 */
// https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/solutions/9289/li-yong-yi-huo-yun-suan-de-xing-zhi-tan-xin-suan-f/
// @lc code=start
// const highBit = 30

// type trie struct {
// 	left, right *trie
// }

// func (t *trie) add(num int) {
// 	cur := t
// 	for i := highBit; i >= 0; i-- {
// 		bit := num >> i & 1
// 		if bit == 0 {
// 			if cur.left == nil {
// 				cur.left = &trie{}
// 			}
// 			cur = cur.left
// 		} else {
// 			if cur.right == nil {
// 				cur.right = &trie{}
// 			}
// 			cur = cur.right
// 		}
// 	}
// }

// func (t *trie) check(num int) (x int) {
// 	cur := t
// 	for i := highBit; i >= 0; i-- {
// 		bit := num >> i & 1
// 		if bit == 0 {
// 			// a_i 的第 k 个二进制位为 0，应当往表示 1 的子节点 right 走
// 			if cur.right != nil {
// 				cur = cur.right
// 				x = x*2 + 1
// 			} else {
// 				cur = cur.left
// 				x = x * 2
// 			}
// 		} else {
// 			// a_i 的第 k 个二进制位为 1，应当往表示 0 的子节点 left 走
// 			if cur.left != nil {
// 				cur = cur.left
// 				x = x*2 + 1
// 			} else {
// 				cur = cur.right
// 				x = x * 2
// 			}
// 		}
// 	}
// 	return
// }

// func findMaximumXOR(nums []int) (x int) {
// 	root := &trie{}
// 	for i := 1; i < len(nums); i++ {
// 		// 将 nums[i-1] 放入字典树，此时 nums[0 .. i-1] 都在字典树中
// 		root.add(nums[i-1])
// 		// 将 nums[i] 看作 ai，找出最大的 x 更新答案
// 		x = max(x, root.check(nums[i]))
// 	}
// 	return
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func findMaximumXOR(nums []int) (x int) {
	const highBit = 30 // 最高位的二进制位编号为 30
	for k := highBit; k >= 0; k-- {
		// 将所有的 pre^k(a_j) 放入哈希表中
		seen := map[int]bool{}
		for _, num := range nums {
			// 如果只想保留从最高位开始到第 k 个二进制位为止的部分
			// 只需将其右移 k 位
			seen[num>>k] = true
		}

		// 目前 x 包含从最高位开始到第 k+1 个二进制位为止的部分
		// 我们将 x 的第 k 个二进制位置为 1，即为 x = x*2+1
		xNext := x*2 + 1
		found := false

		// 枚举 i
		for _, num := range nums {
			if seen[num>>k^xNext] {
				found = true
				break
			}
		}

		if found {
			x = xNext
		} else {
			// 如果没有找到满足等式的 a_i 和 a_j，那么 x 的第 k 个二进制位只能为 0
			// 即为 x = x*2
			x = xNext - 1
		}
	}
	return
}

// @lc code=end
