package offer2

func countBits(n int) []int {

	countBits := make([]int, n+1)
	for i := 0; i <= n; i++ {
		countBits[i] = countBits[i>>1] + (i & 1)
	}

	return countBits
}
