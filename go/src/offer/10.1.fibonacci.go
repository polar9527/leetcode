package main

func fib(n int) int {

	if n < 0 {
		return -1
	}
	result := []int{0, 1}
	if n < 2 {
		return result[n]
	}

	fibN := 0
	fib1 := 0
	fib2 := 1
	for i := 2; i <= n; i++ {
		fibN = fib2 + fib1
		fib1 = fib2
		fib2 = fibN
	}
	return fibN
}
