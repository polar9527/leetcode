package main

func fib(n int) int {

	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = a+b, a
	}
	return a
}
