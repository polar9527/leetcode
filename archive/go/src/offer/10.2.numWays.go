package main

func numWays(n int) int {
	a, b := 1, 1
	m := 1000000007
	for i := 0; i < n; i++ {
		a, b = b%m, (a+b)%m
	}
	return a
}
