package main

import (
	"math"
)

// recursive + dynamic programming
// dp[n][sum] = sum(dp[n-1][sum-i], 1<=i<=6)
// dp[n][sum] = dp[n-1][sum-1] + dp[n-1][sum-2] ...... + dp[n-1][sum-6]
// n 和 sum 之间的关系是 sum >= n
func twoSum(n int) []float64 {

	dp := make(map[int]map[int]int)
	for i := 1; i <= n; i++ {
		dp[i] = make(map[int]int)
	}

	for nthDice := 1; nthDice <= n; nthDice++ {
		for curSum := nthDice; curSum <= 6*nthDice; curSum++ {
			if nthDice == 1 {
				dp[nthDice][curSum] = 1
				continue
			}

			for i := 1; i <= 6; i++ {
				// curSum >= nthDice
				// 同理 curSum-i >= nthDice-1, 防止i比较大，curSum-i 出界
				if curSum-i >= nthDice-1 {
					dp[nthDice][curSum] += dp[nthDice-1][curSum-i]
				}
			}
		}
	}
	result := make([]float64, 0)
	for i := n; i <= 6*n; i++ {
		result = append(result, float64(dp[n][i])*math.Pow(1.0/6.0, float64(n)))
	}
	return result
}

// recursive
func twoSumRecursive(n int) []float64 {
	if n <= 0 {
		return nil
	}
	result := make([]float64, 0)
	for curSum := n; curSum <= 6*n; curSum++ {
		s := countDiceSum(curSum, n)
		r := math.Pow(1.0/6.0, float64(n))
		result = append(result, float64(s)*r)
	}
	return result
}

func countDiceSum(curSum, nthDice int) int {
	if nthDice < 0 || curSum < 0 {
		return 0
	}
	if nthDice == 0 && curSum == 0 {
		return 1
	}
	var sum int
	for i := 1; i <= 6; i++ {
		sum += countDiceSum(curSum-i, nthDice-1)
	}
	return sum
}

func main() {
	twoSumRecursive(1)
}
