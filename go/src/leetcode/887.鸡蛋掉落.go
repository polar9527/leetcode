/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */
package leetcode

import "math"

// @lc code=start
func superEggDrop(k int, n int) int {
	// return superEggDropConverse(k, n)
	// return superEggDropDP(k, n)
	return superEggDropRecursive(k, n)
}

func superEggDropConverse(k int, n int) int {
	var floors = n
	var eggs = k
	// var minimalTry = math.MaxInt32
	if eggs == 1 || floors == 0 || floors == 1 {
		return floors
	}
	// dp[eggs][times<=floors]
	dp := make(map[int]map[int]int)
	for i := 1; i <= eggs; i++ {
		dp[i] = make(map[int]int)
		dp[i][1] = 1
	}
	ans := -1
	// times <= floors
	times := floors
	for itime := 2; itime <= times; itime++ {
		for iegg := 1; iegg <= eggs; iegg++ {
			dp[iegg][itime] = 1 + dp[iegg-1][itime-1] + dp[iegg][itime-1]
		}
		if dp[eggs][itime] >= n {
			ans = itime
			break
		}
	}
	return ans
}

func superEggDropDP(k int, n int) int {
	// return superEggDropRecursive(k, n)
	// dp[k][n]
	dpMap := make(map[int]int, k*n)

	var dpFunc func(int, int) int
	dpFunc = func(k int, n int) int {
		var floors = n
		var eggs = k
		var minimalTry = math.MaxInt32

		if v, ok := dpMap[floors*100+eggs]; !ok {
			if eggs == 1 || floors == 0 || floors == 1 {
				minimalTry = floors
			} else {
				lo, hi := 1, n
				for lo+1 < hi {
					// lo + 1 <= mi
					mi := (lo + hi) / 2

					t1 := dpFunc(eggs-1, mi-1)
					t2 := dpFunc(eggs, floors-mi)

					if t1 < t2 {
						// 所以lo每次至少相对于之前，前进了1位
						lo = mi
					} else if t1 > t2 {
						// hi 至少比 lo 高1位
						hi = mi
					} else {
						lo, hi = mi, mi
					}
				}
				minimalTry = 1 + min(max(dpFunc(eggs-1, lo-1), dpFunc(eggs, floors-lo)), max(dpFunc(eggs-1, hi-1), dpFunc(eggs, floors-hi)))
			}
			dpMap[floors*100+eggs] = minimalTry
		} else {
			minimalTry = v
		}
		return minimalTry
	}
	return dpFunc(k, n)
}

func superEggDropRecursive(k int, n int) int {

	// k=鸡蛋， n=楼层
	if k == 1 || n == 0 || n == 1 {
		return n
	}

	var floors = n
	var eggs = k
	var minimalTry = math.MaxInt32

	for floor := 1; floor <= floors; floor++ {
		breaks := superEggDrop(eggs-1, floor-1)
		noneBreaks := superEggDrop(eggs, floors-floor)
		iTry := 1 + max(breaks, noneBreaks)
		minimalTry = min(iTry, minimalTry)
	}

	return minimalTry

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
