package main

func add(a int, b int) int {
	carry := (a & b) << 1
	sum := a ^ b
	a = sum
	b = carry
	for carry != 0 {
		carry = (a & b) << 1
		sum = a ^ b
		a = sum
		b = carry
	}
	return a
}
