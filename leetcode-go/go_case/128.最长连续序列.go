package go_case

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 *
 * https://leetcode.cn/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Medium (50.50%)
 * Likes:    2471
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 2M
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
 *
 * 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [100,4,200,1,3,2]
 * 输出：4
 * 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,3,7,2,5,8,4,6,0,1]
 * 输出：9
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,0,1,2]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

// @lc code=start
// func longestConsecutive(nums []int) int {
// 	set := make(map[int]bool)
// 	for _, n := range nums {
// 		set[n] = true
// 	}
// 	var res int
// 	for k, _ := range set {
// 		// 找到 起点 num
// 		if _, ok := set[k-1]; ok {
// 			continue
// 		}
// 		// 从起点num 开始 遍历字典 set
// 		count := 0
// 		for start := k; ; {
// 			if set[start] {
// 				count++
// 				start++
// 			} else {
// 				// [k, start)
// 				res = max(res, start-k)
// 				break
// 			}
// 		}
// 	}
// 	return res
// }

// Disjoint Set
func longestConsecutive(nums []int) int {
	size := make(map[int]int)
	ancestor := make(map[int]int)
	ntoi := make(map[int]int)
	for i, num := range nums {
		ancestor[num] = num
		size[num] = 1
		ntoi[num] = i
	}

	var find func(int) int
	find = func(k int) int {
		if k == ancestor[k] {
			return k
		}

		kRoot := find(ancestor[k])
		// 压缩
		ancestor[k] = kRoot
		return kRoot
	}

	join := func(u, v int) {
		ur := find(u)
		vr := find(v)
		if ur == vr {
			return
		}
		ancestor[ur] = vr
		// * update 连通域长度， 存储到根上 *
		size[vr] += size[ur]
	}
	// 获取连通域长度， 从连通域任意一个节点都可以获取这个长度
	getSize := func(k int) int {
		kr := find(k)
		return size[kr]
	}
	res := 0
	for _, num := range nums {
		if _, ok := ntoi[num-1]; ok {
			join(num, num-1)
		}
		if _, ok := ntoi[num+1]; ok {
			join(num, num+1)
		}
		res = max(res, getSize(num))
	}

	return res

}

// @lc code=end
