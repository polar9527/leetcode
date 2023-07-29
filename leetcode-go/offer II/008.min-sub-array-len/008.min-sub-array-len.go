package offer2

import (
	"math"
	"sort"

	common "github.com/polar9527/leetcode/leetcode-go/internal/common"
)

func minSubArrayLen(target int, nums []int) int {
	return minSubArrayLenPrefixSum(target, nums)
}

func minSubArrayLenTwoPtrWindow(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	ret := math.MaxInt32

	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= target {
			ret = common.Min(ret, end-start+1)
			sum -= nums[start]
			start++

		}
		end++
	}
	if ret == math.MaxInt32 {
		return 0
	}
	return ret
}

func minSubArrayLenPrefixSum(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	sums := make([]int, n+1)
	// 为了方便计算，令 size = n + 1
	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
	// 以此类推
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		bound := sort.SearchInts(sums, target)
		//  if bound == n + 1, that means we did not find the target
		if bound <= n {
			ans = min(ans, bound-(i-1))
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
