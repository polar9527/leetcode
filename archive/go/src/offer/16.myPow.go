package main

import (
	"math"
)

func myPow(x float64, n int) float64 {
	if math.Abs(x-0.0) < 0.00001 {
		return 0.0
	}
	if n == 0 {
		return 1.0
	} else if n < 0 {
		x = 1.0 / x
		n = -n
	}
	ret := 1.0
	for n > 1 {
		if n&1 != 0 {
			ret = ret * x
			n--
		} else {
			x = x * x
			n /= 2
		}
	}

	ret *= x

	return ret
}
