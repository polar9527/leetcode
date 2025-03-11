package go_case

import (
	"container/heap"
)

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 *
 * https://leetcode.cn/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (63.53%)
 * Likes:    1650
 * Dislikes: 0
 * Total Accepted:    455.8K
 * Total Submissions: 717.4K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,1,1,2,2,3], k = 2
 * 输出: [1,2]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1], k = 1
 * 输出: [1]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * k 的取值范围是 [1, 数组中不相同的元素的个数]
 * 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
 *
 *
 *
 *
 * 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
 *
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {

	set := make(map[int]int, 0)
	for _, num := range nums {
		set[num]++
	}

	hPri := &kHeap{}
	heap.Init(hPri)

	for num, feq := range set {
		heap.Push(hPri, [2]int{num, feq})
		if hPri.Len() > k {
			heap.Pop(hPri)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(hPri).([2]int)[0]
	}
	return res
}

type kHeap [][2]int

func (k kHeap) Len() int {
	return len(k)
}

func (k kHeap) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k kHeap) Less(i, j int) bool {
	return k[i][1] < k[j][1]
}

func (k *kHeap) Push(x any) {
	*k = append(*k, x.([2]int))
}
func (k *kHeap) Pop() (x any) {
	old := *k
	l := len(old)
	x = old[l-1]
	*k = old[:l-1]
	return x
}

// func topKFrequent(nums []int, k int) []int {

// 	vf := make(map[int]int, 0)
// 	for _, v := range nums {
// 		vf[v]++
// 	}

// 	vfList := []int{}
// 	for v, _ := range vf {
// 		vfList = append(vfList, v)
// 	}

// 	sort.Slice(vfList, func(i, j int) bool {
// 		return vf[vfList[i]] > vf[vfList[j]]
// 	})

// 	return vfList[:k]
// }

// @lc code=end
