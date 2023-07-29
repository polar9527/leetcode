package offer2

func subarraySum(nums []int, k int) int {
	n := len(nums)
	count, preSum := 0, 0
	m := map[int]int{}
	// preSum 与 k 相等，那么前缀和 preSum 自己所代表的 子数组也要 计算在内
	m[0] = 1
	for i := 0; i < n; i++ {
		preSum += nums[i]
		if _, ok := m[preSum-k]; ok {
			// 此处 是计算 preSum[i]和preSum[j] 之间的 前缀和差值，也就是 子数组之和k,
			// 而  m[preSum-k] 的值的含义是，所有满足 preSum-k 的前缀和的个数，
			// 表示此处可能有 j1,j2,...,jN ，一共N个已经存在字典m中的前缀和满足当前 与preSum-k相等的条件
			count += m[preSum-k]
		}
		// 而每次都要往字典的这个键值m[preSum]上累加，因为后续的 某个前缀和 需要在这个 累加值上 计算 满足条件
		// 的子数组的个数
		m[preSum] += 1
	}

	return count
}
