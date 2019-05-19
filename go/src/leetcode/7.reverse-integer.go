/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
package main

func reverse(x int) int {
	// 1111 ...... 1111
	// 2147483648  			uint32
	// 0111 ...... 1111		int32
	// 2147483648 - 1 = 2147483647	int32Max
	const int32Max = int32(^uint32(0) >> 1)

	// 1111 ...... 1111
	// -(2147483648 - 1)  int32
	// 1000 ...... 0000
	// -2147483648        int32		int32Min
	const int32Min = ^int32Max

	var pop int32
	var reverse int32
	pop, reverse = 0, 0

	if x == 0 {
		return x
	}

	for x != 0 {
		pop = int32(x % 10)
		x /= 10

		if reverse > int32Max/10 || (reverse == int32Max/10 && pop > 7) {
			return 0
		}
		if reverse < int32Min/10 || (reverse == int32Max/10 && pop < -8) {
			return 0
		}
		reverse = reverse*10 + pop

	}

	// if x < 0 {
	// 	return -1 * x
	// }
	return int(reverse)
}

// func main() {
// 	print(reverse(-12))
// }
