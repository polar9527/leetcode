package main

func sumNums(n int) int {
	if n == 1 {
		return 1
	}

	return n + sumNums(n-1)
}
